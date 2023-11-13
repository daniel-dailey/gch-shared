package entities

import (
	"github.com/daniel-dailey/gch-shared/packets"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskState int

const (
	TaskStateIdle TaskState = iota
	TaskStateInProgress
	TaskStateDone
	TaskStateParked
)

func (ts TaskState) String() string {
	switch ts {
	case TaskStateIdle:
		return "Idle"
	case TaskStateInProgress:
		return "In Progress"
	case TaskStateDone:
		return "Done"
	case TaskStateParked:
		return "Parked"
	default:
		return ""
	}
}

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
	ID            primitive.ObjectID       `bson:"_id"`
	Title         string                   `bson:"title"`
	Description   string                   `bson:"description"`
	Assignees     []string                 `bson:"assignees"`
	PriorityLevel int                      `bson:"priority"`
	CreatedAt     int64                    `bson:"created_at"`
	Modified      int64                    `bson:"modified_at"`
	Status        TaskState                `bson:"state"`
	Comments      []*packets.MessagePacket `bson:"comments"`
}

func MakeTask(title, description string, assignees []string) *Task {
	t := &Task{
		Title:       title,
		Description: description,
		Assignees:   make([]string, 0),
		Comments:    make([]*packets.MessagePacket, 0),
	}
	return t
}
