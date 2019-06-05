package main

import(
"github.com/gorilla/mux"
"net/http"
"login/common"
)
 var router =mux.NewRouter()

 
func main() {
    router.HandleFunc("/", common.LoginPageHandler) // GET
 
    router.HandleFunc("/index", common.IndexPageHandler) // GET
    router.HandleFunc("/login", common.LoginHandler).Methods("POST")
 
    router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
    router.HandleFunc("/register", common.RegisterHandler).Methods("POST")
 
    router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")
 
    http.Handle("/", router)
    http.ListenAndServe(":8080", nil)
}
