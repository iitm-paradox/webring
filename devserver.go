package main


import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
