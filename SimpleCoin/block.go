package SimpleCoin

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Block : 체인의 한 블록
type Block struct {
	PrevBlockHash []byte            `json:"prevblockhash"`
	TimeStamp     time.Time         `json:"timestamp"`
	Difficulty    uint8             `json:"difficulty"`
	Proof         uint64            `json:"proof"`
	BlockNumber   uint64            `json:"blocknumber"`
	NodeID        string            `json:"nodeid"`
	Transactions  []Transaction     `json:"transactions"`
	BalanceOf     map[string]uint32 `json:"balanceOf"`
}

// BlockChain :
var BlockChain map[uint64]Block
var currentBlockNumber uint64

// HashBlock :
func HashBlock(aBlock Block) []byte {
	var ReturnHash []byte

	Sha256 := sha256.New()
	MarshalBlock, _ := json.Marshal(aBlock)
	Sha256.Write(MarshalBlock)

	ReturnHash = Sha256.Sum(nil)

	return ReturnHash
}

// executeTransactions : transaction 수행
func executeTransactions(Transactions []Transaction) {
	for _, aTransaction := range Transactions {
		if balanceOf[aTransaction.From] >= aTransaction.Amount {
			balanceOf[aTransaction.From] -= aTransaction.Amount
			balanceOf[aTransaction.To] += aTransaction.Amount
			fmt.Printf("Success : %s send %d to %s\n",
				aTransaction.From, aTransaction.Amount, aTransaction.To)
		} else {
			fmt.Printf("FAIL : %s Try to send %d but have %d\n",
				aTransaction.From, aTransaction.Amount, balanceOf[aTransaction.From])
		}
	}
}

// MakeChain : make a blockchain
func MakeChain(aBlock Block) {
	// execute Transaction
	executeTransactions(ReceivedTransactions)
	ReceivedTransactions = ReceivedTransactions[:0]
	BlockChain[currentBlockNumber] = aBlock
	// anounce
}

// Mining :
func Mining() {
	for {
		proof := FindProof(HashBlock(BlockChain[currentBlockNumber]))
		if proof == 0 {
			if AnnouncedBlock.BlockNumber == currentBlockNumber+1 {
				fmt.Println("AnnouncedBlock - if")
				// Get RecentBlock
				var NewBlock Block
				blocknumberstr := strconv.FormatUint(AnnouncedBlock.BlockNumber, 10)
				httpget := AnnouncedBlock.MinedNode.Address + "/getblock?blocknumber=" + blocknumberstr
				// local에서 실행하기 위해선 윗줄 주석 & 아랫줄로 대체
				//httpget := "http://127.0.0.1:" + AnnouncedBlock.MinedNode.Address[len(AnnouncedBlock.MinedNode.Address)-4:] + "/getblock?blocknumber=" + blocknumberstr
				resp, err := http.Get(httpget)
				if err != nil {
					fmt.Println(err)
				}

				// Response 체크.
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				json.Unmarshal(respBody, &NewBlock)
				fmt.Printf("[Mining] %s Mined! %d block : %d at %s\n", NewBlock.NodeID, NewBlock.BlockNumber, NewBlock.Proof, NewBlock.TimeStamp.Format("2006-01-02 15:04:05"))
				fmt.Println(NewBlock.BalanceOf)
				currentBlockNumber = NewBlock.BlockNumber
				MakeChain(NewBlock)
				balanceOf = NewBlock.BalanceOf
				AnnouncedBlock = Announce{0, Node{""}}
			} else {
				fmt.Println("AnnouncedBlock - else")
				// 언제부터 받아야 되는지 확인해봐야 함
			}
		} else {
			balanceOf[Miner] += uint32(MiningReward)
			currentBlockNumber++
			MakeChain(Block{HashBlock(BlockChain[currentBlockNumber-1]), time.Now(), Difficulty, proof, currentBlockNumber, GetMyIP() + ":" + Port, ReceivedTransactions, balanceOf})
			AnnounceMakeBlock(currentBlockNumber)
			//fmt.Println(BlockChain[currentBlockNumber])
			PrintBlock := BlockChain[currentBlockNumber]
			fmt.Printf("[Mining] LocalNode Mined! %d block : %d at %s\n", PrintBlock.BlockNumber, PrintBlock.Proof, PrintBlock.TimeStamp.Format("2006-01-02 15:04:05"))
			fmt.Println(PrintBlock.BalanceOf)
		}
	}
}

// Genesis :
func Genesis() {
	proof := FindProof(GenesisHash)
	currentBlockNumber = 0
	BlockChain = make(map[uint64]Block)
	balanceOf = make(map[string]uint32)
	//	NodeList = make([]Node)
	balanceOf[Miner] = uint32(10)

	BlockChain[currentBlockNumber] = Block{HashBlock(BlockChain[currentBlockNumber]), time.Now(), Difficulty, proof, currentBlockNumber, GetMyIP() + ":" + Port, ReceivedTransactions, balanceOf}
	AnnounceMakeBlock(currentBlockNumber)
	fmt.Println(BlockChain[currentBlockNumber])
}

// concensus
// announcement
// peer들 등록
