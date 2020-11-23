package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	MongoDb *mongo.Database
)

func InitDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))

	if err != nil {
		log.Fatal(err)
	}
	MongoDb = client.Database(os.Getenv("DATA_MONGODB_DATABASE"))
	log.Println("Connected")
}
func GetDb() *mongo.Database {
	return MongoDb
}
