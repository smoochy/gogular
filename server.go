package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"runtime"
)

var (
	router = mux.NewRouter()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

	router.HandleFunc("/", GetIndex)
	router.HandleFunc("/api/about", GetAbout)
	router.HandleFunc("/api/persons", GetAllPersons).Methods("GET")
	router.HandleFunc("/api/persons", CreatePerson).Methods("POST")
	router.HandleFunc("/api/persons/{id:[0-9]+}", GetPerson)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./angapp/app")))

	http.Handle("/", router)

	CreateDBTables()
	// uncomment and run to create Remy
	// CreatePerson()

	fmt.Println("Server is running and listening at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type Person struct {
	Id        int64
	FirstName string `sql:"size:50"`
	LastName  string `sql:"size:50"`
}

func CreateDBTables() {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/gogular")
	if err != nil {
		panic(err)
	}
	db.CreateTable(Person{})
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "angapp/app/index.html")
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	obj := map[string]string{}
	obj["aboutus"] = "Here is some about us data that is coming from an API call."
	someJson, _ := json.Marshal(obj)
	w.Header().Set("Content-Type", "application/json")
	w.Write(someJson)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("wo")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	fmt.Printf(first_name)
	fmt.Printf(last_name)

	person := Person{FirstName: first_name, LastName: last_name}

	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	db.Save(&person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	params := mux.Vars(r)
	id := params["id"]

	person := Person{}
	my_person := db.First(&person, id)

	someJson, _ := json.Marshal(my_person)
	w.Header().Set("Content-Type", "application/json")
	w.Write(someJson)
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	persons := []Person{}
	people := db.Find(&persons)

	// obj := map[string]string{}
	// obj["people"] = "duh person or something"
	someJson, _ := json.Marshal(people)
	w.Header().Set("Content-Type", "application/json")
	w.Write(someJson)
}
