package destination

import (
	"fmt"
	"net/http"
	"vintravel/configs"
	"vintravel/driver"
	"vintravel/middleware/jwt"
	"vintravel/models"
	repo "vintravel/repository/repoimpl"

	"github.com/gin-gonic/gin"
)

// Name         string    	 `form:"name" json:"name" binding:"required"`
// Address      string    	 `form:"address" json:"address" binding:"required"`
// CityProvince string 	 	 `form:"cityProvince" json:"cityProvince" binding:"required"`
// Description  string     	 `form:"description" json:"description" binding:"required"`
// Items        []Item_detail     `form:"items" json:"items" binding:"required"`
// Images       []string  	 `form:"images" json:"images" binding:"required"`
// Related      []int             `form:"related" json:"related" binding:"required"`
func CreateDestination(c *gin.Context) {
  if jwt.TokenValid(c) != nil {
    return 
  }

  var requestData models.Destination_detail
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from reqeust")
    return
  }
  
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  defer db.SQL.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return 
  } 
  destinationRepo := repo.NewDestinationRepo(db.SQL)

  err = destinationRepo.CreateDestination(requestData)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  fmt.Println("Create new destination: ", requestData)
  c.JSON(200, "Create succesful")
}
