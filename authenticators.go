package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var key = []byte{}

func main(){
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	//hmac implementation
	for i := 1; i <= 64; i++{
		key = append(key, byte(i))
	}

	h := sha512.New()
	h.Write([]byte("hello"))
	s := h.Sum(nil)
	fmt.Println(hex.EncodeToString(s))

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

/* bcrypt functions*/
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

/* Bearer tokens and HMAC
   	Bearer tokens: Uses authorization head & keyword "bearer"
   	Used for: prevent faked bearer tokens
   	Used in:  Cryptographic signing --> Way to prove the value was created/validated by a
             specific person
	HMAC -- Hash Message Authentication Code  --> https://godoc.org/crypto/hmac
	In general HMAC is a cryptographic signing algorithm making sure messages have not changed
*/
func signMessage(msg []byte)([]byte, error){
	h := hmac.New(sha512.New, key)
	 _, err := h.Write(msg)
	 if err != nil{
	 	return nil, fmt.Errorf("Error while hashing message %w", err)
	 }
	 signature := h.Sum(nil)   // get back your hash
	 return signature, nil
}

/* Check signature of signed message*/
func checkSig(msg, sig []byte)(bool, error)  {
	//first sign the message
	newSig, err := signMessage(msg)
	if err != nil{
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}

	/* Check if true or false*/
	same := hmac.Equal(newSig, sig)
	return same, nil

}