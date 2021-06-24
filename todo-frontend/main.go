package main

import (
	"crypto/sha1"
	"encoding/hex"
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

func index(w http.ResponseWriter, r *http.Request) {
	pic_file := downloadPicture()
	tmpl_file := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(tmpl_file)
	tmpl.ExecuteTemplate(w, "index.html", pic_file)
}

func main() {

	fs := http.FileServer(http.Dir("/var/cache/jukka"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)

	port := ":3000"
	println("Server address: http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
