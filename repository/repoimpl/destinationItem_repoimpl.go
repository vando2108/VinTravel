package repoimpl

import (
	// "vintravel/models"
	"vintravel/models"
	repo "vintravel/repository"
	// "fmt"
	middleware "vintravel/middleware/database"

	"github.com/jinzhu/gorm"
)

type DestinationItemRepoImpl struct {
  Db *gorm.DB
}

func NewDestinationItemRepo(db *gorm.DB) repo.DestinationItemRepo {
  return &DestinationItemRepoImpl {
    Db: db,
  }
}

func (repo *DestinationItemRepoImpl) CreateDestinationItem(parent_id int, item models.Item) (error) {
  temp := models.Destination_item{
    Parent_id: parent_id,
    Name: item.Name,
    Price: item.Price,
  } 
  err := repo.Db.Table("destination_item").Create(&temp).Error
  if err != nil {
    return err
  }
  if err = middleware.CreateImage(repo.Db, temp.Id, "item_image", item.Images); err != nil {
    return err
  } 
  
  return nil
}
