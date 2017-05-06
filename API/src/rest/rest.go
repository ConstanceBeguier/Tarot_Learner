package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"tarot"
)

type SucceedJson struct {
	Succeed bool `json:"succeed"`
}

type NewPartyJson struct {
	Succeed bool `json:"succeed"`
	Seat    int  `json:"seat"`
}

type ReadyJson struct {
	Ready bool `json:"ready"`
}

type PlayerTurnJson struct {
	PlayerTurn int `json:"playerTurn"`
}

type HistoryJson struct {
	FirstPlayer int                          `json:"firstPlayer"`
	Cards       [tarot.NB_PLAYERS]tarot.Card `json:"cards"`
}

var party tarot.Party

/**
 * @api {get} /hand/:id Request Hand information.
 * @apiName GetHandEndpoint
 * @apiGroup Hand
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {List} cards List of cards.
 */
func GetHandEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	json.NewEncoder(w).Encode(party.Players[id].CardsToJson())
}

/**
 * @api {get} /newparty Start a new party
 * @apiName GetNewpartyEndpoint
 * @apiGroup Newparty
 *
 * @apiSuccess {Boolean} succeed Does the party successfuly start.
 */
func GetNewpartyEndpoint(w http.ResponseWriter, req *http.Request) {
	party = tarot.NewParty()
	party.Seats.AvailableSeats[0] = false
	npJson := NewPartyJson{Succeed: true, Seat: 0}
	json.NewEncoder(w).Encode(npJson)
}

/**
 * @api {get} /newparty/status Request if all seats are ready.
 * @apiName GetNewpartyStatusEndpoint
 * @apiGroup Newparty
 *
 * @apiSuccess {Boolean} ready Readyness of the party.
 */
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

/**
 * @api {get} /newparty/available_seats Request seats availability.
 * @apiName GetNewpartyAvailableseatsEndpoint
 * @apiGroup Newparty
 *
 * @apiSuccess {List} availableSeats List of booleans.
 */
func GetNewpartyAvailableseatsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Seats)
}

/**
 * @api {post} /newparty/available_seats/:id Take place in the 'id' seat.
 * @apiName PostNewpartyAvailableseatsEndpoint
 * @apiGroup Newparty
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {List} availableSeats List of booleans.
 */
func PostNewpartyAvailableseatsEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	party.Seats.AvailableSeats[id] = false
	json.NewEncoder(w).Encode(party.Seats)
}

/**
 * @api {get} /table Request Hand information.
 * @apiName GetTableEndpoint
 * @apiGroup Table
 *
 * @apiSuccess {[2]float32} scores Actual score of defender/taker.
 * @apiSuccess {[NB_PLAYERS]Card} cards on the table.
 * @apiSuccess {Integer} playerTurn ID of the player turn.
 * @apiSuccess {Integer} firstPlayer ID of the first player who played.
 * @apiSuccess {Integer} trickNb Trick's number.
 * @apiSuccess {[NB_PLAYERS]int} isTaker Return the taker status of players.
 */
func GetTableEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(party.Table)
}

/**
* @api {get} /table/valid_cards/:id Request valid cards in Hand.
 * @apiName GetTableValidCardsEndpoint
 * @apiGroup Table
 *
 * @apiSuccess {Boolean} isYourTurn It is your turn to play.
 * @apiSuccess {[]Card} validCards Return valid cards.
*/
func GetTableValidCardsEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	json.NewEncoder(w).Encode(party.ValidCards(id))
}

/**
 * @api {post} /table/:id/:color/:number Play a card.
 * @apiName PostTableEndpoint
 * @apiGroup Table
 *
 * @apiParam {Number} id Users unique ID.
 * @apiParam {Number} color Color of the playing card.
 * @apiParam {Number} number Number of the playing card.
 *
 * @apiSuccess {Boolean} succeed Does the card can be played.
 */
func PostTableEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	color, _ := strconv.Atoi(mux.Vars(req)["color"])
	number, _ := strconv.Atoi(mux.Vars(req)["number"])
	c := tarot.Card{Color: tarot.Color(color), Number: number}
	b := party.PlayCard(c, id)
	succeed := SucceedJson{Succeed: b}
	json.NewEncoder(w).Encode(succeed)
}

/**
 * @api {get} /table/trick Request Trick information.
 * @apiName GetTablePlayerTurnEndpoint
 * @apiGroup Table
 *
 * @apiSuccess {Boolean} playerTurn Current trick.
 */
func GetTablePlayerTurnEndpoint(w http.ResponseWriter, req *http.Request) {
	trick := PlayerTurnJson{PlayerTurn: party.Table.PlayerTurn}
	json.NewEncoder(w).Encode(trick)
}

//TODO: Ready to play the next trick
/**
 * @api {get} /table/:trick/:id Get ready for the next trick.
 * @apiName GetTableTrickIdEndpoint
 * @apiGroup Table
 *
 * @apiParam {Number} trick Trick Number.
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {Empty} Empty Empty brace.
 */
func GetTableTrickIdEndpoint(w http.ResponseWriter, req *http.Request) {
	// trick, _ := strconv.Atoi(mux.Vars(req)["trick"])
	// id, _ := strconv.Atoi(mux.Vars(req)["id"])
	// fmt.Println(trick, id)
	json.NewEncoder(w).Encode("{}")
}

/**
 * @api {get} /history/:trickNb Get old trick.
 * @apiName GetHistoryEndpoint
 * @apiGroup History
 *
 * @apiParam {Number} trickNb Trick number.
 *
 * @apiSuccess {Number} firstPlayer Trick first player.
 * @apiSuccess {[3]Cards} cards Trick cards.
 */
func GetHistoryEndpoint(w http.ResponseWriter, req *http.Request) {
	trickNb, _ := strconv.Atoi(mux.Vars(req)["trickNb"])
	if trickNb >= 0 && trickNb < tarot.NB_CARDS_PER_PLAYER {
		json.NewEncoder(w).Encode(
			HistoryJson{
				FirstPlayer: party.Table.HistoryFirstPlayer[trickNb],
				Cards:       party.Table.HistoryCards[trickNb],
			})
	} else {
		json.NewEncoder(w).Encode(HistoryJson{})
	}
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
	router.HandleFunc("/table/valid_cards/{id}", GetTableValidCardsEndpoint).Methods("GET")
	router.HandleFunc("/table/{id}/{color}/{number}", PostTableEndpoint).Methods("POST")
	router.HandleFunc("/table/trick", GetTablePlayerTurnEndpoint).Methods("GET")
	router.HandleFunc("/table/{trick}/{id}", GetTableTrickIdEndpoint).Methods("GET")
	router.HandleFunc("/history/{trickNb}", GetHistoryEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
