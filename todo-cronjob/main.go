package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var collection = ConnectDB()

type Item struct {
	ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Task string             `json:"task"`
}

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getConnectionURI() string {
	username := Getenv("MONGO_USERNAME", "root")
	password := Getenv("MONGO_PASSWORD", "")
	host := Getenv("MONGO_HOST", "todo-database-svc")
	mongo_uri := "mongodb://" + username + ":" + password + "@" + host
	return mongo_uri
}

func ConnectDB() *mongo.Collection {
	println(getConnectionURI())
	clientOptions := options.Client().ApplyURI(getConnectionURI())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected to MongoDB!")
	db := client.Database("todo")
	col := db.Collection("items")
	return col
}

func main() {

	resp, err := http.Get("https://en.wikipedia.org/wiki/Special:Random")
	if err != nil {
		log.Fatal(err)
		return
	}

	url := "https://" + resp.Request.URL.Host + "/" + resp.Request.URL.Path

	println("Random url:", url)

	var item Item
	item.ID = primitive.NewObjectID()
	item.Task = "Read wikipedia link " + url

	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		fmt.Println("Failed to insert to database:", err)
		return
	}
	println("Result", result)
}
