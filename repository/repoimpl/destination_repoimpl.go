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
  // CityProvince string 	 	 `form:"cityprovince" json:"cityProvince" binding:"required"`
  // Description  string     	 `form:"description" json:"description" binding:"required"`
  // Coordinate   Point     	 `form:"coordinate" json:"coordinate" binding:"required"`
  // Items        []Item_detail     `form:"items" json:"items" binding:"required"`
  // Images       []string  	 `form:"images" json:"images" binding:"required"`
  // Related      []string          `form:"related" json:"related" binding:"required"`
  // Tags         		[]string 	`form:"tags" json:"tags" binding:"required"`
  // Types 		[]string 	`form:"types" json:"types" binding:"required"`
  // Functionalities 	[]string        `form:"functionalities" json:"functionalities" binding:"required"`
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
  err = middleware.CreateFunctionality(d.Db, destination_detail.Id, "destination_functionality", destination.Functionalities)
  if err != nil {
    return err
  }
  err = middleware.CreateTag(d.Db, destination_detail.Id, "destination_tag", destination.Tags)
  if err != nil {
    return err
  }
  err = middleware.CreateType(d.Db, destination_detail.Id, "destination_type", destination.Types)
  if err != nil {
    return err
  }

  destinationItemRepo := NewDestinationItemRepo(d.Db)
  for i := range destination.Items {
    err := destinationItemRepo.CreateDestinationItem(destination_detail.Id, destination.Items[i])
    if err != nil {
      return err
    }
  }
  
  return nil
}
