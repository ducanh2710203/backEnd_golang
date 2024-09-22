package main

import (
	"fmt"
	"go-backend/config"
	routes "go-backend/routers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Chờ cho đến khi ConnectDB() hoàn tất hoặc gặp lỗi
	config.ConnectDB()

	// Chờ cho tới khi ConnectDB hoàn thành
	time.Sleep(10 * time.Second)
	// Sau khi kết nối thành công, tiếp tục khởi tạo router và đăng ký routes
	fmt.Println("Kết nối MongoDB thành công, tiếp tục khởi tạo server...")

	r := gin.Default()

	// Đăng ký các routes
	routes.UserRoutes(r)

	// Chạy server trên cổng mặc định 8080
	r.Run() // Chạy ứng dụng
}
