package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/utils"
	"github.com/anish-sekar/literature-backend/models"
	"net/http"
	"log"
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

		//Check if game has started
		if world.Games[i.GameCode].HasStarted == true {
			ctx.JSON(200, gin.H{
				"status":  "failure",
				"details":"Game has already started",
			})
		}
		
		//Set the game state as started
		world.Games[i.GameCode].HasStarted = true 

		//Shuffle the players
		players := make([]string, 0, len(world.Games[i.GameCode].Players))
		for k := range world.Games[i.GameCode].Players {
			players = append(players, k)
		}

		//Assign to Team1 or Team2
		for indx,player := range players{

			if indx % 2 == 0{
				world.Games[i.GameCode].Team1 = append(world.Games[i.GameCode].Team1,player)
			}else{
				world.Games[i.GameCode].Team2 = append(world.Games[i.GameCode].Team2,player)
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
		deckLength := len(deck)
		playerCount := len(world.Games[i.GameCode].Players)

		ptr := 0
		for key,_ := range world.Games[i.GameCode].Players{
			world.Games[i.GameCode].Players[key].Cards = make([]string,48)
			world.Games[i.GameCode].Players[key].Cards = deck[ptr:(ptr+(deckLength/playerCount))-1]
			ptr = ptr + (deckLength/playerCount) 
			deckLength = deckLength - (deckLength/playerCount)
			playerCount = playerCount -1
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
	world.Games[game_code] = &models.Match{}
	
	
	ctx.JSON(200, gin.H{
		"status":  "success",
		"game_code":game_code,
		"details":world.Games[game_code],
	})

}

func (world *World) JoinGame(ctx *gin.Context){

	var input map[string]interface{}
	var username string
	var gameCode string


	// //Bind request body to i
	// var i models.JoinGamePayload
	// ctx.BindJSON(&i)

	// //Check if game has started
	// if world.Games[i.GameCode].HasStarted != false  {
	// 	ctx.JSON(200, gin.H{
	// 		"status":  "failure",
	// 		"details":"Game has either started or does not exist",
	// 	})
	// 	return
	// }

	// if world.Games[i.GameCode].Players == nil{

	// 	world.Games[i.GameCode].Players = make([]*models.Player,0)

	// }
	
	// //Use i to add user to the game
	// world.Games[i.GameCode].Players  = append(world.Games[i.GameCode].Players,&models.Player{ UserName:i.UserName})

	//Start up a socket connection
	//var upgrader = websocket.Upgrader{	CheckOrigin: func(r *http.Request) bool {return true}}
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}else 
		{
			//log.Print(c.RemoteAddr(),": Joined")
			//Entire game logic to constantly listen to any inputs the players emit
			
			for {
				if err := c.ReadJSON(&input); err != nil {
					if websocket.IsCloseError(err,1000){
						log.Println(c.RemoteAddr(),"left voluntarily")
					}
					break
				}else{
					
					switch {
					case input["messageType"].(string) == "connect":
							//Need to add error handling to check for empty inputs from the client
							username = input["userName"].(string)
							gameCode = input["gameCode"].(string)
							connectUser(world,username,gameCode,c)
							log.Println("User has connected")

					case input["messageType"].(string) == "move":

						target := input["target"].(string)
						query  := input["query"].(string)
						
						//Perform the move

							//Check if the target has the card
							hit := performQuery(world,gameCode,target,query)
						
							if hit {

								for player,value := range world.Games[gameCode].Players{
							
									abstracted_players := make(map[string]*models.AbstractedPlayer)
									var players_cards []string
									
									for key,v := range world.Games[gameCode].Players{
										if player != key{
											abstracted_players[key] = &models.AbstractedPlayer{Cards : len(v.Cards)} 
										}	else{
											players_cards = value.Cards
										}
									}
		
									//Generate the state for each user
									userstate := &models.UserState{
										HasStarted : world.Games[gameCode].HasStarted,
										CurrentTurn : world.Games[gameCode].CurrentTurn,
										Players : abstracted_players,
										YourCards : players_cards,
										Team1: world.Games[gameCode].Team1,
										Team2: world.Games[gameCode].Team1,
		
									}
									//Pass the user state to the appropriate user
									publish(world,player,gameCode,userstate)
								}



							}else{
								for player,_ := range world.Games[gameCode].Players{
								publishMiss(world,player,gameCode,"MISS!")
								}
							}


					default:

					}
				}

			log.Println(input)
			}
			return
		}


	// // ctx.JSON(200, gin.H{
	// // 	"status":  "success",
	// // 	"username":i.UserName,
	// // })
	// return

}