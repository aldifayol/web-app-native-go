package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type M map[string]interface{}

// later unused
/*
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "welcome"
	w.Write([]byte(message))
}
*/

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "hello world"
	w.Write([]byte(message))
}

// chapter 4
func handlerHtml(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
		return
	}

	fmt.Println(*tmpl)

	var data = map[string]interface{}{
		"title": "Mastering Go Web",
		"name": "Aegon",
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError )
	}
}

func main() {	
	// chapter 4
	var tmpl, err = template.ParseGlob(("views/*"))
	if err != nil {
		panic(err.Error())
	}

	// chapter 3
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Iron Man"}
		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError )
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Tony Stark"}
		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError )
		}
	})

	// chapter 1

	http.HandleFunc("/", handlerHtml)
	// http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// chapter 2

	http.HandleFunc("/again", func(w http.ResponseWriter, r *http.Request) {
		data := "again"
		w.Write([]byte(data))
	})

	address := ":8080"
	fmt.Printf("server started at %s\n", address)
	http.ListenAndServe(address, nil) 

}