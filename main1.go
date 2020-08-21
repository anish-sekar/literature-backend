package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/anish-sekar/literature-backend/controller"
)

func main() {

	c:= controller.NewController()

	r := gin.Default()


	r.GET("/create", c.CreateGame)

	r.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}