package SimpleCoin

// Transaction 구조체
type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint32 `json:"amount"`
}

// ReceivedTransactions : 받은 Transaction들
var ReceivedTransactions []Transaction

/*
func main() {
	aTransaction := Transaction{"raynear", "cuteyun", 10}
	fmt.Println(json.Marshal(aTransaction))
}
*/
