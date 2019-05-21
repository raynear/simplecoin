package main

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func err(msg string, e error) {
	if e != nil {
		fmt.Println(msg, e)
		log.Fatal(msg, e)
	}
}

func main() {
	db, e := leveldb.OpenFile("/db/db1.db", nil)
	err("DB Open Error", e)
	defer db.Close()

	e = db.Put([]byte("key1"), []byte("data1"), nil)
	err("DB Input Error", e)
	e = db.Put([]byte("key2"), []byte("data2"), nil)
	err("DB Input Error", e)
	e = db.Put([]byte("key3"), []byte("data3"), nil)
	err("DB Input Error", e)
	e = db.Put([]byte("key4"), []byte("data4"), nil)
	err("DB Input Error", e)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s | Value: %s\n", key, value)
	}

	fmt.Println()

	for ok := iter.Seek([]byte("key2")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s | Value: %s\n", key, value)
	}

	fmt.Println()

	for ok := iter.First(); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s | Value: %s\n", key, value)
	}

	iter.Release()
	e = iter.Error()
	err("iter error", e)
}
