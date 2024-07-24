package main

import (
	"engineering_task/controllers"
	"engineering_task/initializers"
	"engineering_task/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController  
	BookController 		controllers.BookController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.AuthController
	UserRouteController routes.UserRouteController 
	BookRouterController routes.BookController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)

	BookController = controllers.NewBookControllers(initializers.DB)

	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewAuthController(initializers.DB)
	
	BookController = controllers.NewBookControllers(initializers.DB)
	
	UserRouteController = routes.NewRouteUserController(UserController)

	BookRouterController = routes.NewBookController(BookController)

	server = gin.Default()


}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowMethods  =[]string{"POST", "GET", "DELETE", "PUT"}
	corsConfig.AllowHeaders = []string{"Accept", "Accept-Language", "Content-Type"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)

	
	UserRouteController.UserRoute(router)

	BookRouterController.BookRoute(router)
	
	log.Fatal(server.Run(":" + config.ServerPort))
}