package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const port = ":8000"

func OpenConnection() *sql.DB {

	pg_host := "pingpong-database-svc"
	pg_port := 5432
	pg_user := "postgres"
	pg_password := os.Getenv("POSTGRES_PASSWORD")
	pg_dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg_host, pg_port, pg_user, pg_password, pg_dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func InitDB() {
	db := OpenConnection()
	db.Exec("CREATE TABLE IF NOT EXISTS PingPongCounter (cnt int);")
	var numberOfRows int
	db.QueryRow("SELECT COUNT(*) FROM PingPongCounter;").Scan(&numberOfRows)
	println("Number of rows in database =", numberOfRows)
	if numberOfRows == 0 {
		db.Exec("INSERT INTO PingPongCounter VALUES (0)")
	}
}

func IncreasePingPongCounter() int {
	db := OpenConnection()
	var cnt int
	db.QueryRow("SELECT cnt FROM PingPongCounter LIMIT 1;").Scan(&cnt)
	db.Exec("UPDATE PingPongCounter SET cnt = $1;", cnt+1)
	return cnt
}

func RegisterPing(w http.ResponseWriter, r *http.Request) {
	cnt := IncreasePingPongCounter()
	println("PingPong called, cnt =", cnt)
	fmt.Fprint(w, cnt)
}

func main() {
	InitDB()
	http.HandleFunc("/", RegisterPing)
	println("Ping/pong server listening in address http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
