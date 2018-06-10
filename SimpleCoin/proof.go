package SimpleCoin

import (
	"bytes"
	"crypto/sha256"
	"hash"
	"strconv"
)

// ShaHash : Sha256 hash variable
var ShaHash hash.Hash

// Zeros :
var Zeros []byte

func init() {
	ShaHash = sha256.New()
	Difficulty = 3

	Zeros = []byte("0")
	var i uint8
	for i = 1; i < Difficulty; i++ {
		Zeros = append(Zeros, byte(48))
	}
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

		if bytes.Compare(aHash[len(aHash)-int(Difficulty):], Zeros) == 0 {
			return i
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
