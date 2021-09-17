package mongodb

import (
	"api_test_hexagonal/config"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var _ = godotenv.Load(".env")

func Connect(collection string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.ConnectionUriMongo))
	if err != nil {
		log.Fatal("err",err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("err",err)
	}
	//defer client.Disconnect(ctx)
	return client.Database(os.Getenv("db_name_mongo")).Collection(collection)
}
