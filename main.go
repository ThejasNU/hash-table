package main

import "github.com/ThejasNU/hash-table/hashtable"

func main() {
	hastTable := hashtable.CreateNewHashTable()
	hashtable.DeleteHashTable(hastTable)
}
