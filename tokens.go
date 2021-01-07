package main

import (
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	_ "hash/fnv"
	"io"
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
	SessionID int64
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
	signedToken, err := t.SignedString(keys[currentKid].key)
	if err != nil{
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}

func generateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil{
		return fmt.Errorf("Error generatingNewKey while generting key %w", err)
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("Error in generateNewKey while generating kid: %w", err)
	}
	keys[uid.String()] = key{
		key:     newKey,
		created: time.Now(),
	}
	currentKid = uid.String()
	return nil
}

type key struct {
	key []byte
	created time.Time
}
var currentKid = ""
var keys = map[string]key{}

// parse token and return claim
func parseToken(signedToken string)(*UserClaims, error)  {

	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func (t *jwt.Token)(interface{}, error){
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg(){
			return nil, fmt.Errorf("Invalid signing algorithm")
		}
		kid, ok := t.Header["kid"].(string)
		if !ok{
			return nil, fmt.Errorf("Invalid key ID")
		}
		k, ok := keys[kid]
		if !ok{
			return nil, fmt.Errorf("Invalid key ID")
		}
		return k.key, nil
	})
	if err != nil{
		return nil, fmt.Errorf("Error in parseToken while verifying: %w", err)
	}
	if !t.Valid{
		return nil, fmt.Errorf("Error in parseToken, token is not valid ")
	}
	return t.Claims.(*UserClaims), nil //Assert
}