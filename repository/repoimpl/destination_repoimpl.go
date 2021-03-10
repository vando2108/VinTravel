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

// type Destination_detail struct {
//   Id 	       int 	       	 `form:"id" json:"id"`
//   Name         string    	 `form:"name" json:"name" binding:"required"`
//   Address      string    	 `form:"address" json:"address" binding:"required"`
//   CityProvince string 	 	 `form:"cityprovince" json:"cityprovince" binding:"required"`
//   Description  string     	 `form:"description" json:"description" binding:"required"`
//   Coordinate   string     	 `form:"coordinate" json:"coordinate" binding:"required"`
// }
func (d *DestinationRepoImpl) ReadDestination(name string) (error) {
  var result models.DestinationAPI
  var destination_detail models.Destination_detail
  err := d.Db.Table("destination_detail").Find(&destination_detail, "name = ?", name).Error
  if err != nil {
    return err
  }
  result.Id = destination_detail.Id
  result.Name = destination_detail.Name
  result.Address = destination_detail.Address
  result.CityProvince = destination_detail.CityProvince
  result.Description = destination_detail.Description
  result.Coordinate = destination_detail.Coordinate
  err = middleware.ReadFunctionality(d.Db, destination_detail.Id, "destination_functionality", &result)  
  if err != nil {
    return err
  }
  err = middleware.ReadItem(d.Db, destination_detail.Id, "destination_item", &result)    
  return nil
}
