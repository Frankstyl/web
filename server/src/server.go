package main

import(
		"net/http"
		"fmt"
)

func hi(w http.ResponseWriter, r *http.Request){
	 t:="hello"
	 w.Write([]byte(t))
}	
func helloWorld(w http.ResponseWriter, r *http.Request){
word:="hello World my name is Gopher"
fmt.Fprintf(w,word)
}


func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/hi",hi)
	mux.HandleFunc("/hello",helloWorld)
	
	server:=&http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}