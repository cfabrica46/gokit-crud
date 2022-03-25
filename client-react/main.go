package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	log.Println("Listening on :3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
