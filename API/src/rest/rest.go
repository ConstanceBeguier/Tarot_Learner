package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tarot"
)

var party tarot.Party

func GetTableEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Table)
}

func GetTableCardsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Table.Cards)
}

func GetTableScoresEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Table.Scores)
}

func GetPlayEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "{yourTurn: true}")
}

func PostPlayCardEndpoint(w http.ResponseWriter, req *http.Request) {
	var card tarot.Card
	_ = json.NewDecoder(req.Body).Decode(&card)
	fmt.Println(card)
	json.NewEncoder(w).Encode(card)
}

func main() {
	party = tarot.NewParty()
	/*router := mux.NewRouter()
	router.HandleFunc("/table", GetTableEndpoint).Methods("GET")
	router.HandleFunc("/table/cards", GetTableCardsEndpoint).Methods("GET")
	router.HandleFunc("/table/scores", GetTableScoresEndpoint).Methods("GET")
	router.HandleFunc("/play", GetPlayEndpoint).Methods("GET")
	router.HandleFunc("/play/card", PostPlayCardEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))*/
}
