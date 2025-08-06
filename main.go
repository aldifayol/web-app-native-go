package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address string
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}

type Person struct {
	Name string
	Gender string
	Hobbies []string
	Info Info
}

func main() {	
	// serving static files from assets folder
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name: "Tony Stark",
			Gender: "male",
			Hobbies: []string{"flying", "fvcking"},
			Info: Info{"Stark Industries", "10880 Malibu Point, 90265"},
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