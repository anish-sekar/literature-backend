package models


type JoinGamePayload struct{

	UserName string `json:"userName" binding:"required"`
	GameCode string `json:"gameCode" binding:"required"`

}