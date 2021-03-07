package repoimpl

import (
	// "vintravel/models"
	middleware "vintravel/middleware/database"
	"vintravel/models"
	repo "vintravel/repository"

	// "fmt"

	"github.com/jinzhu/gorm"
)

type DestinationRepoImpl struct {
  Db *gorm.DB
}

func NewDestinationRepo(db *gorm.DB) repo.DestinationRepo {
  return &DestinationRepoImpl {
    Db: db,
  }
}

  // Id 	       int 	       	 `form:"id" json:"id"`
  // Name         string    	 `form:"name" json:"name" binding:"required"`
  // Address      string    	 `form:"address" json:"address" binding:"required"`
  // CityProvince string 	 	 `form:"cityProvince" json:"cityProvince" binding:"required"`
  // Description  string     	 `form:"description" json:"description" binding:"required"`
  // Coordinate   Point     	 `form:"coordinate" json:"coordinate" binding:"required"`
  // Items        []Item_detail     `form:"items" json:"items" binding:"required"`
  // Images       []string  	 `form:"images" json:"images" binding:"required"`
  // Related      []string          `form:"related" json:"related" binding:"required"`
func (d *DestinationRepoImpl) CreateDestination(destination models.Destination) (error) {
  destination_detail := models.Destination_detail {
    Name: destination.Name,
    Address: destination.Address,
    CityProvince: destination.CityProvince,
    Description: destination.Description,
    Coordinate: destination.Coordinate,
  }
  err := d.Db.Table("destination_detail").Create(&destination_detail).Error
  if err != nil {
    return err
  }
  
  err = middleware.CreateImage(d.Db, destination_detail.Id, "destination_image", destination.Images)
  if err != nil {
    return err
  }
  err = middleware.CreateRelated(d.Db, destination_detail.Id, "destination_related", destination.Related)
  if err != nil {
    return err
  }

  destinationItemRepo := NewDestinationItemRepo(d.Db)
  for i := range destination.Items {
    err := destinationItemRepo.CreateDestinationItem(destination.Items[i])
    if err != nil {
      return err
    }
  }
  
  return nil
}
