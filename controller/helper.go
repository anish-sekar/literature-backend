package controller

import (
	"log"
	"github.com/anish-sekar/literature-backend/models"
	//"github.com/anish-sekar/literature-backend/utils"
	"github.com/gorilla/websocket"
)

func joinGameRoom(world *World, username string, gameCode string, sock *websocket.Conn) bool{

	// When the first player is joining, create the Players map. Refer models.Mat struct
	if world.Games[gameCode].Players==nil{
		world.Games[gameCode].Players = make(map[string]*models.Player)
	}

	//Add user to the game 
	if _, ok := world.Games[gameCode].Players[username]; ok {
		// Update the socket information
		log.Println("Someone with the same username is already in the game, updating the socket conn.")
		world.Games[gameCode].Players[username].SocketConnection = sock
	}else{
		//Create new player with socket information
		world.Games[gameCode].Players[username] = &models.Player{SocketConnection:sock}
	}

	return true
}

func publish(world *World, targetUser string, gameCode string,gameState *models.UserState) bool{ 

		err := world.Games[gameCode].Players[targetUser].SocketConnection.WriteJSON(gameState);
		if err != nil{
			log.Println("Error publishing to player",targetUser,"| reason:",err )
			return false
		}else{
			return true
		}

}

func publishMiss(world *World, targetUser string, gameCode string, queryFailResponse *models.QueryFailureResponse) bool{ 

			
	err := world.Games[gameCode].Players[targetUser].SocketConnection.WriteJSON(queryFailResponse);
	if err != nil{
		log.Println("Error publishing to player",targetUser,"| reason:",err )
		return false
	}else{
		return true
	}

}


func performQuery(world *World,gameCode string,source string,target string, query string) bool{

	//Check if the target has the card
	for indx,card := range world.Games[gameCode].Players[target].Cards{
		if query==card{
			world.Games[gameCode].Players[target].Cards = remove(world.Games[gameCode].Players[target].Cards,indx)
			world.Games[gameCode].Players[source].Cards = append(world.Games[gameCode].Players[source].Cards, card)
			world.Games[gameCode].CurrentTurn = source
			return true
		}
	}
	world.Games[gameCode].CurrentTurn = target
	return false
}


func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}
