package app
 
import (
    "github.com/gin-gonic/gin"
) 

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Примеры маршрутов
    r.GET("/api/hello", HelloHandler)
    r.POST("/api/user", CreateUserHandler)
    r.GET("/api/user/:id", GetUserHandler)
    r.PUT("/api/user/:id", UpdateUserHandler)
    r.DELETE("/api/user/:id", DeleteUserHandler)

    return r
}

func HelloHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Привет, мир!"})
}

func CreateUserHandler(c *gin.Context) {
    // Обработка создания пользователя
}

func GetUserHandler(c *gin.Context) {
    // Обработка получения пользователя по ID
}

func UpdateUserHandler(c *gin.Context) {
    // Обработка обновления пользователя по ID
}

func DeleteUserHandler(c *gin.Context) {
    // Обработка удаления пользователя по ID
}
