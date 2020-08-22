package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/controller"
	"fmt"
)

func main() {

	w:= controller.NewController()
	w.Games= make(map[string]*controller.Match)
	

	fmt.Print(w)
	r := gin.Default()


	r.GET("/create", w.CreateGame)
	r.GET("/games",w.GetGames)
	r.POST("start".w.StartGame)
	r.PUT("/join",w.JoinGame)

	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}