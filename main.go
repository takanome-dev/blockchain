package main

import (
	"log"
	"net/http"
)

func main() {
	// port := flag.Int("p", 5000, "port to connect to")
	// flag.Parse()

	http.HandleFunc("/blockchain", func (w http.ResponseWriter, r *http.Request) {
		// return the blockchain
	})

	http.HandleFunc("/transaction", func (w http.ResponseWriter, r *http.Request) {
		// return the blockchain
	})
	
	http.HandleFunc("/mine", func (w http.ResponseWriter, r *http.Request) {
		// return the blockchain
	})

	log.Printf("connected to localhost:5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}