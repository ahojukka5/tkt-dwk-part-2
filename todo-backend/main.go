package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var db = []Item{
	{
		ID:   1,
		Task: "Buy coffee",
	},
	{
		ID:   2,
		Task: "Drink coffee",
	},
}

var curId = 2

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status": "ok"}`)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	println("todo-backend: getTodos")
	w.Header().Set("Content-Type", "application/json")
	db, err := json.Marshal(db)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, string(db))
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	println("todo-backend: postTodo")
	decoder := json.NewDecoder(r.Body)
	var item Item
	err := decoder.Decode(&item)
	if err != nil {
		panic(err)
	}

	println("New todo item: " + item.Task)
	curId += 1
	item.ID = curId
	db = append(db, item)

	resp, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, string(resp))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", postTodo).Methods("POST")
	println("todo-backend listening on port 8000")
	http.ListenAndServe(":8000", router)
}
