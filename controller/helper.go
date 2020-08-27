package controller

import (
	"log"
	"github.com/anish-sekar/literature-backend/models"
	//"github.com/anish-sekar/literature-backend/utils"
	"github.com/gorilla/websocket"
)

func connectUser(world *World, username string, gameCode string, sock *websocket.Conn) bool{

	//Check if game has players & add to list
	// if val, ok := world.Games[gameCode].Players[username]; ok {
	// 	//do something here
	// 	world.Games[gameCode].Players = make(map[string]*utils.Player)
	// }
	if world.Games[gameCode].Players==nil{
		world.Games[gameCode].Players = make(map[string]*models.Player)
	}

	//Add user to the game with socket info
	if _, ok := world.Games[gameCode].Players[username]; ok {
		log.Println("Someone with the same username is already in the game, updating the socket conn.")
		world.Games[gameCode].Players[username].SocketConnection = sock
	}else{
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

func publishMiss(world *World, targetUser string, gameCode string,message string) bool{ 

			
	err := world.Games[gameCode].Players[targetUser].SocketConnection.WriteJSON(message);
	if err != nil{
		log.Println("Error publishing to player",targetUser,"| reason:",err )
		return false
	}else{
		return true
	}

}




func performQuery(world *World,gameCode string,target string, query string) bool{

	//Check if the target has the card
	for indx,card := range world.Games[gameCode].Players[target].Cards{
		if query==card{
			world.Games[gameCode].Players[target].Cards = remove(world.Games[gameCode].Players[target].Cards,indx)
		return true
		}
	}
	return false
}


func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}