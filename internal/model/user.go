package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id"`
	TelegramID int64         `bson:"telegram_id"`
	Phone      string        `bson:"phone"`
	Username   string        `bson:"username"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at"`
}
