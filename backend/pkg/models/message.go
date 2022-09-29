package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `json:"messageId" bson:"_id"`
	Type       int                `json:"type"`
	UserID     string             `json:"userId"`
	Username   string             `json:"username"`
	Created_At time.Time          `json:"created_at"`
	Body       string             `json:"body"`
}

type MessageRequest struct {
	Username string `json:"username"`
	Body     string `json:"body"`
}
