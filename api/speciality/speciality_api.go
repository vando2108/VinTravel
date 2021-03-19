package speciality

import (
	"fmt"
	"net/http"
	"vintravel/configs"
	"vintravel/driver"
	"vintravel/models"
	repoimpl "vintravel/repository/repoimpl"
	jwt "vintravel/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func CreateNewSpeciality(c *gin.Context) {
  if err := jwt.Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return
  }

  var requestData models.SpecialityApi 
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data")
    return
  }
  
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }
  defer db.SQL.Close()

  specialityRepo := repoimpl.NewSpecialityRepo(db.SQL)
  err = specialityRepo.CreateSpeciality(&requestData)  
  if err != nil {
    c.JSON(401, "Cannot create new speciality")
    return
  }
  c.JSON(http.StatusOK, "Create new speciality succesful")
  fmt.Println("Create new speciality: ", requestData)
}

type deleteApi struct {
  Name string `form:"name" json:"name" binding:"required"`
}

func DeleteSpeciality(c *gin.Context) {
  if err := jwt.Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return
  }
  var requestData struct {
    Name string `json:"name" form:"name" binding:"required"`
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
  defer db.SQL.Close()

  specialityRepo := repoimpl.NewSpecialityRepo(db.SQL)

  err = specialityRepo.DeleteSpeciality(requestData.Name)

  if err != nil {
    c.JSON(401, err.Error())
    return
  }
  
  c.JSON(http.StatusOK, "Delete Succesful")
  fmt.Println("Delete speciality: ", requestData.Name)
}

func ReadSpeciality(c *gin.Context) {
  if err := jwt.Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return
  }
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
  defer db.SQL.Close()

  specialityRepo := repoimpl.NewSpecialityRepo(db.SQL)

  ret, err := specialityRepo.ReadSpeciality(requestData.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  
  c.JSON(http.StatusOK, ret)
}

func ReadAllSpeciality(c *gin.Context) {
  if err := jwt.Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return
  }
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  }
  defer db.SQL.Close()

  specialityRepo := repoimpl.NewSpecialityRepo(db.SQL)
  ret, err := specialityRepo.ReadAllSpeciality()
  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, ret)
}

func ReadListSpecialityByCategories(c *gin.Context) {
  if err := jwt.Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return
  }
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
  defer db.SQL.Close()

  specialityRepo := repoimpl.NewSpecialityRepo(db.SQL)

  ret, err := specialityRepo.ReadListSpecialityByCategories(requestData.Categories)

  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  c.JSON(http.StatusOK, ret)
}

func ReadListSpecialityByListName(c *gin.Context) {
  if err := jwt.Verify(c.Request.Header["Authorization"][0]); err != nil {
    c.JSON(http.StatusNonAuthoritativeInfo, err.Error())
    return
  }
  var requestData struct {
    Name []string `json:"name" form:"name" binding:"required"`
  }

  if err := c.ShouldBindJSON(&requestData); err != nil {
    c.JSON(http.StatusBadRequest, err.Error)
    return
  }

  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error)
    return
  }
  defer db.SQL.Close()

  specialityRepo := repoimpl.NewSpecialityRepo(db.SQL)

  var ret []models.SpecialityApi
  for i := range requestData.Name {
    queryData, err := specialityRepo.ReadSpeciality(requestData.Name[i])
    if err != nil {
      c.JSON(http.StatusInternalServerError, err.Error)
      return
    }
    ret = append(ret, queryData)
  }
  c.JSON(http.StatusOK, ret)
}
