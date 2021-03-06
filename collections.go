package minidb

import (
	"encoding/json"
	"path"
	"sync"
)

// base function for creating a new collection
func minicollection(filename string, defaultValue interface{}) *MiniCollections {
	db := &MiniCollections{
		content: []interface{}{},
		mutexes: make(map[int]*sync.Mutex),
		BaseMiniDB: BaseMiniDB{
			db:    filename,
			mutex: &sync.Mutex{},
		},
	}

	content, f := ensureInitialDB(db.db, defaultValue, "[]")
	err := json.Unmarshal(content, &db.content)
	logError(err, "(collections) failed to unmarshall content to struct")

	if f {
		db.writeToDB()
	}

	return db
}

// Collections creates a new key with an array / slice value.
func (db *MiniDB) Collections(key string) *MiniCollections {
	d := db.getOrCreateMutex("cols_" + key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.content.Collections[key]
	if !ok {
		filename = generateFileName("cols")
	}

	db.content.Collections[key] = filename
	db.writeToDB()

	return minicollection(path.Join(db.path, filename), nil)
}

// CollectionsWithDefault creates a new key with an array / slice value. If the key doesn't exist,
// it will write the defaultValue as the first data to the db.
func (db *MiniDB) CollectionsWithDefault(key string, defaultValue interface{}) *MiniCollections {
	d := db.getOrCreateMutex("cols_" + key)
	d.Lock()
	defer d.Unlock()

	// if the key exists, get the file's name,
	// otherwise, create a new one
	filename, ok := db.content.Collections[key]
	if !ok {
		filename = generateFileName("cols")
	}

	db.content.Collections[key] = filename
	db.writeToDB()

	return minicollection(path.Join(db.path, filename), defaultValue)
}
