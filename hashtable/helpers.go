package hashtable

import (
	"fmt"
	"math"
)

var HASHTABLE_BASE_SIZE = 5

func createNewItem(k string, v string) *item {
	newItem := item{
		key:   k,
		value: v,
	}
	return &newItem
}

func createNewHashTable(hashTableSize int) *table {
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

func deleteHashTable(hashTable *table) {
	for i := 0; i < len(hashTable.items); i++ {
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
	fmt.Println(ht.size, ht.itemsCount)
	newHashTable := createNewHashTable(newSize)
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
	fmt.Println(ht.size, ht.itemsCount)
	deleteHashTable(newHashTable)
}

func (ht *table) scaleUp() {
	newSize := getNextPrime(ht.size * 2)
	ht.resize(newSize)
}

func (ht *table) scaleDown() {
	newSize := getNextPrime(ht.size / 2)
	ht.resize(newSize)
}

func getNextPrime(num int) int {
	for !isPrime(num) {
		num++
	}

	return num
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if num < 4 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i < int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
