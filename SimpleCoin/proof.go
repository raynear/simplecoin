package SimpleCoin

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hash"
	"strconv"
)

// ShaHash : Sha256 hash variable
var ShaHash hash.Hash

// StopFlag :
var StopFlag bool

// Zeros :
var Zeros []byte

func init() {
	ShaHash = sha256.New()

	Zeros = append(Zeros, byte(0))
	Zeros = append(Zeros, byte(0))
	Zeros = append(Zeros, byte(0))
	//	var i uint8
	//	for i = 1; i < Difficulty; i++ {
	Zeros = append(Zeros, byte(Difficulty))
	//	}

	StopFlag = false

	fmt.Println(Zeros)
}

// ValidProof find proof
func ValidProof(prevHash []byte, proof string) []byte {
	ShaHash.Reset()
	ShaHash.Write(append(prevHash, proof...))
	return ShaHash.Sum(nil)
}

// FindProof :
func FindProof(PrevHash []byte) uint64 {
	var i uint64
	for i = 0; ; i++ {
		aHash := ValidProof(PrevHash, strconv.FormatUint(i, 36))

		if bytes.Compare(aHash[4:], Zeros) < 0 {
			return i
		}

		if StopFlag == true {
			StopFlag = false
			return 0
		}
	}
}

/*
func main() {
	PrevHash := "test123"
	proof := FindProof(PrevHash)
	fmt.Println(proof)
}
*/
