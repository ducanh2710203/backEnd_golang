package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	TaskID primitive.ObjectID `bson:"taskId,omitempty" json:"taskId,omitempty"` // Liên kết tới bảng Task
}
