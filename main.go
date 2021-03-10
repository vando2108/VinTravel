package main

import (
	// "github.com/gin-gonic/gin"
	"net/http"
	auth "vintravel/api/auth"
	speciality_api "vintravel/api/speciality"
	destination_api "vintravel/api/destination"
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

  router.GET("/index", func (c *gin.Context) {
    c.String(http.StatusOK, "heelo")
  })
  router.POST("/auth/Register", auth.RegsiterUser)
  router.POST("/auth/Login", auth.Login)
  router.GET("/auth/ReadUserData", auth.ReadUserData)
  router.POST("/auth/UpdateUser", auth.UpdateUser)

  router.POST("/speciality/CreateNewSpeciality", speciality_api.CreateNewSpeciality)
  router.POST("/speciality/DeleteSpeciality", speciality_api.DeleteSpeciality)
  router.GET("/speciality/ReadSpeciality", speciality_api.ReadSpeciality)
  router.GET("/speciality/ReadAllSpeciality", speciality_api.ReadAllSpeciality)
  router.GET("speciality/ReadListSpecialityByCategories", speciality_api.ReadListSpecialityByCategories)
  router.GET("speciality/ReadListSpecialityByListName", speciality_api.ReadListSpecialityByListName)

  router.POST("/destination/CreateDestination", destination_api.CreateDestination)

  router.Run()


  defer db.SQL.Close()
}
