package main

import (
	"fmt"
)

type block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func main() {
	fmt.Println("a")
}
