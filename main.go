package main

import (
	"fmt"
	"log"

	"github.com/ThejasNU/hash-table/hashtable"
)

func main() {
	hastTable := hashtable.InitTable()

	hastTable.Set("abc", "123")
	hastTable.Set("abc", "456")

	val, err := hastTable.Get("abc")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(val)
	}

	delErr := hastTable.Delete("abc")
	if delErr != nil {
		log.Fatal(delErr)
	}

	val2, err2 := hastTable.Get("abc")
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println(val2)
	}

}
