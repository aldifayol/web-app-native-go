package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type M map[string]interface{}

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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Iron Man"}
		var tmpl = template.Must(template.ParseFiles("views/index.html", "views/_header.html", "views/_message.html"))
		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError )
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Tony Stark"}
		var tmpl = template.Must(template.ParseFiles("views/about.html", "views/_header.html", "views/_message.html"))
		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError )
		}
	})


	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/hello", handlerHello)

	address := ":8080"
	fmt.Printf("server started at %s\n", address)
	http.ListenAndServe(address, nil) 

}