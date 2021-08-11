package main

import (
	"github.com/gin-gonic/gin"
	"github.com/web/config"
	"github.com/web/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectDb()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDbConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.GET("/", index)

	r.Run(":3000")
}
