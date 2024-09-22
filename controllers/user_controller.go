package controllers

import (
	"context"
	"net/http"

	"go-backend/config"
	"go-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tạo user mới (POST /users)
func CreateUser(c *gin.Context) {
	var collection = config.GetCollection("users")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = primitive.NewObjectID()
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Đã thêm user", "userID": insertResult.InsertedID})
}

// Lấy danh sách tất cả users (GET /users)
func GetUsers(c *gin.Context) {
	var collection = config.GetCollection("users")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể truy xuất danh sách users"})
		return
	}
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi xử lý dữ liệu"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Lấy thông tin user theo ID (GET /users/:id)
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var user models.User
	var collection = config.GetCollection("users")

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Cập nhật user theo ID (PUT /users/:id)
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		},
	}
	var collection = config.GetCollection("users")

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể cập nhật user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User đã được cập nhật"})
}

// Xóa user theo ID (DELETE /users/:id)
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}
	var collection = config.GetCollection("users")

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User đã được xóa"})
}
