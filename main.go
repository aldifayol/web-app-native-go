package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
}

func main() {	
	// chapter 1

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// chapter 2

	http.HandleFunc("/again", func(w http.ResponseWriter, r *http.Request) {
		data := "again"
		w.Write([]byte(data))
	})

	// chapter 3
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	address := "localhost:8080"
	fmt.Printf("server started at %s\n", address)
	err :=	http.ListenAndServe(address, nil) 
	if err != nil {
		fmt.Println(err.Error())
	}

}