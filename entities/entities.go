package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID
	Name     string
	Age      int
	Birthday string
	Bio      string
}
