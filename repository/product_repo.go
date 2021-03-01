package repository

import "vintravel/models"

type ProductRepo interface {
  GetImages(id int) ([]*models.Product_image, error) 
  CreateProduct(product *models.ProductApi) (error)
  ReadProduct(name string) (models.ProductApi, error)
  ReadAllProduct() ([]models.ProductApi, error)
  ReadListProductByCategories(categories string) ([]models.ProductApi, error)
  ReadListProductByListName(listName []string) ([]models.ProductApi, error)
  DeleteProduct(name string) (error)
}
