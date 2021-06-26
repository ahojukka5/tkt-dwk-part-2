package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const port = ":8000"

var ctx = context.TODO()
var collection = ConnectDB()

type ItemWithoutID struct {
	Task string `json:"task"`
}

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
	}
	fmt.Println("Connected to MongoDB!")
	db := client.Database("todo")
	col := db.Collection("items")
	return col
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status": "ok"}`)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	println("getTodos")
	w.Header().Set("Content-Type", "application/json")
	var items []Item
	println("getTodos: fetching collection")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cur.Close(ctx)
	println("getTodos: looping collection")
	for cur.Next(ctx) {
		var item Item
		err := cur.Decode(&item)
		if err != nil {
			fmt.Println(err)
			return
		}
		items = append(items, item)
	}
	println("getTodos: encoding to json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	println("postTodo")
	w.Header().Set("Content-Type", "application/json")
	var item Item
	var itemWithoutID ItemWithoutID
	err := json.NewDecoder(r.Body).Decode(&itemWithoutID)
	if err != nil {
		fmt.Println("Failed to parse input:", err)
		return
	}
	item.ID = primitive.NewObjectID()
	item.Task = itemWithoutID.Task
	println("postTodo: new todo item: " + item.Task)
	if len(item.Task) > 140 {
		println("postTodo: message is too long, over 140 characters!")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		fmt.Println("Failed to insert to database:", err)
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Println("Failed to encode to json:", err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", postTodo).Methods("POST")
	println("Server listening in address http://localhost" + port)
	http.ListenAndServe(port, router)
}
