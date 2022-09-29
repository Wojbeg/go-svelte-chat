package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Wojbeg/go-svelte-chat/pkg/controllers"
	"github.com/Wojbeg/go-svelte-chat/pkg/models"
	"github.com/Wojbeg/go-svelte-chat/pkg/util"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID       string
	Username string
	Conn     *websocket.Conn
	Pool     *Pool
	mu       sync.Mutex
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		var body string

		var request models.MessageRequest
		err = json.Unmarshal(p, &request)

		if err != nil {
			log.Println("error: ", err)
			body = string(p)
		} else {
			body = request.Body
		}

		if c.Username == "" {
			c.Username = request.Username
		}

		message := models.Message{
			ID:       primitive.NewObjectID(),
			Type:     messageType,
			Username: c.Username,
			UserID:   c.ID,
			Body:     body,
		}

		//Save message
		controllers.AddMessage(message)

		c.Pool.Broadcast <- message
		log.Printf("Message recived: %+v by user %+v \n", message, c.ID)
	}
}

func (c *Client) SendChatMessanges(messages []models.Message) {

	infoToSend := util.Pair[[]models.Message, *Client]{
		First:  messages,
		Second: c,
	}
	c.Pool.SendManyToUser <- infoToSend
}
