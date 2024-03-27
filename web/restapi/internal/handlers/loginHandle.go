package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func AuthGameRoute(r *gin.Engine, authGroup *gin.RouterGroup) {

	{
		authGroup.POST("/login", func(c *gin.Context) {
			var loginInfo LoginInfo
			if err := c.ShouldBindJSON(&loginInfo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if loginInfo.Password == "admin" {
				c.JSON(http.StatusOK, gin.H{"status": "раздача на спавне!"})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": "вам выдали бан по железу!"})
			}
		})

		authGroup.POST("/register", func(c *gin.Context) {
			var registerInfo RegisterInfo
			if err := c.ShouldBindJSON(&registerInfo); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// Обработка данных registerInfo...
			c.JSON(http.StatusOK, gin.H{"status": "Вы успешно зарегистрировались!"})
		})
	}

	// Группа путей для игры, для примера
	gameGroup := r.Group("/game")
	{
		gameGroup.GET("/info", func(c *gin.Context) {
			// Пример обработки запроса
			c.JSON(http.StatusOK, gin.H{"game": "Информация об игре"})
		})
	}
}
