package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/Wojbeg/go-svelte-chat/pkg/controllers"
	"github.com/Wojbeg/go-svelte-chat/pkg/models"
	"github.com/Wojbeg/go-svelte-chat/pkg/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pool struct {
	Register       chan *Client
	Unregister     chan *Client
	Clients        map[*Client]bool
	Broadcast      chan models.Message
	SendManyToUser chan util.Pair[[]models.Message, *Client]
}

func NewPool() *Pool {
	return &Pool{
		Register:       make(chan *Client),
		Unregister:     make(chan *Client),
		Clients:        make(map[*Client]bool),
		Broadcast:      make(chan models.Message),
		SendManyToUser: make(chan util.Pair[[]models.Message, *Client]),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

			messTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			//I decided to use type -1 to send client.ID to user
			message := models.Message{
				ID:         primitive.NewObjectID(),
				Type:       -1,
				Username:   "Server",
				UserID:     client.ID,
				Created_At: messTime,
				Body:       "You have joined a chat",
			}

			client.Conn.WriteJSON(
				message,
			)

			message.Type = 1
			message.Body = "New user joined..."

			controllers.AddMessage(message)

			for allclient, _ := range pool.Clients {
				if allclient.ID != client.ID {
					allclient.Conn.WriteJSON(
						message,
					)
				}
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			messTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			message := models.Message{
				ID:         primitive.NewObjectID(),
				Type:       1,
				UserID:     client.ID,
				Username:   client.Username,
				Created_At: messTime,
				Body:       fmt.Sprintf("User %s disconnected...", client.Username),
			}

			controllers.AddMessage(message)

			for client, _ := range pool.Clients {

				client.Conn.WriteJSON(
					message,
				)
			}
			break
		case message := <-pool.Broadcast:

			messTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			message.Created_At = messTime
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					return
				}
			}
			break
		case pair := <-pool.SendManyToUser:

			pair.Second.Conn.WriteJSON(
				pair.First,
			)

			break
		}
	}
}
