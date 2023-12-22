package hashtable

func createNewItem(k string, v string) *item {
	newItem := item{
		key:   k,
		value: v,
	}
	return &newItem
}

func CreateNewHashTable() *table {
	newTable := table{
		size:       53,
		itemsCount: 0,
		items:      make([]*item, 53),
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
