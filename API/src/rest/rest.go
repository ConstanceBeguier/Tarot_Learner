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

func GetHandEndpoint(w http.ResponseWriter, req *http.Request) {
	// id := mux.Vars(req)['id']
	//json.NewEncoder(w).Encode(party.Players[0].CardsRemaining)
	json.NewEncoder(w).Encode(party.Table)
}

func GetNewpartyEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(party.Table)
	json.NewEncoder(w).Encode(party.Table)
}

func GetNewpartyAvailableseatEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("GetNewpartyAvailableseatEndpoint")
}

func PostNewpartyAvailableseatEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("PostNewpartyAvailableseatEndpoint")
}

func GetTableEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("GetTableEndpoint")
}

func PostTableEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("PostTableEndpoint")
}

func GetTableTurnEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("GetTableTurnEndpoint")
}

// func GetTableCardsEndpoint(w http.ResponseWriter, req *http.Request) {
// 	json.NewEncoder(w).Encode(party.Table.Cards)
// }

// func GetTableScoresEndpoint(w http.ResponseWriter, req *http.Request) {
// 	json.NewEncoder(w).Encode(party.Table.Scores)
// }

// func GetPlayEndpoint(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintln(w, "{yourTurn: true}")
// }

// func PostPlayCardEndpoint(w http.ResponseWriter, req *http.Request) {
// 	var card tarot.Card
// 	_ = json.NewDecoder(req.Body).Decode(&card)
// 	fmt.Println(card)
// 	json.NewEncoder(w).Encode(card)
// }

func main() {
	party = tarot.NewParty()
	router := mux.NewRouter()
	router.HandleFunc("/hand/{id}", GetHandEndpoint).Methods("GET")
	router.HandleFunc("/newparty", GetNewpartyEndpoint).Methods("GET")
	router.HandleFunc("/newparty/available_seat", GetNewpartyAvailableseatEndpoint).Methods("GET")
	router.HandleFunc("/newparty/available_seat/{id}", PostNewpartyAvailableseatEndpoint).Methods("POST")
	router.HandleFunc("/table", GetTableEndpoint).Methods("GET")
	router.HandleFunc("/table", PostTableEndpoint).Methods("POST")
	router.HandleFunc("/table/{turn}", GetTableTurnEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
