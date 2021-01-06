package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main(){
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	passwrd := "God!sGoodT0m3"
	hashedPasswrd, err := hashPassword(passwrd)
	if err != nil{
		panic(err)
	}
	err = comparePassword(passwrd, hashedPasswrd)
	if err != nil{
		log.Fatalln("Not logged in")
	}
	log.Println("Logged in!")
}

func hashPassword(password string)([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass,[]byte(password)) // returns an error on failure
	if err != nil{
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}