package product

import (
	"fmt"
	"net/http"
	"vintravel/configs"
	"vintravel/driver"
	"vintravel/models"
	repoimpl "vintravel/repository/repoimpl"

	"github.com/gin-gonic/gin"
)

func CreateNewProduct(c *gin.Context) {
  var requestData models.ProductApi 
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data")
    return
  }
  
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }

  productRepo := repoimpl.NewProductRepo(db.SQL)
  err = productRepo.CreateProduct(&requestData)  
  if err != nil {
    c.JSON(401, "Cannot create new product")
    return
  }
  c.JSON(http.StatusOK, "Create new product succesful")
  fmt.Println("Create new product: ", requestData)
}

type deleteApi struct {
  Name string `form:"name" json:"name" binding:"required"`
}

func DeleteProduct(c *gin.Context) {
  var requestData deleteApi
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data")
    return
  }

  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }
  productRepo := repoimpl.NewProductRepo(db.SQL)

  err = productRepo.DeleteProduct(requestData.Name)

  if err != nil {
    c.JSON(401, err.Error())
    return
  }
  
  c.JSON(http.StatusOK, "Delete Succesful")
  fmt.Println("Delete product: ", requestData.Name)
}

func ReadProduct(c *gin.Context) {
  var requestData deleteApi
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data")
    return
  }

  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }
  productRepo := repoimpl.NewProductRepo(db.SQL)

  ret, err := productRepo.ReadProduct(requestData.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  
  c.JSON(http.StatusOK, ret)
}
