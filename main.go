package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/takanome-dev/blockchain/blockchain"
	"github.com/takanome-dev/blockchain/utils"
)

func main() {
	// port := flag.Int("p", 5000, "port to connect to")
	// flag.Parse()

	http.HandleFunc("/blockchain", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/blockchain" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "only GET requests are accepted for this endpoint", http.StatusMethodNotAllowed)
			return
		}

		chain := blockchain.InitBlockchain()
		
		err := utils.WriteJSON(w, chain)
		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/transaction", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/transaction" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if r.Method != "POST" {
			http.Error(w, "only post requests are accepted for this endpoint", http.StatusMethodNotAllowed)
			return
		}

		body, err := utils.ReadJSON[blockchain.Transaction](r.Body)
		if err != nil {
			utils.WriteError(w, err, http.StatusBadRequest)
			return
		}

		fmt.Printf("request body: %+v\n", body)

		type response struct{
			Message string
		}

		err = utils.WriteJSON(w, response{
			Message: "Body received!!",
		})
		if err != nil {
			utils.WriteError(w, err, http.StatusInternalServerError)
			return
		}
	})
	
	http.HandleFunc("/mine", func (w http.ResponseWriter, r *http.Request) {
		// return the blockchain
	})

	log.Printf("connected to localhost:5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}