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

type SucceedJson struct {
	Succeed bool `json:"succeed"`
}

type ReadyJson struct {
	Ready bool `json:"ready"`
}

type PlayerTurnJson struct {
	PlayerTurn int `json:"playerTurn"`
}

var party tarot.Party

func GetHandEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	json.NewEncoder(w).Encode(party.Players[id].CardsToJson())
}

func GetNewpartyEndpoint(w http.ResponseWriter, req *http.Request) {
	party = tarot.NewParty()
	succeed := SucceedJson{Succeed: true}
	json.NewEncoder(w).Encode(succeed)
}

func GetNewpartyStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	ready := true
	for _, seat := range party.Seats.AvailableSeats {
		if seat {
			ready = false
		}
	}
	r := ReadyJson{Ready: ready}
	json.NewEncoder(w).Encode(r)
}

func GetNewpartyAvailableseatsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Seats)
}

func PostNewpartyAvailableseatsEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	party.Seats.AvailableSeats[id] = false
	json.NewEncoder(w).Encode(party.Seats)
}

func GetTableEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Table)
}

func PostTableEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	color, _ := strconv.Atoi(mux.Vars(req)["color"])
	number, _ := strconv.Atoi(mux.Vars(req)["number"])
	c := tarot.Card{Color: tarot.Color(color), Number: number}
	b := party.PlayCard(c, id)
	succeed := SucceedJson{Succeed: b}
	json.NewEncoder(w).Encode(succeed)
}

func GetTablePlayerTurnEndpoint(w http.ResponseWriter, req *http.Request) {
	turn := PlayerTurnJson{PlayerTurn: party.Table.PlayerTurn}
	json.NewEncoder(w).Encode(turn)
}

//TODO: Ready to play the newt trick
func GetTableTrickIdEndpoint(w http.ResponseWriter, req *http.Request) {
	trick, _ := strconv.Atoi(mux.Vars(req)["trick"])
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	fmt.Println(trick, id)
	json.NewEncoder(w).Encode("{}")
}

func main() {
	party = tarot.NewParty()
	router := mux.NewRouter()
	router.HandleFunc("/hand/{id}", GetHandEndpoint).Methods("GET")
	router.HandleFunc("/newparty", GetNewpartyEndpoint).Methods("GET")
	router.HandleFunc("/newparty/status", GetNewpartyStatusEndpoint).Methods("GET")
	router.HandleFunc("/newparty/available_seats", GetNewpartyAvailableseatsEndpoint).Methods("GET")
	router.HandleFunc("/newparty/available_seats/{id}", PostNewpartyAvailableseatsEndpoint).Methods("POST")
	router.HandleFunc("/table", GetTableEndpoint).Methods("GET")
	router.HandleFunc("/table/{id}/{color}/{number}", PostTableEndpoint).Methods("POST")
	router.HandleFunc("/table/turn", GetTablePlayerTurnEndpoint).Methods("GET")
	router.HandleFunc("/table/{trick}/{id}", GetTableTrickIdEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
