package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}
func main(){
	p1 := person{
		First: "Evadne",
	}
	p2 := person{
		First: "Tennille",
	}

	xp := []person{p1,p2}
	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(bs)
	}
	fmt.Println("PRINT JS",string(bs))

	xp2 :=  []person{}
	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Back into a go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){

}

func bar(w http.ResponseWriter, r *http.Request){

}

