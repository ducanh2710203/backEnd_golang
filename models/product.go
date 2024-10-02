package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name"`
	Img    string             `bson:"img" json:"img"`
	State  bool               `bson:"state" json:"state"`                       // true hoặc false
	TaskID primitive.ObjectID `bson:"taskId,omitempty" json:"taskId,omitempty"` // Liên kết tới bảng Task
}
