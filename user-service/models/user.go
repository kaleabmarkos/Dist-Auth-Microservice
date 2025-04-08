package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID			primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Email		string				`bson:"email" json:"email"`
	Name		string				`bson:"name" json:"name"`
	Avatar		string				`bson:"avatar,omitempty" json:"avatar,omitempty"`
	Role		string				`bson:"role" json:"role"`
	CreateAt	time.Time			`bson:"created_at" json:"created_at"`
}

