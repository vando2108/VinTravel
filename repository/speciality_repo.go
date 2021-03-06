package repository

import "vintravel/models"

type SpecialityRepo interface {
  GetImages(id int) ([]*models.Speciality_image, error) 
  CreateSpeciality(speciality *models.SpecialityApi) (error)
  ReadSpeciality(name string) (models.SpecialityApi, error)
  ReadAllSpeciality() ([]models.SpecialityApi, error)
  ReadListSpecialityByCategories(categories string) ([]models.SpecialityApi, error)
  DeleteSpeciality(name string) (error)
}
