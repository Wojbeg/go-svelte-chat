package controllers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Wojbeg/go-svelte-chat/pkg/database"
	"github.com/Wojbeg/go-svelte-chat/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotValidUser = errors.New("user is trying to delete someone else's message")
)

var MessagesController *mongo.Collection = database.MessagesData(database.Client, "Messages")

func AddMessage(message models.Message) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	defer ctx.Done()

	messTime, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	message.Created_At = messTime

	if err != nil {
		log.Println("Can not insert message to database, error: ", err)
		return err
	}

	_, err = MessagesController.InsertOne(ctx, message)

	if err != nil {
		log.Println("Can not insert message to database, error: ", err)
		return err
	}

	return nil
}

func GetMessages() (messagesList []models.Message, err error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	defer ctx.Done()

	cursor, err := MessagesController.Find(ctx, bson.D{{}})

	if err != nil {
		log.Println("Can not find cursor in messages, error: ", err)
		return nil, err
	}

	err = cursor.All(ctx, &messagesList)

	if err != nil {
		log.Println("Can not find any messages, error: ", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.Err(); err != nil {
		log.Println("The cursor has encountered this problem, error: ", err)
		return nil, err
	}

	return
}

//Update and delete functions are not used yet. In the next version it is planned to add new functionalities related to this

func DeleteMessage(id primitive.ObjectID, userId string) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	defer ctx.Done()

	var message models.Message
	err = MessagesController.FindOne(ctx, bson.M{"_id": id}).Decode(&message)

	if err != nil {
		log.Printf("Can't find message with id: %+v, error: %+v \n", id, err.Error())
		return
	}

	if message.UserID != userId {
		log.Printf("User is trying to delete someone else's message, invalid id: %+v\n", message.UserID)
		err = ErrNotValidUser
		return
	}

	_, err = MessagesController.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		log.Printf("Can't delete message with id: %+v, error: %+v \n", id, err.Error())
		return
	}

	return
}

//Update and delete functions are not used yet. In the next version it is planned to add new functionalities related to this

func UpdateMessage(updatedMessage models.Message) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	defer ctx.Done()

	var message models.Message
	err := MessagesController.FindOne(ctx, bson.M{"_id": updatedMessage.ID}).Decode(&message)

	if err != nil {
		log.Printf("Can't find message with id: %+v, error: %+v \n", updatedMessage.ID, err.Error())
		return err
	}

	message.Type = updatedMessage.Type
	message.Body = updatedMessage.Body

	_, err = MessagesController.UpdateOne(ctx, bson.M{"_id": updatedMessage.ID}, &message)

	if err != nil {
		log.Printf("Can't update message with id: %+v, error: %+v \n", updatedMessage.ID, err.Error())
		return err
	}

	return nil
}
