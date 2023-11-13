package main

import (
	"go-html-test/controllers"
	"go-html-test/initializers"
	"go-html-test/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// действия до основной функции
func init() {
	initializers.LoadEnvVaribles()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// точка входа
func main() {
	//запуск web сервера
	r := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	//main page
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "main page",
		})
	})

	//controllers
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", middleware.RequireAuth, controllers.Logout)

	r.Run() // listen and serve on 0.0.0.0:env
}
