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

func ReadAllProduct(c *gin.Context) {
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }
  productRepo := repoimpl.NewProductRepo(db.SQL)
  ret, err := productRepo.ReadAllProduct()
  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, ret)
}

func ReadListProductByCategories(c *gin.Context) {
  var requestData struct {
    Categories string `json:"categories" form:"categories" binding:"required"`
  }
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

  ret, err := productRepo.ReadListProductByCategories(requestData.Categories)

  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot read data from database")
    return
  }

  var res []struct {
    Name string
    Images []string
  }
  
  for i := range ret {
    fmt.Println(ret[i])
    res = append(res, 
      struct {
        Name string
        Images []string
      } {
        ret[i].Name,
        ret[i].Images,
      },
    )
  }
  c.JSON(http.StatusOK, res)
}

func ReadListProductByListName(c *gin.Context) {
  var requestData struct {
    Name []string `json:"name" form:"name" binding:"required"`
  }

  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from request")
    return
  }

  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }
  productRepo := repoimpl.NewProductRepo(db.SQL)
  
  var res [] struct {
    Name string
    Images []string
  }

  for i := range requestData.Name {
    queryData, err := productRepo.ReadProduct(requestData.Name[i])
    if err != nil {
      c.JSON(http.StatusInternalServerError, "Cannot read data from database")
      return
    }
    res = append(res,
      struct {
	Name string
	Images []string
      } {
	queryData.Name,
	queryData.Images,
      },
    )
  }
  c.JSON(http.StatusOK, res)
}
