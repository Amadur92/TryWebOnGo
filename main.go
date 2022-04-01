package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    uint16
	Hobbie []string
}

func homePage(page http.ResponseWriter, r *http.Request) {
	user1 := User{"Person", 44, []string{"asd", "da", "asing poor"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, user1)
}

func mishka(page http.ResponseWriter, r *http.Request) {
	user1 := User{"Mike", 29, []string{"mast", "codiqweng", "sfgg"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, user1)
}

func HandleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/s/", mishka)

	http.ListenAndServe(":8080", nil)

}

func main() {
	HandleRequest()
}
