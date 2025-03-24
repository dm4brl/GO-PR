package user

import "github.com/gin-gonic/gin"

type User struct {
	ID     string
	Name   string
	GeoLat float64
	GeoLng float64
}

type Repository interface {
	GetByID(id string) (*User, error)
	Create(user *User) error
}

// RegisterRoutes регистрирует маршруты для работы с пользователями
func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Users list"})
	})
}
