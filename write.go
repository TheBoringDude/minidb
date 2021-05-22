package minidb

import (
	"io/ioutil"
)

// writes the file
// it just wraps the ioutil
func write(file string, data []byte) error {
	return ioutil.WriteFile(file, data, 0755)
}

// writeToDB write the db.store to the defined json db file.
func (db *MiniDB) writeToDB() {
	data := marshalStore(db.store)

	write(db.db, data)
}

// writeToDB writes the cols.store to the defined json db file.
func (cols *MiniCollections) writeToDB() {
	data := marshalStore(cols.store)

	write(cols.db, data)
}

// writeToDB writes the st.store to the defined json db file.
func (st *MiniStore) writeToDB() {
	data := marshalStore(st.store)

	write(st.db, data)
}
