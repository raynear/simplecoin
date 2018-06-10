package main

import (
	coin "SimpleCoin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/transaction", transaction).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(mux)
	go n.Run(":" + coin.Port)

	coin.Genesis()
	coin.Mining()

}
