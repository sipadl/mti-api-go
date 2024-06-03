package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	User "mti-cm-be-vendor/controllers/MainUser"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.New()
	User.Init(db)

	r.GET("/user", User.GetUsers)
	return r
}
