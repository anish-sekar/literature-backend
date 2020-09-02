# literature-backend

## Game flow possibilities

User opens game
Uses the /join route to establish socket connection to the game server
  ws = new WebSocket("ws://localhost:5000/join");
Lifecycle methods to handle events in js are [here](https://javascript.info/websocket) 

* Once the user has socket connection established, at any time he can join a game room
* Messages exchange format - JSON
* Important field {"messageType":""}
  * messsageType tells the server what the client wants to do
  * Clients are allowed to do the following:
      Join game
          {
              gameCode: "XVlB",
              userName: "anish#11252",
              messageType: "connect",
          }
      Play move
          {
              messageType: "move",
              target: "anish#11252",
              query: "H4",
          }
      Declare - tbd
    
 At the end of any user event - (play move/declare etc) the server broadcasts the game's updated state  along with a broadcast message.
            
Things to note about the message:
 It will have  a message code/type so the client can handle the message.
 currentTurn field - so UI can highlight the player who is playing the turn (also prevent playing out of turn)
 broadcast field has a message that UI will display on a banner on top
 I'm thinking instead of saying 5 Heart -> I can send the card code and you can display an image of the card. But that's just nice to have.
            








{
                  "messageCode": 0,
                  "messageType": "gameState",
                  "hasStarted": true,
                  "currentTurn": "lakshman",
                  "players": {
                    "anish": {
                      "cards": 6
                    },
                    "sekar": {
                      "cards": 7
                    },
                    "rajesh": {
                      "cards": 7
                    },
                    "hari": {
                      "cards": 7
                    },
                    "kiran": {
                      "cards": 7
                    }
                  },
                  "yourCards": [
                    "D10",
                    "C7",
                    "S2",
                    "H7",
                    "HK",
                    "H4",
                    "D9",
                    "C3"
                  ],
                  "team1": [
                    "anish",
                    "sekar",
                    "hari"
                  ],
                  "team2": [
                    "rajesh",
                    "kiran"
                  ],
                  "Broadcast": "hari took C3 from anish"
            }
