package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// SHA1 hashes using sha1 algorithm
func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Download the picture of the day to cache if does not exist
func downloadPicture() string {
	time := time.Now().UTC().Format("2006-01-02")
	path := "/var/cache/jukka/"
	filename := SHA1(time) + ".jpg"
	full_filename := path + filename
	println("Time is", time)
	println("Filename is", full_filename)
	if !Exists(full_filename) {
		println("File does not exist, downloading")
		url := "https://picsum.photos/600"
		DownloadFile(full_filename, url)
	} else {
		println("File already exists, not downloading.")
	}
	return filename
}

type Item struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

/*
var items = []Item{
	{
		ID:   1,
		Task: "Buy coffee",
	},
	{
		ID:   2,
		Task: "Drink coffee",
	},
}
*/

type IndexData struct {
	Picture string
	Items   []Item
}

var Client = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	println("todo-frontend: getJson", url)
	resp, err := Client.Get(url)
	if err != nil {
		println("todo-frontend: getJson: failed to fetch data")
		log.Panic(err)
	}
	defer resp.Body.Close()
	/*
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println("todo-frontend: getJson: failed to read data")
			log.Panic(err)
		}
		bodyString := string(bodyBytes)
		println("todo-frontend: getJson: resp.body ", bodyString)
	*/
	return json.NewDecoder(resp.Body).Decode(target)
}

const port = ":3000"
const api_url = "http://todo-backend-svc:8000" // "http://localhost:8000"

func index(w http.ResponseWriter, r *http.Request) {
	println("todo-frontend: index")
	pic_file := downloadPicture()
	var items []Item
	getJson(api_url+"/todos", &items)
	tmpl_file := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(tmpl_file)
	data := IndexData{pic_file, items}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func add_todo(w http.ResponseWriter, r *http.Request) {
	println("todo-frontend: add_todo")
	r.ParseForm()
	item := Item{0, r.Form.Get("task")}
	jsonValue, _ := json.Marshal(item)
	http.Post(api_url+"/todos", "application/json", bytes.NewBuffer(jsonValue))
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	fs := http.FileServer(http.Dir("/var/cache/jukka"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/add_todo", add_todo)
	println("Server address: http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
