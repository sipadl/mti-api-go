package mainuser

import (
	models "mti-cm-be-vendor/models/mainUser"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}

func GetUsers(c *gin.Context) {
	var users []models.User
	DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
