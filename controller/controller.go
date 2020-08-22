package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/utils"
	"github.com/anish-sekar/literature-backend/models"
)
// Controller example
type World struct {

Games map[string]*Match

}

type Match struct{

CurrentTurn int
Players []*Player 
Team1 []int 
Team2 []int

}

type Player struct {
	
	UserName string `json:"username"`
	Cards []string `json:"cards"`
	SocketConnection *websocket.Conn

}

// NewController example
func NewController() *World {
	return &World{}
}

func (world *World) StartGame(ctx *gin.Context){ 

	

	}

func (world *World) GetGames(ctx *gin.Context){


	ctx.JSON(200, gin.H{
		"status":  "success",
		"details":world,
	})


}



func (world *World) CreateGame(ctx *gin.Context){

	game_code := utils.GenerateGameCode(4)
	world.Games[game_code] = &Match{}
	
	
	ctx.JSON(200, gin.H{
		"status":  "success",
		"game_code":game_code,
		"details":world.Games[game_code],
	})

	
	// //Generating the cards for the game
	// suits :=[...]string{"S", "D", "C", "H"}
	// values :=[...]string{"A", "2", "3", "4", "5", "6", "7", "9", "10", "J", "Q", "K"}

	// var deck [] string

	// for _, suit := range suits {
	// 	for _, value := range values {
	// 		deck = append(deck, (suit+value))
	// 	}
	// }
	// deck= utils.Shuffle(deck)

	// match.P1.Cards=deck[0:7]
	// match.P2.Cards=deck[8:15]
	// match.P3.Cards=deck[16:23]
	// match.P4.Cards=deck[24:31]
	// match.P5.Cards=deck[32:39]
	// match.P6.Cards=deck[40:47]

}

func (world *World) JoinGame(ctx *gin.Context){

	//Bind request body to i
	var i models.JoinGamePayload
	ctx.BindJSON(&i)

		
	if world.Games[i.GameCode].Players == nil{

		world.Games[i.GameCode].Players = make([]*Player,0)

	}
	
	//Use i to add user to the game
	world.Games[i.GameCode].Players  = append(world.Games[i.GameCode].Players,&Player{ UserName:i.UserName})


	ctx.JSON(200, gin.H{
		"status":  "success",
		"username":i.UserName,
	})



}