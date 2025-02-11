package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/LainForge/Neura-Launch-Dashboard/dashboard/controllers"
	"github.com/LainForge/Neura-Launch-Dashboard/dashboard/initializers"
	"github.com/LainForge/Neura-Launch-Dashboard/dashboard/middlewares"
)

func init() {
	initializers.LoanEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	// Initializing the gin router/engine
	r := gin.Default()

	// Allowing Cross Origin Requests
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Auth URLs
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	// Generate Token
	r.POST("/token", middlewares.RequireAuth, controllers.TokenController)

	// Upload Code to S3
	r.POST("/upload", controllers.UploadFile)

	// Project URLs
	r.POST("/project/new", middlewares.RequireAuth, controllers.CreateNewProject)
	r.GET("/projects", middlewares.RequireAuth, controllers.GetProjects)
	r.GET("/project/:token", middlewares.RequireAuth, controllers.GetProject)
	r.DELETE("/project/:token", middlewares.RequireAuth, controllers.DeleteProject)

	// Ping endpoint
	r.GET("/ping", controllers.Ping)
	r.GET("", controllers.Hello)

	// Starting the Server
	fmt.Println("Server started at port 8080")
	r.Run()

}
