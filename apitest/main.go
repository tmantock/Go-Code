package main

import (
	"log"
	"net/http"

	"encoding/json"

	"strconv"

	"github.com/gorilla/mux"
)

//Person implements a new Person type for reqresenting a person and their information
type Person struct {
	ID        int
	Firstname string
	Lastname  string
	Address   *Address
}

//Address implements a new type that represents a person's address
type Address struct {
	City  string
	State string
}

var people []Person

//GetPersonEndpoint handles the request for a specific person
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//GetPeopleEndpoint handles the request for a all persons
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//CreatePersonEndpoint handles the request for creating a person
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	i, _ := strconv.Atoi(params["id"])
	person.ID = i
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

//DeletePersonEndpoint handles the request for deleting a specific person
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	for index, item := range people {
		if item.ID == i {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	//create a new router
	router := mux.NewRouter()

	people = append(people, Person{ID: 1, Firstname: "Nicholas", Lastname: "Claus", Address: &Address{City: "North Pole", State: "Arctic"}})
	people = append(people, Person{ID: 2, Firstname: "Jessie", Lastname: "Claus", Address: &Address{City: "North Pole", State: "Greenland"}})
	people = append(people, Person{ID: 3, Firstname: "Saint", Lastname: "Krispus", Address: &Address{City: "South Pole", State: "Antartica"}})
	people = append(people, Person{ID: 4, Firstname: "Amy", Lastname: "Krispus", Address: &Address{City: "West Pole", State: "Maui"}})

	//route for serving a list of people
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	//route for serving a specific person
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	//route that for creating a new person
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	//route for deleting a person
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	//serve our API over PORT 8080 and log any errors
	log.Fatal(http.ListenAndServe(":8080", router))
}
