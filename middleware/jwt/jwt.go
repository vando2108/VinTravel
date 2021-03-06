package jwt

import (
	// "fmt"
	"errors"
	"net/http"
	"time"
	"vintravel/configs"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Create(username string) (string, error) {
  claims := jwt.MapClaims{}
  claims["authorized"] = true
  claims["username"] = username
  claims["exp"] = time.Now().Add(12 * time.Hour).Unix()
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(configs.Serect_key))
} 

func Verify(userToken string) (error) {
  token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
    return []byte(configs.Serect_key), nil
  })

  if err != nil {
    return err
  }

  if token.Valid {
    return nil
  }

  return errors.New("Token is not valid")
} 

func TokenValid(c *gin.Context) error {
  if err := Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return err
  }
  return nil
}
