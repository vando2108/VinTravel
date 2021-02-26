package repoimpl

import (
	"vintravel/models"
	repo "vintravel/repository"
	"github.com/jinzhu/gorm"
	"fmt"
)

type UserRepoImpl struct {
  Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repo.UserRepo {
  return &UserRepoImpl{
    Db: db,
  }
}

func (u *UserRepoImpl) CreateNewUser(user *models.User) (error) {
  err := u.Db.Create(&user).Error
  if err != nil {
    fmt.Println(err)
    return err
  }
  fmt.Println("Create new user: ", user)
  return nil
}

func (u *UserRepoImpl) ReadUser(username string) (models.User, error) {
  var result models.User
  fmt.Println(username)
  err := u.Db.Table("users").Where("username = ?", username).Find(&result).Error
  return result, err
}

func (u *UserRepoImpl) UpdateUser(user *models.User) (error) {
  return u.Db.Table("users").Where("username=?", user.Username).Update(&user).Error
}
