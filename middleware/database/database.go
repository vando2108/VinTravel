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

func CreateFunctionality(db *gorm.DB, parent_id int, tableName string, functionalities []string) (error) {
  for i := range functionalities {
    temp := models.Functionality_table {
      Parent_id: parent_id,
      Functionality: functionalities[i],
    }
    err := db.Table(tableName).Create(&temp).Error
    if err != nil {
      return err
    }
  }
  return nil
}

func CreateType(db *gorm.DB, parent_id int, tableName string, types []string) (error) { 
  for i := range types {
    temp := models.Type_table {
      Parent_id: parent_id,
      Name: types[i],
    }
    err := db.Table(tableName).Create(&temp).Error
    if err != nil {
      return err
    }
  }
  return nil
}

func CreateTag(db *gorm.DB, parent_id int, tableName string, tags []string) (error) {  
  for i := range tags {
    temp := models.Tag_table {
      Parent_id: parent_id,
      Tag: tags[i],
    }
    err := db.Table(tableName).Create(&temp).Error
    if err != nil {
      return err
    }
  }
  return nil
}
