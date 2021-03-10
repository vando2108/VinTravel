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

func ReadImage(db *gorm.DB, parent_id int, tableName string, ) (error) { 

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

func ReadFunctionality(db *gorm.DB, parent_id int, tableName string, destination *models.DestinationAPI) ([]string, error) {
  var listFunctionality []*models.Functionality_table
  err := db.Table(tableName).Find(&listFunctionality, "parent_id = ?", parent_id).Error
  if err != nil {
    return nil, err
  }
  var ret []string
  for i := range listFunctionality {
    ret = append(ret, listFunctionality[i].Functionality)
  }
  return ret, nil
}

func ReadItem(db *gorm.DB, parent_id int, tableName string, destination *models.DestinationAPI) (error) {
  var listItemDetail []*models.Item_detail
  err := db.Table(tableName).Find(&listItemDetail, "parent_id = ?", parent_id).Error
  if err != nil {
    return err
  }
  var result []models.Item
  for _, it := range listItemDetail {
    temp := models.Item {
      Id: it.Id,
      Name: it.Name,
      Price: it.Price,
    }
    var listImage []*models.Image_table
    err := db.Table("item_image").Find(&listImage, "parent_id = ?", it.Id).Error
    if err != nil {
      return err
    }
    result = append(result, temp)
 }
 destination.Items = result
 return nil
}
