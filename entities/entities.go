package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
	Birthday string             `bson:"birthday"`
	Bio      string             `bson:"bio"`
}
