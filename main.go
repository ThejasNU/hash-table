package main

import (
	"github.com/ThejasNU/hash-table/hashtable"
)

func main() {
	hastTable := hashtable.CreateNewHashTable(53)
	hashtable.DeleteHashTable(hastTable)
}
