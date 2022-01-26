package main

import (
	"html/template"
	"net/http"
)

type User struct {
	name   string
	age    uint16
	hobbie []string
}

func homePage(page http.ResponseWriter, r *http.Request) {
	user1 := User{"Mike", 29, []string{"mast", "coding", "crying of being poor"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, user1)
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}

func main() {

	HandleRequest()
}
