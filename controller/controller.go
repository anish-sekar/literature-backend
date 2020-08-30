package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/utils"
	"github.com/anish-sekar/literature-backend/models"
	"net/http"
	"log"
)


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

		//Check if theere are minimum number of players
		
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
		world.Games[i.GameCode].CurrentTurn=world.Games[i.GameCode].Team1[0]

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

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//Request to upgrade the socket connection
	connectionToClient, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	defer connectionToClient.Close()
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
				ctx.JSON(418, gin.H{
					"status":  "failure",
					"error":err,
				})
		}
		return
	}else 
		{

			// Listen for messages from the client 
			for {
				if err := connectionToClient.ReadJSON(&input); err != nil {
					if websocket.IsCloseError(err,1000){
						log.Println(connectionToClient.RemoteAddr(),"left voluntarily")
						//Notify that #player left the game
						if username != ""{

							for key,v := range world.Games[gameCode].Players{
								if username != key{
									
									err := v.SocketConnection.WriteJSON(`{"broadcast":"`+username+` left the game"}`);
									if err != nil{
										log.Println(username + " left the game| reason:",err )
									}else{
									}
								}	
							}
						}

					}
					break
				}else{
					
					// Switch to handle different client behaviours
					switch {
							
						case input["messageType"].(string) == "connect":
								//Need to add error handling to check for empty inputs from the client
								username = input["userName"].(string)
								gameCode = input["gameCode"].(string)
								joinGameRoom(world,username,gameCode,connectionToClient)
								log.Println("User has connected")

						case input["messageType"].(string) == "move":

							target := input["target"].(string)
							query  := input["query"].(string)

							//Check if their move was valid
							if username != world.Games[gameCode].CurrentTurn {

								err := world.Games[gameCode].Players[username].SocketConnection.WriteJSON(`{"messageCode":2,"messageType":"notYourTurn"}`);
								if err != nil{
									log.Println("Error publishing to player",username,"| reason:",err )
								}
								break	
							}
							
							//Perform the move

								//Check if the target has the card
								hit := performQuery(world,gameCode,username,target,query)
							
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
											MessageCode : 0 ,
											MessageType : "gameState",
											HasStarted : world.Games[gameCode].HasStarted,
											CurrentTurn : world.Games[gameCode].CurrentTurn,
											Players : abstracted_players,
											YourCards : players_cards,
											Team1: world.Games[gameCode].Team1,
											Team2: world.Games[gameCode].Team2,
											Broadcast: username + " took " + query + " from " + target,
			
										}
										//Pass the user state to the appropriate user
										publish(world,player,gameCode,userstate)
									}



								}else{
									for player,_ := range world.Games[gameCode].Players{
										
										
									publishMiss(world,player,gameCode, &models.QueryFailureResponse{MessageCode : 1 ,
										MessageType : "queryFailureResponse",
										CurrentTurn : world.Games[gameCode].CurrentTurn,
										Broadcast: target + " does not have "+ query })
									}
								}


						case input["messageType"].(string) == "declare":

							//Check if it's the players turn

						
						default:
							unknownMessage := `
							{	
								
								"messageType":"unknown messageType",
								"description: "The message"input["messageType"]
							}
							`
							connectionToClient.WriteJSON(unknownMessage)
							

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