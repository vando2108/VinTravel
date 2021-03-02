package jwt

import (
	// "fmt"
	"time"
	"vintravel/configs"

	jwt "github.com/dgrijalva/jwt-go"
)

func Create(username string) (string, error) {
  claims := jwt.MapClaims{}
  claims["authorized"] = true
  claims["username"] = username
  claims["iat"] = time.Now().Add(time.Hour * 12).Unix()
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(configs.Serect_key))
} 

func Verify(userToken string) (error) {
  _, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
    return []byte(configs.Serect_key), nil
  })

  return err
} 
