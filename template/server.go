''package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Users struct {
	Name   string
	Age    int
	Phone  string
	Adress string
}

func index(w http.ResponseWriter, r *http.Request) {
	//calling a template
	t, err := template.ParseFiles("test.html")
	if err != nil {
		fmt.Println(err)
	}
	firstuser := Users{
		Name:   "Juan",
		Age:    29,
		Phone:  "56123489",
		Adress: "Dinamarca",
	}
	t.Execute(w, firstuser)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
