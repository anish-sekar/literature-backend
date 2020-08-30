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

	HasStarted bool
	CurrentTurn string
	Players map[string]*Player 
	Team1 []string 
	Team2 []string
	
	}
	
type Player struct {
	
	Cards []string `json:"cards"`
	SocketConnection *websocket.Conn

}
type AbstractedPlayer struct {


	Cards  int `json:"cards"`

}


type UserState struct{
	MessageCode int `json:"messageCode"`
	MessageType string `json:"messageType"`
	HasStarted bool `json:"hasStarted"`
	CurrentTurn string `json:"currentTurn"`
	Players map[string]*AbstractedPlayer `json:"players"`
	YourCards []string `json:"yourCards"`
	Team1 []string `json:"team1"`
	Team2 []string `json:"team2"`
	Broadcast string

}

type QueryFailureResponse struct {
	MessageCode int `json:"messageCode"`
	MessageType string `json:"messageType"`
	CurrentTurn string `json:"currentTurn"`
	Broadcast string `json:"broadcast"`
}