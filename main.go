package main

import (

	"github.com/anish-sekar/literature-backend/utils"
	"github.com/anish-sekar/literature-backend/models"
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


var activeConnections []websocket.Conn
var abs_state models.AbsoluteState

func main() {


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




// err := json.Unmarshal([]byte(s), &abs_state)
// if err != nil {
// 	log.Println("read:", err)
// }

// fmt.Println(abs_state.State.P1)



http.HandleFunc("/", echo)
log.Fatal(http.ListenAndServe(*addr, nil))

}

//var abs_state models.AbsoluteState

var addr = flag.String("addr", "localhost:8081", "http service address")

var upgrader = websocket.Upgrader{	CheckOrigin: func(r *http.Request) bool {
	return true
},} 


func echo(w http.ResponseWriter, r *http.Request) {
	
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	activeConnections = append(activeConnections,*c)

	defer c.Close()
	//c.SetCloseHandler(printCloseMessage)
	log.Print("JUST CONNECTED:")
	log.Println(c.RemoteAddr())


	for {

		c.ReadJSON(&abs_state)
		if err != nil {
			log.Println("read:", err)
			break
		}
		shoutOut(activeConnections)
		fmt.Println(abs_state)
		

		// mt, message, err := c.ReadMessage()
		// if err != nil {
		// 	log.Println("read:", err)
		// 	break
		// }

		// log.Printf("recv: %s", message)
		// err = c.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}
}


func shoutOut(c []websocket.Conn){

	s := `{
		"id": "XVlBzgbaiC",
		"turn": "anish#2345",
		"state": {
		  "P1": {
			"username": "a",
			"cards": ["H4","SA"]
		  },
		  "P2": {
			"username": "b",
			"cards": ["H4","SA"]
		  },
		  "P3": {
			"username": "c",
			"cards": ["H4","SA"]
		  },
		  "P4": {
			"username": "d",
			"cards": ["H4","SA"]
		  },
		  "P5": {
			"username": "e",
			"cards": ["H4","SA"]
		  },
		  "P6": {
			"username": "e",
			"cards": ["H4","SA"]
		  }
		}
	  }`
for _,con:= range c{

	con.WriteMessage(1,[]byte(s) )

}

}

