package hashtable

import (
	"math"
)

var HASHTABLE_BASE_SIZE = 7

func createNewItem(k string, v string) *item {
	newItem := item{
		key:   k,
		value: v,
	}
	return &newItem
}

func CreateNewHashTable(hashTableSize int) *table {
	newTable := table{
		size:       hashTableSize,
		itemsCount: 0,
		items:      make([]*item, hashTableSize),
	}

	return &newTable
}

func deleteItem(htItem *item) {
	htItem = nil
}

func DeleteHashTable(hashTable *table) {
	for i := 0; i < hashTable.size; i++ {
		hashTable.items[i] = nil
	}
	hashTable = nil
}

// Polynomial rolling hash function
func polynomialRollingHashFunc(s string, primeNum int, numBuckets int) int {
	var hashVal int64 = 0
	n := len(s)
	for i := 0; i < n; i++ {
		hashVal += int64(math.Pow(float64(primeNum), float64(n-i-1))) * int64(s[i])
		hashVal %= int64(numBuckets)
	}
	return int(hashVal)
}

func (ht *table) resize(newSize int) {
	if newSize < HASHTABLE_BASE_SIZE {
		return
	}

	newHashTable := CreateNewHashTable(newSize)
	for i := 0; i < ht.size; i++ {
		curItem := ht.items[i]
		if curItem != nil && curItem != &DELETED_ITEM {
			newHashTable.Set(curItem.key, curItem.value)
		}
	}

	ht.size = newHashTable.size
	ht.itemsCount = newHashTable.itemsCount
	ht.items = newHashTable.items
	newHashTable.items = make([]*item, 0)

	DeleteHashTable(newHashTable)
}

func (ht *table) scaleUp(){
	newSize:=ht.size*2
	ht.resize(newSize)
}