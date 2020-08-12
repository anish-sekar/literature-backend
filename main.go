package main

import (

	"github.com/anish-sekar/literature-backend/utils"
	//"flag"
	//"html/template"
	//"log"
	//"net/http"
    //"errors"
    "fmt"
    //"math/rand"
    //"time"

	//"github.com/gorilla/websocket"
)

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

}