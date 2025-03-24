package device

import "github.com/gin-gonic/gin"

// Пример структуры для устройства
type Device struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Пример функции для регистрации маршрутов (если необходимо)
func RegisterRoutes(router *gin.RouterGroup) {
	// Пример регистрации маршрутов
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Devices list"})
	})
}
