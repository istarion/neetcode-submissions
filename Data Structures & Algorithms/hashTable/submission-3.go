
type HashTable struct {
	entries  []*HashEntry
	size     int
	capacity int
}

type HashEntry struct {
	key   int
	value int
}

const HASH_TABLE_LOAD_FACTOR = 0.5

func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		entries:  make([]*HashEntry, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ht *HashTable) Insert(key, value int) {
	if float64(ht.size+1)/float64(ht.capacity) >= HASH_TABLE_LOAD_FACTOR {
		ht.Resize()
	}
	hash := key % ht.capacity
	for {
		if ht.entries[hash] == nil {
			ht.entries[hash] = &HashEntry{key: key, value: value}
			ht.size++
			return
		} else if ht.entries[hash].key == key {
			ht.entries[hash].value = value
			return
		}
		hash = (hash + 1) % ht.capacity
	}
}

func (ht *HashTable) Get(key int) int {
	hash := key % ht.capacity
	for {
		if ht.entries[hash] != nil && ht.entries[hash].key == key {
			return ht.entries[hash].value
		}
		if ht.entries[hash] == nil {
			return -1
		}
		hash = (hash + 1) % ht.capacity
	}
}

func (ht *HashTable) Remove(key int) bool {
	hash := key % ht.capacity
	for {
		if ht.entries[hash] != nil && ht.entries[hash].key == key {
			ht.entries[hash] = nil
			ht.size--
			return true
		}
		if ht.entries[hash] == nil {
			return false
		}
		hash = (hash + 1) % ht.capacity
	}
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

func (ht *HashTable) GetCapacity() int {
	return ht.capacity
}

func (ht *HashTable) Resize() {
	newCapacity := ht.capacity << 1
	newEntries := make([]*HashEntry, newCapacity)
	for _, entry := range ht.entries {
		if entry != nil {
			hash := entry.key % newCapacity
			for {
				if newEntries[hash] == nil {
					newEntries[hash] = entry
					break
				}
				hash = (hash + 1) % newCapacity
			}
		}
	}

	ht.entries = newEntries
	ht.capacity = newCapacity
}
