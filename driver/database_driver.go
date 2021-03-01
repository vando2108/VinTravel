package driver

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "fmt"
)

type PostgresDB struct {
  SQL *gorm.DB
}

var Postgres = &PostgresDB{}

func Connect(host, port, user, password, dbname string) (*PostgresDB, error) {
  connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
				host, port, user, password, dbname)
  db, err := gorm.Open("postgres", connectionStr)
  if err != nil {
    return nil, err
  }
  
  fmt.Println("Connect sucessful")

  Postgres.SQL = db
  return Postgres, nil
}
