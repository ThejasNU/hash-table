package hashtable

import "errors"

var PRIME1 = 151
var PRIME2 = 163
var DELETED_ITEM = item{"", ""}

func getHash(s string, numBuckets int, attempt int) int {
	hash1 := polynomialRollingHashFunc(s, PRIME1, numBuckets)
	hash2 := polynomialRollingHashFunc(s, PRIME2, numBuckets) + 1

	if hash2%numBuckets == 0 {
		return (hash1 + (attempt * (hash2 - 1))) % numBuckets
	} else {
		return (hash1 + (attempt * (hash2))) % numBuckets
	}
}

func (ht *table) Set(key string, value string) {
	load := (ht.itemsCount * 100) / ht.size
	if load > 70 {
		ht.scaleUp()
	}

	newItem := createNewItem(key, value)
	index := getHash(key, ht.size, 0)
	curItem := ht.items[index]
	attempt := 1
	for curItem != nil && curItem != &DELETED_ITEM {
		//for update case
		if curItem.key == key {
			ht.items[index].value = value
			return
		}
		//continue searching if key doesn't exist
		index = getHash(key, ht.size, attempt)
		curItem = ht.items[index]
		attempt++
	}

	ht.items[index] = newItem
	ht.itemsCount++
}

func (ht *table) Get(key string) (string, error) {
	index := getHash(key, ht.size, 0)
	curItem := ht.items[index]
	attempt := 1
	for curItem != nil {
		if curItem != &DELETED_ITEM && curItem.key == key {
			return curItem.value, nil
		}
		index = getHash(key, ht.size, attempt)
		curItem = ht.items[index]
		attempt++
	}
	return "", errors.New("key not found in hash-table!!")
}

// if we delete an item at an index, it may lead to not reaching the item at the end of collision chain while searching,so just mark element at that position as deleted
func (ht *table) Delete(key string) error {
	load := (ht.itemsCount * 100) / ht.size
	if load < 10 {
		ht.scaleDown()
	}

	index := getHash(key, ht.size, 0)
	curItem := ht.items[index]
	attempt := 1
	for curItem != nil {
		if curItem != &DELETED_ITEM {
			if curItem.key == key {
				deleteItem(curItem)
				ht.items[index] = &DELETED_ITEM
				ht.itemsCount--
				return nil
			}
		}
		index = getHash(key, ht.size, attempt)
		curItem = ht.items[index]
		attempt++
	}
	return errors.New("key not found!!")
}
