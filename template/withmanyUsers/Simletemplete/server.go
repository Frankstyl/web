package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

//ManageErr the next function is used to manage the error
func ManageErr(err error) {
	fmt.Println(err)
}

//Data with the following struct we build the data that will be used in the template
type Data struct {
	Title   string
	Content string
}

//HelloTemplate is used to write a hello world in a web browser
func HelloTemplate(w http.ResponseWriter, r *http.Request) {
	//to print all the header of the url fmt.Println(r)
	//to convert the url to strings
	t := r.URL.String()
	//we replacing the / with "" wherever it is found
	t = strings.Replace(t, "/", "", -1)
	//we only print the url in console each time the HandleFunc HelloTemplate is called
	fmt.Println("the url is " + t)
	D := Data{
		Title: t,
	}
	//urlwords[0] = strings.Replace(urlwords[1], " ", "", -1)
	switch t {
	case "helloWorld", "hello", "HelloWorld", "hi", "HI", "HELLOWORLD":
		//we serve the hello World template.
		templ := template.Must(template.ParseFiles("helloWorld.html"))
		mesage := "We are building the future with Gopher"
		D.Content = mesage
		templ.Execute(w, D)
		// this is used for know what is the url fmt.Println("the url is " + t)

	case "about", "About", "ABOUT", "aBout", "abOut":
		templ := template.Must(template.ParseFiles("about.html"))
		//data for about
		mesage := "we are building about now :=)"
		D.Content = mesage
		templ.Execute(w, D)
		// this is used for know what is the url fmt.Println("the url is " + t)

	case "info", "INFORMATION", "information", "INFO", "":
		templ := template.Must(template.ParseFiles("info.html"))
		mesage := "So Sorry We are building this place for you"
		D.Content = mesage
		templ.Execute(w, D)
	}

}

//DinamicPhat is used to response any request.
func DinamicPhat(w http.ResponseWriter, r *http.Request) {
	//we save the name's file requested
	name := r.URL.Path[len("/home/"):]
	fmt.Println(r.URL.Path)
	w.Write([]byte(name))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/info", HelloTemplate)
	mux.HandleFunc("/hello", HelloTemplate)
	mux.HandleFunc("/about", HelloTemplate)
	mux.HandleFunc("/home/", DinamicPhat)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
