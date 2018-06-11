package main

import (
	coin "SimpleCoin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func transaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("transaction post")
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
	fmt.Println("Receive a Transaction")
	fmt.Println(aTransaction)
	coin.ReceivedTransactions = append(coin.ReceivedTransactions, aTransaction)

	// w 수행 결과 알리기

}

func push2node(aNode string) {
	MyIP := coin.GetMyIP()
	if MyIP == "" {
		return
	}
	resp, err := http.PostForm(aNode+"/addnode", url.Values{"node": {MyIP}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func pullnodes(aNode string) {
	// nodelist에서 nodes 가져오기
	resp, err := http.Get(aNode + "/nodelist")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func addnode(w http.ResponseWriter, r *http.Request) {
	// request를 어떤 방식으로 받을 것인가?
	r.ParseForm()
	newNode := r.Form.Get("node")
	coin.NodeList = append(coin.NodeList, coin.Node{newNode})

	pullnodes(newNode)
}

func nodelist(w http.ResponseWriter, r *http.Request) {
	jData, err := json.Marshal(coin.NodeList)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/transaction", transaction).Methods("POST")
	mux.HandleFunc("/addnode", addnode).Methods("POST")
	mux.HandleFunc("/nodelist", addnode).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(mux)
	go n.Run(":" + coin.Port)

	coin.Genesis()
	coin.Mining()

}
