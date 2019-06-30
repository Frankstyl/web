package main

import(
		"net/http"
)

func hi(w http.ResponseWriter, r *http.Request){
	 t:="hello"
	 w.Write([]byte(t))
}	


func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/hi",hi)
	
	server:=&http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}