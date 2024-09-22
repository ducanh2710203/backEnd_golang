package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Hàm để kết nối đến MongoDB
func ConnectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("lỗi khi kết nối MongoDB: %v", err)
	}

	// Kiểm tra kết nối
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("không thể kết nối đến MongoDB: %v", err)
	}

	fmt.Println("Đã kết nối thành công đến MongoDB!")
	DB = client.Database("mydatabase") // Gán database vào biến toàn cục DB
	return nil
}

// Lấy collection từ database
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database chưa được khởi tạo, cần gọi ConnectDB() trước")
	}
	return DB.Collection(collectionName)
}
