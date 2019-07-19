package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func ManageErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type User struct {
	Name       string
	Age        int
	Profession string
	Adress     string
}

var (
	//create a slice of 10 elements to save all users
	Users             = make([]User, 3)
	a                 User
	nam, adr, profess string
	age               int
)

func home(w http.ResponseWriter, r *http.Request) {

	//we calling the template as follow
	t, err := template.ParseFiles("names.html")
	ManageErr(err)

	//wuth the next line we pass all the object.(struct)
	t.Execute(w, &Users)
}

//GenerateUsers is used to generate user's data
func GenerateUsers() {
	for i := 0; i < 3; i++ {
		//we ask for the data of every single user in console. and
		//we create 10 Users
		fmt.Println("Give me the name of the user.")
		fmt.Scan(&nam)
		fmt.Println("Give me the Age of the user.")
		fmt.Scan(&age)
		fmt.Println("Give me the Profession of the user.")
		fmt.Scan(&profess)
		fmt.Println("Give me the Adress of the user.")
		fmt.Scan(&adr)
		a.New(nam, profess, adr, age)
		Users[i] = a
	}
}

func (u *User) New(nam, profess, adr string, age int) {
	u.Name = nam
	u.Age = age
	u.Profession = profess
	u.Adress = adr
}

func main() {
	//we firts Generate data
	GenerateUsers()
	//we create a handlermux
	mux := http.NewServeMux()
	//adding the first handlermux
	mux.HandleFunc("/", home)
	//with the follow two lines we are creating the Server
	Server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	Server.ListenAndServe()
}
