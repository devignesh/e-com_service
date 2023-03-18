package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var mongoClient *mongo.Client

func Db() *mongo.Client {
	once.Do(func() {

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error loading .env file")
		}

		mongoDb := os.Getenv("MONGODB_URL")

		//creating context
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(mongoDb)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			fmt.Println("Error connecting to mongodb: ", err)
		}

		mongoClient = client
		fmt.Println("Successfully established connection to mongodb!!")

	})
	return mongoClient
}
