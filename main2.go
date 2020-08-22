package main

import (

	"github.com/anish-sekar/literature-backend/utils"
	//"github.com/anish-sekar/literature-backend/models"
	"flag"
	//"html/template"
	"log"
	"net/http"
    //"errors"
    "fmt"
    //"math/rand"
	//"time"
	//"reflect"
	//"encoding/json"
	"github.com/gorilla/websocket"
)


type Game struct{

	ActiveConnections []*websocket.Conn

}

//var abs_state models.AbsoluteState

func smain() {




gameObject := Game{}

//Generating the cards for the game
suits :=[...]string{"S", "D", "C", "H"}
values :=[...]string{"A", "2", "3", "4", "5", "6", "7", "9", "10", "J", "Q", "K"}

var deck [] string

for _, suit := range suits {
    for _, value := range values {
        deck = append(deck, (suit+value))
    }
}
deck= utils.Shuffle(deck)

var players [][] string 

players = append(players,deck[0:7])
players = append(players,deck[8:15])
players = append(players,deck[16:23])
players = append(players,deck[24:31])
players = append(players,deck[32:39])
players = append(players,deck[40:47])

for _, player:= range players{
		fmt.Println(player)
}

ret := utils.GenerateGameCode(10)
fmt.Println(ret)



http.HandleFunc("/", gameObject.connect)
log.Fatal(http.ListenAndServe(*addr, nil))

}

//var abs_state models.AbsoluteState

var addr = flag.String("addr", "localhost:5000", "http service address")

var upgrader = websocket.Upgrader{	CheckOrigin: func(r *http.Request) bool {
	return true
},} 


func (game *Game) connect(w http.ResponseWriter, r *http.Request) {
	
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	game.ActiveConnections = append(game.ActiveConnections,c)

	defer c.Close()
	//c.SetCloseHandler(printCloseMessage)
	log.Print("JUST CONNECTED:")
	log.Println(c.RemoteAddr())


	for {

		var i map[string]interface{}
		c.ReadJSON(&i)
		if err != nil {
			log.Println("read:", err)
			break
		}


		shoutOut(game.ActiveConnections)

	}

}


func shoutOut(c []*websocket.Conn){

	s := `{
		"hello":"padowan"
	  }`
for _,con:= range c{

	con.WriteMessage(1,[]byte(s) )

}

}

