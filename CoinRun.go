package main

import (
	coin "SimpleCoin"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// MyPort from command line
var MyPort string

// MyWallet from command line
var MyWallet string

func transaction(w http.ResponseWriter, r *http.Request) {
	Body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var aTransaction coin.Transaction
	err = json.Unmarshal(Body, &aTransaction)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(aTransaction)
	coin.ReceivedTransactions = append(coin.ReceivedTransactions, aTransaction)

	// w 수행 결과 알리기

}

func push2node(aNode string) {
	MyIP := coin.GetMyIP()
	if MyIP == "" {
		return
	}
	var bNode coin.Node
	bNode = coin.Node{Address: "http://" + MyIP + ":" + MyPort}
	nodebyte, _ := json.Marshal(bNode)
	buff := bytes.NewBuffer(nodebyte)

	resp, err := http.Post(aNode+"/addnode", "application/json", buff)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func pullnodes(aNode string) []coin.Node {

	var Nodes []coin.Node
	// nodelist에서 nodes 가져오기
	resp, err := http.Get(aNode + "/getnodelist")
	if err != nil {
		fmt.Println(err)
		return Nodes
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		fmt.Println(str)
	}
	json.Unmarshal(respBody, &Nodes)
	fmt.Println(Nodes)

	return Nodes
}

func addnode(w http.ResponseWriter, r *http.Request) {
	// request를 어떤 방식으로 받을 것인가?
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var node coin.Node
	err = json.Unmarshal(b, &node)
	if err != nil {
		panic(err)
	}
	coin.NodeList = append(coin.NodeList, coin.Node{Address: node.Address})

	fmt.Println(coin.NodeList)

	if strings.Contains(node.Address, "localhost") || strings.Contains(node.Address, "127.0.0.1") {
		return
	}
	Nodes := pullnodes(node.Address)

	for _, aNode := range Nodes {
		push2node(aNode.Address)
	}
}

func getnodelist(w http.ResponseWriter, r *http.Request) {
	jData, err := json.Marshal(coin.NodeList)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func listenmakeblock(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var Announce coin.Announce
	err = json.Unmarshal(b, &Announce)
	if err != nil {
		panic(err)
	}
	fmt.Println(Announce)
	coin.AnnouncedBlock = Announce

	coin.StopFlag = true
}

func getblock(w http.ResponseWriter, r *http.Request) {
	blocknumberstr := r.URL.Query()["blocknumber"]
	blocknumber, err := strconv.Atoi(blocknumberstr[0])
	if err != nil {
		fmt.Println(err)
	}
	jData, err := json.Marshal(coin.BlockChain[uint64(blocknumber)])
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func main() {

	MyPort = os.Args[1]
	coin.Port = MyPort

	MyWallet = os.Args[2]
	coin.Miner = MyWallet

	mux := mux.NewRouter()
	mux.HandleFunc("/transaction", transaction).Methods("POST")
	mux.HandleFunc("/addnode", addnode).Methods("POST")
	mux.HandleFunc("/getnodelist", getnodelist).Methods("GET")
	mux.HandleFunc("/listenmakeblock", listenmakeblock).Methods("POST")
	mux.HandleFunc("/getblock", getblock).Methods("GET")

	//	n := negroni.Classic()
	n := negroni.New()
	n.UseHandler(mux)
	go n.Run(":" + MyPort)

	coin.Genesis()
	coin.Mining()
}
