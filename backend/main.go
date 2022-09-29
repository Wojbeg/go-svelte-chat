package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Wojbeg/go-svelte-chat/pkg/controllers"
	"github.com/Wojbeg/go-svelte-chat/pkg/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		serveWS(pool, w, r)
	})
}

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err.Error())
	}

	client := &websocket.Client{
		ID:   primitive.NewObjectID().Hex(),
		Conn: conn,
		Pool: pool,
	}

	//Register user
	pool.Register <- client

	//Fetch data from db
	messages, err := controllers.GetMessages()

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err.Error())
		log.Println(err)
	} else if messages != nil {
		client.SendChatMessanges(messages)
	}

	client.Read()
}

func main() {
	fmt.Println("Wojbeg's websockets chat app started")
	setupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	fmt.Println("Listening port ", port)
	http.ListenAndServe(port, nil)
}
