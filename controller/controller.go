package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/utils"
	"github.com/anish-sekar/literature-backend/models"
)
// Controller example
type World struct {

Games map[string]*models.Match

}

// NewController example
func NewController() *World {
	return &World{}
}

func (world *World) StartGame(ctx *gin.Context){ 

		//Bind request body to i
		var i models.StartGamePayload
		ctx.BindJSON(&i)
		
		//Shuffle the players
		world.Games[i.GameCode].Players = utils.ShufflePlayers(world.Games[i.GameCode].Players)
		//Assign to Team1 or Team2
		for indx,value := range world.Games[i.GameCode].Players{

			if indx % 2 == 0{
				world.Games[i.GameCode].Team1 = append(world.Games[i.GameCode].Team1,indx)
			}else{
				world.Games[i.GameCode].Team1 = append(world.Games[i.GameCode].Team2,indx)
			}

		}
		// //Generating the cards for the game
		suits :=[...]string{"S", "D", "C", "H"}
		values :=[...]string{"A", "2", "3", "4", "5", "6", "7", "9", "10", "J", "Q", "K"}

		var deck [] string

		for _, suit := range suits {
			for _, value := range values {
				deck = append(deck, (suit+value))
			}
		}
		deck= utils.ShuffleCards(deck)

		playerCount := len(world.Games[i.GameCode].Players)

		ite := 0
		for _,value := range deck{
			if ite > (playerCount-1) {ite=0}
			world.Games[i.GameCode].Players[ite] = append(world.Games[i.GameCode].Players[ite],value)
			ite++
		}
		// match.P1.Cards=deck[0:7]
		// match.P2.Cards=deck[8:15]
		// match.P3.Cards=deck[16:23]
		// match.P4.Cards=deck[24:31]
		// match.P5.Cards=deck[32:39]
		// match.P6.Cards=deck[40:47]



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