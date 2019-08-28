package dwarfdb

import (
  "errors"
  "gopkg.in/vmihailenco/msgpack.v2"
  "io/ioutil"
  "os"
)

func pathExists(path string) (bool, error) {
  _, err := os.Stat(path)
  if err == nil {
    return true, nil
  }
  if os.IsNotExist(err) {
    return false, nil
  }
  return false, err
}


// DwarfDB structure
type DwarfDB struct {
  path  string
  force bool
  db    map[string]interface{}
}

// DwarfDBLoad loads a new DwarfDB
func DwarfDBLoad(path string, force bool) DwarfDB {
  exist, err := pathExists(path)
  if err != nil {
    panic(err)
  }
  ddb := DwarfDB{path, force, make(map[string]interface{})}
  if exist {
    ddb.loaddb()
  }
  return ddb
}

// Dump force the writing to the file system
func (ddb *DwarfDB) Dump() bool {
  ddb.dumpdb(true)
  return true
}

func (ddb *DwarfDB) loaddb() bool {
  input, err := ioutil.ReadFile(ddb.path)
  if err != nil {
    panic(err)
  }

  if err := msgpack.Unmarshal([]byte(input), &ddb.db); err != nil {
    panic(err)
  }

  return true
}

func (ddb *DwarfDB) dumpdb(force bool) bool {
  dump, _ := msgpack.Marshal(ddb.db)
  err := ioutil.WriteFile(ddb.path, []byte(string(dump)), 0644)
  if err != nil {
    panic(err)
  }
  return true
}

// Set a new key/value in DwarfDB
func (ddb *DwarfDB) Set(key string, value interface{}) bool {
  ddb.db[key] = value
  ddb.dumpdb(ddb.force)
  return true
}

// Set a new key/value in DwarfDB
func (ddb *DwarfDB) Exists(key string) bool {
  _, ok := ddb.db[key]
  return ok
}

// Get a value in DwarfDB
func (ddb *DwarfDB) Get(key string) (interface{}, error) {
  value, ok := ddb.db[key]
  if ok {
    return value, nil
  }
  return nil, errors.New("not found")
}


// GetAll keys in the DwarfDB
func (ddb *DwarfDB) GetAll() []string {
  keys := make([]string, 0, len(ddb.db))
  for k := range ddb.db {
    keys = append(keys, k)
  }
  return keys
}


// Len number of keys in the DwarfDB
func (ddb *DwarfDB) Len() int {
  return len(ddb.db)
}

// Rem removes a key/value od DwarfDB
func (ddb *DwarfDB) Rem(key string) bool {
  delete(ddb.db, key)
  ddb.dumpdb(ddb.force)
  return true
}

// DelDB clean the DwarfDB
func (ddb *DwarfDB) DelDB() bool {
  ddb.db = make(map[string]interface{})
  ddb.dumpdb(ddb.force)
  return true
}
