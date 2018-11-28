package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handlePost(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	t := time.Now()
	timestamp := t.Format("Mon 02 Jan 2006 15:04:05 MST")

	logEntry := fmt.Sprintf("[%s] %s", timestamp, body)
	fmt.Fprint(rw, logEntry)
	log.Println(logEntry)
}

func main() {
	log.Print("Starting server on port 8080...")
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
