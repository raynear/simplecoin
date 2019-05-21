package main

import (
	"fmt"

	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

func main() {
	aString := "this is test sentence"
	bString := "raynear"
	fmt.Println(aString)
	fmt.Println(bString)

	h := sha3.New256()

	h.Write([]byte(aString))
	txt := h.Sum(nil)

	aa := hex.EncodeToString(txt)
	fmt.Println(aa)
	bb := base58.Encode(txt)
	fmt.Println(bb)

	h.Write([]byte(bString))
	ttt := h.Sum(nil)

	aa = hex.EncodeToString(ttt)
	fmt.Println(aa)
	bb = base58.Encode(ttt)
	fmt.Println(bb)
}
