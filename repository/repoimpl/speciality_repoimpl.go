package repoimpl

import (
	// "vintravel/models"
	"vintravel/models"
	repo "vintravel/repository"
	"fmt"

	"github.com/jinzhu/gorm"
)

type SpecialityRepoImpl struct {
  Db *gorm.DB
}

func NewSpecialityRepo(db *gorm.DB) repo.SpecialityRepo {
  return &SpecialityRepoImpl {
    Db: db,
  }
}

func (p *SpecialityRepoImpl) GetImages(id int) ([]*models.Speciality_image, error) {
  var ret []*models.Speciality_image
  err := p.Db.Table("speciality_image").Find(&ret).Error
  return ret, err
}

func (p *SpecialityRepoImpl) ReadSpeciality(name string) (models.SpecialityApi, error) {
  var speciality models.Speciality_detail
  speciality.Name = name
  err := p.Db.Table("speciality_detail").First(&speciality, "name = ?", name).Error
  if err != nil{
    return models.SpecialityApi{}, err
  }

  var image_query []*models.Speciality_image
  err = p.Db.Table("speciality_image").Find(&image_query, "speciality_parent_id= ?", speciality.Id).Error
  if err != nil {
    return models.SpecialityApi{}, err
  }

  var related_query []*models.Speciality_related
  err = p.Db.Table("speciality_related").Find(&related_query, "speciality_parent_id = ?", speciality.Id).Error
  if err != nil {
    return models.SpecialityApi{}, err
  }

  var categories_query []*models.Speciality_categories
  err = p.Db.Table("speciality_categories").Find(&categories_query, "speciality_parent_id = ?", speciality.Id).Error
  if err != nil {
    return models.SpecialityApi{}, err
  }
  

  ret := models.SpecialityApi{
    Name: speciality.Name,
    Origin: speciality.Origin,
    Voice: speciality.Voice,
    Description: speciality.Description,
    // Categories: speciality.Categories,
  }

  for i := range image_query {
    ret.Images = append(ret.Images, image_query[i].Url)
  }

  for i := range related_query {
    ret.Related = append(ret.Related, related_query[i].Name)
  }

  for i := range categories_query {
    ret.Categories = append(ret.Categories, categories_query[i].Cate)
  }
  
  return ret, nil
}

func (p *SpecialityRepoImpl) CreateSpeciality(speciality *models.SpecialityApi) (error) {
  newSpeciality := models.Speciality_detail{
    Name: speciality.Name,
    Origin: speciality.Origin,
    Voice: speciality.Voice,
    Description: speciality.Description,
    // Categories: speciality.Categories,
  }

  err := p.Db.Table("speciality_detail").Create(&newSpeciality).Error
  fmt.Println(newSpeciality)
  if err != nil {
    return err
  }

  for i :=  range speciality.Categories {
    temp := models.Speciality_categories{Speciality_parent_id: newSpeciality.Id, Cate: speciality.Categories[i]}
    err := p.Db.Table("speciality_categories").Create(&temp).Error
    if err != nil {
      return err
    }
  }

  for i := range speciality.Images {
    temp := models.Speciality_image{Speciality_parent_id: newSpeciality.Id, Url: speciality.Images[i]}
    err := p.Db.Table("speciality_image").Create(&temp).Error
    if err != nil {
      return err
    }
  }
  
  for i := range speciality.Related {
    fmt.Println(speciality.Related[i])
    temp := models.Speciality_related{Speciality_parent_id: newSpeciality.Id, Name: speciality.Related[i]}
    err := p.Db.Table("speciality_related").Create(&temp).Error
    if err != nil {
      return err
    }    
  }

  return nil
}

func (p *SpecialityRepoImpl) DeleteSpeciality(name string) (error) {
  var queryData models.Speciality_detail
  queryData.Name = name
  err := p.Db.Table("speciality_detail").First(&queryData, "name = ?", name).Error
  if err != nil{
    return err
  }

  err = p.Db.Table("speciality_image").Where("speciality_parent_id = ?", queryData.Id).Delete(&models.Speciality_image{}).Error
  if err != nil {
    return err
  }

  err = p.Db.Table("speciality_related").Where("speciality_parent_id = ?", queryData.Id).Delete(&models.Speciality_related{}).Error
  if err != nil {
   return err
  }

  err = p.Db.Table("speciality_detail").Delete(&queryData).Error
  return err
}

func (p *SpecialityRepoImpl) ReadAllSpeciality() ([]models.SpecialityApi, error) {
  var specialityDetails []*models.Speciality_detail
  err := p.Db.Table("speciality_detail").Find(&specialityDetails).Error
  if err != nil {
    return nil, err
  }
  
  var ret []models.SpecialityApi
  for i := range specialityDetails {
    var images_query []*models.Speciality_image
    var related_query []*models.Speciality_related
    var categories_query []*models.Speciality_categories
    var temp models.SpecialityApi
    err = p.Db.Table("speciality_image").Find(&images_query, "speciality_parent_id = ?", specialityDetails[i].Id).Error
    if err != nil {
      return nil, err
    }
    fmt.Println(images_query)

    err = p.Db.Table("speciality_related").Find(&related_query, "speciality_parent_id = ?", specialityDetails[i].Id).Error
    if err != nil {
      return nil, err
    }

    err = p.Db.Table("speciality_categories").Find(&categories_query, "speciality_parent_id = ?", specialityDetails[i].Id).Error
    if err != nil {
      return nil, err
    }

    temp = models.SpecialityApi{
      Name: specialityDetails[i].Name,
      Origin: specialityDetails[i].Origin,
      Voice: specialityDetails[i].Voice,
      Description: specialityDetails[i].Description,
    }

    for j := range images_query {
      temp.Images = append(temp.Images, images_query[j].Url)
    }
    for j := range related_query {
      temp.Related = append(temp.Related, related_query[j].Name)
    }
    for j := range categories_query {
      temp.Categories = append(temp.Categories, categories_query[j].Cate)
    }
    ret = append(ret, temp)
  }

  return ret, nil
}

func (p *SpecialityRepoImpl) ReadListSpecialityByCategories(categories string) ([]models.SpecialityApi, error) {
  var categories_query []*models.Speciality_categories
  err := p.Db.Table("speciality_categories").Find(&categories_query, "cate = ?", categories).Error
  if err != nil {
    return nil, err
  }

  var ret []models.SpecialityApi 
  for i := range categories_query {
    var specialityQuery models.Speciality_detail
    err := p.Db.Table("speciality_detail").Find(&specialityQuery, "id = ?", categories_query[i].Speciality_parent_id).Error
    if err != nil {
      return nil, err
    }
    var temp models.SpecialityApi
    temp = models.SpecialityApi{
      Name: specialityQuery.Name,
      Origin: specialityQuery.Origin,
      Voice: specialityQuery.Voice,
      Description: specialityQuery.Description,
    }
    var images_query []*models.Speciality_image
    err = p.Db.Table("speciality_image").Find(&images_query, "speciality_parent_id = ?", categories_query[i].Speciality_parent_id).Error
    if err != nil {
      return nil, err
    }
    for j := range images_query {
      temp.Images = append(temp.Images, images_query[j].Url)
    }
    var related_query []*models.Speciality_related
    err = p.Db.Table("speciality_related").Find(&related_query, "speciality_parent_id = ?", categories_query[i].Speciality_parent_id).Error
    if err != nil {
      return nil, err
    }
    for j := range related_query {
      temp.Related = append(temp.Related, related_query[j].Name)
    }
    var categories_query2 []*models.Speciality_categories
    err = p.Db.Table("speciality_categories").Find(&categories_query2, "speciality_parent_id = ?", categories_query[i].Speciality_parent_id).Error
    if err != nil {
      return nil, err
    }
    for j := range categories_query2 {
      temp.Categories = append(temp.Categories, categories_query2[j].Cate)
    }
    ret = append(ret, temp)
  }
  return ret, nil
}
