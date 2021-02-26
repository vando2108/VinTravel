package repository

import (
  "vintravel/models"
)

type UserRepo interface {
  CreateNewUser(user *models.User) (error)
  ReadUser(username string) (models.User, error)
  UpdateUser(user *models.User) (error)
}
