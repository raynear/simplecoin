package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func err(msg string, e error) {
	if e != nil {
		fmt.Println(msg, e)
		log.Fatal(msg, e)
	}
}

func main() {
	e := ioutil.WriteFile("c:/GoCode/src/test/filetest/file/file.txt", []byte("data1"), 0644)
	err("File Write Error", e)

	//	f, e := os.Create("c:/GoCode/src/test/filetest/file/file.txt")
	//	err("File Create Error", e)

	//	defer f.Close()
}
