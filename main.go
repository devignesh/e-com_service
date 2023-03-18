package main

import (
	routes "e-com/src"
	"e-com/utils/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	r := gin.New()

	setupRoutes(r)
	setupDatabase()
	startServer(port, r)
}

// route setup function
func setupRoutes(r *gin.Engine) {
	routes.ProductRoutes(r)
}

// db setup func
func setupDatabase() {
	database.Db()

}

// startServer to start the server
func startServer(port string, r *gin.Engine) {
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Listen: %s\n", err)
	}

}
