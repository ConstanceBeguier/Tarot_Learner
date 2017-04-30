package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tarot"
)

func GetRoundTableEndpoint(w http.ResponseWriter, req *http.Request) {
	var cards []tarot.Card
	cards = append(cards, tarot.Card{Color: 4, Number: 21})
	cards = append(cards, tarot.Card{Color: 4, Number: 15})
	json.NewEncoder(w).Encode(cards)
}

func GetRoundScoresEndpoint(w http.ResponseWriter, req *http.Request) {
	var scores tarot.TableScores
	scores.Scores[0] = 6
	scores.Scores[1] = 15
	json.NewEncoder(w).Encode(scores)
}

func GetPlayEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "{yourTurn: yes}")
}

func PostPlayCardEndpoint(w http.ResponseWriter, req *http.Request) {
	var card tarot.Card
	_ = json.NewDecoder(req.Body).Decode(&card)
	fmt.Println(card)
	json.NewEncoder(w).Encode(card)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/round/table", GetRoundTableEndpoint).Methods("GET")
	router.HandleFunc("/round/scores", GetRoundScoresEndpoint).Methods("GET")
	router.HandleFunc("/play", GetPlayEndpoint).Methods("GET")
	router.HandleFunc("/play/card", PostPlayCardEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}
