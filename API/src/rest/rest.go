package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"tarot"
)

type NewPartyJson struct {
	Succeed bool `json:"succeed,omitempty"`
}

var party tarot.Party

func GetHandEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	json.NewEncoder(w).Encode(party.Players[id].CardsToJson())
}

func GetNewpartyEndpoint(w http.ResponseWriter, req *http.Request) {
	party = tarot.NewParty()
	newPartyJson := NewPartyJson{Succeed: true}
	json.NewEncoder(w).Encode(newPartyJson)
}

func GetNewpartyAvailableseatEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Seats)
}

func PostNewpartyAvailableseatEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	party.Seats.AvailableSeats[id] = false
	json.NewEncoder(w).Encode(party)
}

func GetTableEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Table)
}

func PostTableEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	color, _ := strconv.Atoi(mux.Vars(req)["color"])
	number, _ := strconv.Atoi(mux.Vars(req)["number"])
	c := tarot.Card{Color: tarot.Color(color), Number: number}
	party.PlayCard(c, id)
	json.NewEncoder(w).Encode(party.Table)
}

//TODO
// Ready to play the newt trick
func GetTableTrickIdEndpoint(w http.ResponseWriter, req *http.Request) {
	trick, _ := strconv.Atoi(mux.Vars(req)["trick"])
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	fmt.Println(trick, id)
	json.NewEncoder(w).Encode("GetTableTurnIdEndpoint")
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
	//router.HandleFunc("/newparty/status", GetNewpartyEndpoint).Methods("GET")
	router.HandleFunc("/newparty/available_seat", GetNewpartyAvailableseatEndpoint).Methods("GET")
	router.HandleFunc("/newparty/available_seat/{id}", PostNewpartyAvailableseatEndpoint).Methods("POST")
	router.HandleFunc("/table", GetTableEndpoint).Methods("GET")
	router.HandleFunc("/table/{id}/{color}/{number}", PostTableEndpoint).Methods("POST")
	//router.HandleFunc("/table/{turn}", GetTableTurnEndpoint).Methods("GET")
	router.HandleFunc("/table/{trick}/{id}", GetTableTrickIdEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
