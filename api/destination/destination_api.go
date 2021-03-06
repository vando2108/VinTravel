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

  var requestData models.Destination
  if err := c.ShouldBindJSON(&requestData); err != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from request")
    return
  }
  fmt.Println(requestData)
  
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return 
  } 
  defer db.SQL.Close()

  destinationRepo := repo.NewDestinationRepo(db.SQL)

  err = destinationRepo.CreateDestination(requestData)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  fmt.Println("Create new destination: ", requestData)
  c.JSON(200, "Create succesful")
}

func ReadDestination(c *gin.Context) {
  if jwt.TokenValid(c) != nil {
    return 
  }
  var requestData struct {
    Name string `json:"name" form:"name" binding:"required"`
  }
  if err := c.ShouldBindJSON(&requestData); err != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data")
    return 
  }
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return 
  } 
  defer db.SQL.Close()

  destinationRepo := repo.NewDestinationRepo(db.SQL)
  response, err := destinationRepo.ReadDestination(requestData.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return 
  }
  c.JSON(http.StatusOK, response)
}
