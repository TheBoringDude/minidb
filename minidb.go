package minidb

import (
	"sync"
)

const Version = "0.1.4"

type (
	// BaseMiniDB is the base db structure.
	BaseMiniDB struct {
		db    string // combined path and filename
		mutex *sync.Mutex
	}

	// MiniDB is the base store file.
	MiniDB struct {
		path     string
		filename string
		content    MiniDBContent
		mutexes  map[string]*sync.Mutex
		BaseMiniDB
	}

	// MiniDBContent is the types of MiniDB.store
	MiniDBContent struct {
		Keys        map[string]string `json:"keys"`
		Collections map[string]string `json:"collections"`
		Store       map[string]string `json:"store"`
	}

	// MiniCollections is a collections store.
	MiniCollections struct {
		content   []interface{}
		mutexes map[int]*sync.Mutex
		BaseMiniDB
	}

	// MiniStore is a key-value store.
	MiniStore struct {
		content   map[string]interface{}
		mutexes map[string]*sync.Mutex
		BaseMiniDB
	}
)

// New creates a new MiniDB struct.
// The dir will be created if it doesn't exist and a file named `__default.json` will also be generated.
// It is better to use this in managing multiple json files.
func New(dir string) *MiniDB {
	return minidb(dir, "__default.json")
}

// NewMiniStore creates and returns a new key-store collection json db.
func NewMiniStore(f string) *MiniStore {
	return ministore(f)
}

// NewMiniCollections creates and returns a new collections json db.
func NewMiniCollections(f string) *MiniCollections {
	return minicollection(f)
}

// ListCollections returns the list of collections created.
func (db *MiniDB) ListCollections() []string {
	cols := []string{}

	for i := range db.content.Collections {
		cols = append(cols, i)
	}

	return cols
}

// ListCollections returns the list of collections created.
func (db *MiniDB) ListStores() []string {
	stores := []string{}

	for i := range db.content.Store {
		stores = append(stores, i)
	}

	return stores
}
