
package repoimpl

import (
	// "vintravel/models"
	"vintravel/models"
	repo "vintravel/repository"
	// "fmt"

	"github.com/jinzhu/gorm"
)

type DestinationRepoImpl struct {
  Db *gorm.DB
}

func NewDestinationRepo(db *gorm.DB) repo.DescriptionRepo {
  return &DestinationRepoImpl {
    Db: db,
  }
}

func (d *DestinationRepoImpl) CreateDestination(destination models.Destination_detail) (error) {
  err := d.Db.Table("destination_detail").Create(&destination).Error
  if err != nil {
    return err
  }
  
  // for i := range destination.Images {
  //
  // }
  
  return nil
}
