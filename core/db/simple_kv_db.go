package db

type SimpleKVDatabase struct {
	mp map[uint32]uint32
}

func NewSimpleKVDatabase() *SimpleKVDatabase {
	return &SimpleKVDatabase{make(map[uint32]uint32)}
}

func (db *SimpleKVDatabase) Get(key uint32) (uint32, bool) {
	value, exist := db.mp[key]
	return value, exist
}

func (db *SimpleKVDatabase) Put(key uint32, value uint32) {
	db.mp[key] = value
}
