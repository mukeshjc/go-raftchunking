package raftchunking

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/go-raftchunking/types"
	"github.com/hashicorp/raft"
)

var _ raft.FSM = (*ChunkingFSM)(nil)
var _ raft.ConfigurationStore = (*ChunkingConfigurationStore)(nil)

type ChunkingFSM struct {
	underlying raft.FSM
	store      ChunkStorage
	lastTerm   uint64
}

type ChunkingConfigurationStore struct {
	*ChunkingFSM
	underlyingConfigurationStore raft.ConfigurationStore
}

func NewChunkingFSM(underlying raft.FSM, store ChunkStorage) raft.FSM {
	ret := &ChunkingFSM{
		underlying: underlying,
		store:      store,
	}
	if store == nil {
		ret.store = NewInmemChunkStorage()
	}
	return ret
}

func NewChunkingConfigurationStore(underlying raft.ConfigurationStore, store ChunkStorage) raft.ConfigurationStore {
	ret := &ChunkingConfigurationStore{
		ChunkingFSM: &ChunkingFSM{
			underlying: underlying,
			store:      store,
		},
		underlyingConfigurationStore: underlying,
	}
	if store == nil {
		ret.ChunkingFSM.store = NewInmemChunkStorage()
	}
	return ret
}

// Apply applies the log, handling chunking as needed. The return value will
// either be an error or whatever is returned from the underlying Apply.
func (c *ChunkingFSM) Apply(l *raft.Log) interface{} {
	// Not chunking or wrong type, pass through
	if l.Type != raft.LogCommand || l.Extensions == nil {
		return c.underlying.Apply(l)
	}

	if l.Term != c.lastTerm {
		// Term has changed. A raft library client that was applying chunks
		// should get an error that it's no longer the leader and bail, and
		// then any client of (Consul, Vault, etc.) should then retry the full
		// chunking operation automatically, which will be under a different
		// opnum. So it should be safe in this case to clear the map.
		if err := c.store.ClearAll(); err != nil {
			return err
		}
		c.lastTerm = l.Term
	}

	// Get chunk info from extensions
	var ci types.ChunkInfo
	if err := proto.Unmarshal(l.Extensions, &ci); err != nil {
		return errwrap.Wrapf("error unmarshaling chunk info: {{err}}", err)
	}

	// Store the current chunk and find out if all chunks have arrived
	done, err := c.store.StoreChunk(&ChunkInfo{
		OpNum:       ci.OpNum,
		SequenceNum: ci.SequenceNum,
		NumChunks:   ci.NumChunks,
		Data:        l.Data,
	})
	if err != nil {
		return err
	}
	if !done {
		return nil
	}

	// All chunks are here; get the full set and clear storage of the op
	chunks, err := c.store.FinalizeOp(ci.OpNum)
	if err != nil {
		return err
	}

	finalData := make([]byte, 0, len(chunks)*raft.SuggestedMaxDataSize)
	for _, chunk := range chunks {
		finalData = append(finalData, chunk.Data...)
	}

	// Use the latest log's values with the final data
	logToApply := &raft.Log{
		Index:      l.Index,
		Term:       l.Term,
		Type:       l.Type,
		Data:       finalData,
		Extensions: ci.NextExtensions,
	}

	return c.Apply(logToApply)
}

func (c *ChunkingFSM) Snapshot() (raft.FSMSnapshot, error) {
	return c.underlying.Snapshot()
}

func (c *ChunkingFSM) Restore(rc io.ReadCloser) error {
	return c.underlying.Restore(rc)
}

// Note: this is used in tests via the Raft package test helper functions, even
// if it's not used in client code
func (c *ChunkingFSM) Underlying() raft.FSM {
	return c.underlying
}

func (c *ChunkingFSM) CurrentState() (ChunkMap, error) {
	return c.store.GetAll()
}

func (c *ChunkingConfigurationStore) StoreConfiguration(index uint64, configuration raft.Configuration) {
	c.underlyingConfigurationStore.StoreConfiguration(index, configuration)
}
