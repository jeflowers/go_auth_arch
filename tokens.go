package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/* JSON Web Tokens aka JWT's
	Has two JSON objects and a signature
	Made up of three base64 sections, separated by periods.
		{JWT standard fields}.{custom fields}.{Signature}
	1) General JSON information:  Expiration, etc...
	2) Customized information: Session ID, etc ...
	3) signature
	https://godoc.org/github.com/dgrijalva/jwt-go
	jwt-go requires the following:
	Signing method which is an interface which provides three methods (HMAC, ECDSA, or RSA)
	Claims type (MapClaims, or StandardClaims)
*/

// create claim type
type UserClaims struct {
	// Using StandardClaims
	jwt.StandardClaims
	SessoionID int64
}

// recommended to override the Valid function
func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true){
		return fmt.Errorf("Token has expired")
	}

	if u.SessionID == 0{
		return fmt.Errorf("Invalid session ID")
	}
	return nil
}

func main() {

}

func createToken(c *UserClaims) (string, error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err :=t.SignedString(key)
	if err != nil{
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}