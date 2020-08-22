package models
import (
"github.com/gorilla/websocket"
)
type JoinGamePayload struct{

	UserName string `json:"userName" binding:"required"`
	GameCode string `json:"gameCode" binding:"required"`

}

type StartGamePayload struct{

	GameCode string `json:"gameCode" binding:"required"`

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