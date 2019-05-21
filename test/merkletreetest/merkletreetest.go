package main

import (
    "crypto/sha256"
    "fmt"
    "log"

    "github.com/btcsuite/btcutil/base58"
    "golang.org/x/crypto/sha3"
)

type MerkleNode struct {
    Left  *MerkleNode
    Right *MerkleNode
    Data  []byte
}

func err(msg string, e error) {
    if e != nil {
        fmt.Println(msg, e)
        log.Fatal(msg, e)
    }
}

func NewMerkleNode(left *MerkleNode, right *MerkleNode, data []byte) *MerkleNode {
    mNode := MerkleNode{}

    if left == nil && right == nil {
        hash := sha256.Sum256(data)
        mNode.Data = hash[:]
    } else {
        prevHashes := append(left.Data, right.Data...)
        hash := sha256.Sum256(prevHashes)
        mNode.Data = hash[:]
    }

    mNode.Left = left
    mNode.Right = right

    return &mNode
}

func NewMerkleTree(data [][]byte) *MerkleNode {
    var nodes []MerkleNode

    if len(data)%2 != 0 {
        data = append(data, data[len(data)-1])
    }

    for _, aData := range data {
        node := NewMerkleNode(nil, nil, aData)
        nodes = append(nodes, *node)
    }

    for i := 0; i < len(data)/2; i++ {
        var newLevel []MerkleNode

        for j := 0; j < len(nodes); j += 2 {
            node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
            newLevel = append(newLevel, *node)
        }
        nodes = newLevel
    }

    MerkleRoot := nodes[0]

    return &MerkleRoot
}

func str2Base58Hash(str string) string {
    h := sha3.New256()

    h.Write([]byte(str))
    txt := h.Sum(nil)

    base58str := base58.Encode(txt)
    fmt.Println("base58str", base58str)

    return base58str
}

func sumHash(lhs string, rhs string) string {
    base58lhs := str2Base58Hash(lhs)
    base58rhs := str2Base58Hash(rhs)

    sumString := base58lhs + base58rhs
    fmt.Println("sumString", sumString)

    base58Ss := str2Base58Hash(sumString)

    return base58Ss
}

func main() {
    s := []string{"one", "two", "three", "four", "five", "six"}

    for i := 0; i < len(s)/2.0; i++ {
        s0 := s[2*i]
        s1 := s[2*i+1]

        sumHash(s0, s1)
    }
}
