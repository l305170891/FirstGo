package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("hello")

	fmt.Fprintln(w, "hello world.")
}
