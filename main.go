package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"log"
	"net/http"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
func main(){


	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request){
		var user Person
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil{
			log.Println("decode error")
		}

		fmt.Fprintf(w, "%s %s is a good worker", user.Firstname, user.Lastname)
	})


	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request){
		statement := Person{
			Firstname: "John",
			Lastname: "Doe",
		}

		err := json.NewEncoder(w).Encode(statement)
		if err != nil{
			log.Println("Bad data to encode, printing out help statement", err)
		}
	})
	http.ListenAndServe(":8090", nil)
}



