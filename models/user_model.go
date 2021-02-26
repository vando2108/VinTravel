package models

import (
	"html"
	"time"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
  Id int `form:"id" json:"id" binding:"required"`
  Name string `form:"name" json:"name" binding:"required"`
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
  Created_at time.Time `form:"created_at" json:"created_at" binding:"required"`
}

func Hash(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) (error) {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) (string) {
  data = html.EscapeString(strings.TrimSpace(data))
  return data
}
