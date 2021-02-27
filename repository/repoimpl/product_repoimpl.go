package repoimpl

import (
	// "vintravel/models"
	"vintravel/models"
	repo "vintravel/repository"
	"fmt"

	"github.com/jinzhu/gorm"
)

type ProductRepoImpl struct {
  Db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repo.ProductRepo {
  return &ProductRepoImpl {
    Db: db,
  }
}

func (p *ProductRepoImpl) GetImages(id int) ([]*models.Product_image, error) {
  var ret []*models.Product_image
  err := p.Db.Table("product_images").Find(&ret).Error
  return ret, err
}

func (p *ProductRepoImpl) CreateProduct(product *models.ProductApi) (error) {
  newProduct := models.Product_details{
    Name: product.Name,
    Origin: product.Origin,
    Voice: product.Voice,
    Description: product.Description,
    Categories: product.Categories,
    Related: product.Related,
  }

  err := p.Db.Table("product_details").Create(&newProduct).Error
  fmt.Println(newProduct)
  if err != nil {
    return err
  }

  for i := range product.Images {
    fmt.Println(product.Images[i])
    temp := models.Product_image{Id: newProduct.Id, Url: product.Images[i]}
    err := p.Db.Table("product_images").Create(&temp).Error
    if err != nil {
      return err
    }
  }

  return nil
}

func (p *ProductRepoImpl) DeleteProduct(name string) (error) {
  var queryData models.Product_details
  queryData.Name = name
  err := p.Db.Table("product_details").First(&queryData, "name = ?", name).Error
  if err != nil{
    return err
  }

  err = p.Db.Table("product_images").Where("id = ?", queryData.Id).Delete(&models.Product_image{}).Error
  if err != nil {
    return err
  }
  err = p.Db.Table("product_details").Delete(&queryData).Error
  return err
}

func (p *ProductRepoImpl) ReadProduct(name string) (models.ProductApi, error) {
  var product models.Product_details
  product.Name = name
  err := p.Db.Table("product_details").First(&product, "name = ?", name).Error
  if err != nil{
    return models.ProductApi{}, err
  }

  var images_query []*models.Product_image
  err = p.Db.Table("product_images").Find(&images_query, "id = ?", product.Id).Error
  if err != nil {
    return models.ProductApi{}, err
  }

  ret := models.ProductApi{
    Name: product.Name,
    Origin: product.Origin,
    Voice: product.Voice,
    Description: product.Description,
    Categories: product.Categories,
    Related: product.Related,
  }

  for i := range images_query {
    ret.Images = append(ret.Images, images_query[i].Url)
  }
  
  return ret, nil
}

func (p *ProductRepoImpl) ReadAllProduct() ([]models.ProductApi, error) {
  var productDetails []*models.Product_details
  err := p.Db.Table("product_details").Find(&productDetails).Error
  if err != nil {
    return nil, err
  }
  
  var ret []models.ProductApi
  for i := range productDetails {
    var images_query []*models.Product_image
    var temp models.ProductApi
    err = p.Db.Table("product_images").Find(&images_query, "id = ?", productDetails[i].Id).Error
    if err != nil {
      return nil, err
    }
    fmt.Println(images_query)

    temp = models.ProductApi{
      Name: productDetails[i].Name,
      Origin: productDetails[i].Origin,
      Voice: productDetails[i].Voice,
      Description: productDetails[i].Description,
      Categories: productDetails[i].Categories,
      Related: productDetails[i].Related,
    }

    for j := range images_query {
      temp.Images = append(temp.Images, images_query[j].Url)
    }
    ret = append(ret, temp)
  }

  return ret, nil
}
