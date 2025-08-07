package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
    Name    string
    Alias   string
    Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
    return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {	
	// serving static files from assets folder
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
            Name:    "Tony Stark",
            Alias:   "Iron Man",
            Friends: []string{"Thor", "Captain America", "Superman"},
        }

		var tmpl = template.Must(template.ParseFiles("views/view.html"))
		
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError )
		}
	})

	address := ":8080"
	fmt.Printf("server started at %s\n", address)
	http.ListenAndServe(address, nil) 

}