package database

import (
	"vintravel/models"

	"github.com/jinzhu/gorm"
)

func CreateImage(db *gorm.DB, parent_id int, tableName string, images []string) (error) {
  for i := range images {
    image := models.Image_table {
      Parent_id: parent_id,
      Url: images[i],
    }
    
    err := db.Table(tableName).Create(&image).Error
    if err != nil {
      return err
    }
  }
  return nil
}

func CreateRelated(db *gorm.DB, parent_id int, tableName string, related []string) (error) {
  for i := range related {
    temp := models.Related_table {
      Parent_id: parent_id,
      Name: related[i],
    }
    err := db.Table(tableName).Create(&temp).Error
    if err != nil {
      return err
    }
  }
  return nil
}
