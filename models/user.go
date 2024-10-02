package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	Email   string             `bson:"email" json:"email"`
	Pass    string             `bson:"pass" json:"pass"`
	GroupID primitive.ObjectID `bson:"groupId,omitempty" json:"groupId,omitempty"` // Liên kết tới bảng Group
	TaskID  primitive.ObjectID `bson:"taskId,omitempty" json:"taskId,omitempty"`   // Liên kết tới bảng Task
}
