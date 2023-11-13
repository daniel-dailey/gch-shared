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

type Story struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string
	Tasks []*Task
}

type Task struct {
	ID            primitive.ObjectID `bson:"_id"`
	Title         string             `bson:"title"`
	Description   string             `bson:"description"`
	Assignees     []string           `bson:"assignees"`
	PriorityLevel int                `bson:"priority"`
	CreatedAt     int64              `bson:"created_at"`
	Modified      int64              `bson:"modified_at"`
}

func MakeTask(title, description string, assignees []string) *Task {
	t := &Task{
		Title:       title,
		Description: description,
		Assignees:   make([]string, 0),
	}
	return t
}
