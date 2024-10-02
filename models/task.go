package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Định nghĩa enumTaskState dưới dạng kiểu dữ liệu tùy chỉnh
type enumTaskState string

const (
	StatePending   enumTaskState = "Pending"
	StateCompleted enumTaskState = "Completed"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID primitive.ObjectID `bson:"productId,omitempty" json:"productId,omitempty"` // Liên kết tới bảng Product
	UserID    primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`       // Liên kết tới bảng User
	GroupID   primitive.ObjectID `bson:"groupId,omitempty" json:"groupId,omitempty"`     // Liên kết tới bảng Group
	StartTime time.Time          `bson:"startTime" json:"startTime"`
	EndTime   time.Time          `bson:"endTime" json:"endTime"`
	Status    enumTaskState      `bson:"status" json:"status"`
}
