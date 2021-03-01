package main

import (
	// "github.com/gin-gonic/gin"
	// "fmt"
	auth "vintravel/api/auth"
	product_api "vintravel/api/product"
	config "vintravel/configs"
	"vintravel/driver"

	"github.com/gin-gonic/gin"
)

func main() {
  db, err := driver.Connect(config.Host, config.Port, config.User, config.Password, config.Name)

  if err != nil {
    panic(err)
  }

  router := gin.Default()

  router.POST("/auth/regsiter", auth.RegsiterUser)
  router.POST("/auth/login", auth.Login)
  router.GET("/auth/readuserdata", auth.ReadUserData)
  router.POST("/auth/updateuser", auth.UpdateUser)

  router.POST("/product/create", product_api.CreateNewProduct)
  router.POST("/product/delete", product_api.DeleteProduct)
  router.GET("/product/readproduct", product_api.ReadProduct)
  router.GET("/product/readallproduct", product_api.ReadAllProduct)
  router.GET("product/ReadListProductByCategories", product_api.ReadListProductByCategories)

  router.Run(":3000")

  defer db.SQL.Close()
}
