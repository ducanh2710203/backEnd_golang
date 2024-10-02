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

// Đăng ký người dùng (POST /users/register)
func RegisterUser(c *gin.Context) {
	var user models.User
	var collection = config.GetCollection("users")

	// Bind dữ liệu JSON vào struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra xem email đã tồn tại hay chưa
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "Email đã tồn tại"})
		return
	}

	// Tạo ObjectID mới cho người dùng
	user.ID = primitive.NewObjectID()

	// Thêm người dùng vào cơ sở dữ liệu
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Lỗi khi tạo người dùng"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Đăng ký thành công", "user": user})
}

// Đăng nhập người dùng (POST /users/login)
func LoginUser(c *gin.Context) {
	var loginData struct {
		Email string `json:"email" binding:"required"`
		Pass  string `json:"pass" binding:"required"`
	}
	var collection = config.GetCollection("users")

	// Bind dữ liệu JSON vào struct loginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tìm kiếm người dùng theo email
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email hoặc mật khẩu không chính xác"})
		return
	}

	// Kiểm tra mật khẩu
	if user.Pass != loginData.Pass {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email hoặc mật khẩu không chính xác"})
		return
	}

	// Đăng nhập thành công
	c.JSON(http.StatusOK, gin.H{"message": "Đăng nhập thành công", "user": user})
}
