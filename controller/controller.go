package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/utils"
)
// Controller example
type Games struct {

Room map[string] Match

}

type Match struct{

ID string
P1 Player `json:"P1"`
P2 Player `json:"P2"`
P3 Player `json:"P3"`
P4 Player `json:"P4"`
P5 Player `json:"P5"`
P6 Player `json:"P6"`

}

type Player struct {
	
	UserName string `json:"username"`
	Cards []string `json:"cards"`
	SocketConnection *websocket.Conn

}

// NewController example
func NewController() *Games {
	return &Games{}
}

func (games *Games) CreateGame(ctx *gin.Context){

	game_code := utils.GenerateGameCode(4)
	match := games.Room[game_code]
	match.ID = game_code
	
	//Generating the cards for the game
	suits :=[...]string{"S", "D", "C", "H"}
	values :=[...]string{"A", "2", "3", "4", "5", "6", "7", "9", "10", "J", "Q", "K"}

	var deck [] string

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, (suit+value))
		}
	}
	deck= utils.Shuffle(deck)

	match.P1.Cards=deck[0:7]
	match.P2.Cards=deck[8:15]
	match.P3.Cards=deck[16:23]
	match.P4.Cards=deck[24:31]
	match.P5.Cards=deck[32:39]
	match.P6.Cards=deck[40:47]

}