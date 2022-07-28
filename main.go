package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "item 1", Done: true},
			{Item: "item 2", Done: false},
			{Item: "item 3", Done: false},
			{Item: "item 4", Done: false},
			{Item: "item 5", Done: true},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	// tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static")) // File server
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// Request handler
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":9091", mux))
}
