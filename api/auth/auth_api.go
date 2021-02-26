package auth

import (
	// "fmt"
	"fmt"
	"net/http"
	"time"
	"vintravel/configs"
	"vintravel/driver"
	"vintravel/models"
	repo "vintravel/repository/repoimpl"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type regsiterForm struct {
  Name string `form:"name" json:"name" binding:"required"`
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
}

func RegsiterUser(c *gin.Context) {
  var requestData regsiterForm
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from request")
    return
  }

  if govalidator.IsNull(requestData.Name) || govalidator.IsNull(requestData.Username) || govalidator.IsNull(requestData.Password) {
    c.JSON(http.StatusBadRequest, "Data cannot empty")
    return
  }

  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  defer db.SQL.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  } 
  var id int 
  db.SQL.Model(&models.User{}).Count(&id)

  requestData.Username = models.Santize(requestData.Username)
  requestData.Password = models.Santize(requestData.Password)
  requestData.Name = models.Santize(requestData.Name)

  requestData.Password, err = models.Hash(requestData.Password)

  if err != nil {
    c.JSON(http.StatusBadRequest, err.Error)
    return
  }
  
  newUser := models.User {
    Id: id + 1,
    Name: requestData.Name,
    Username: requestData.Username,
    Password: requestData.Password,
    Created_at: time.Now(),
  }
  userRepo := repo.NewUserRepo(db.SQL) 

  err = userRepo.CreateNewUser(&newUser) 	
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return 
  }
  c.JSON(http.StatusOK, "Create succesful")
}

type loginForm struct {
  Username string `form:"username" json:"username" binding:"required"`
  Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
  var requestData loginForm
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from request")
    return
  }
  requestData.Username = models.Santize(requestData.Username)
  requestData.Password = models.Santize(requestData.Password)

  fmt.Println("Login reqeust: ",  requestData)

  //Connect to database
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  defer db.SQL.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  } 
  userRepo := repo.NewUserRepo(db.SQL)
  var queryUser models.User
  queryUser, err = userRepo.ReadUser(requestData.Username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  err = models.CheckPasswordHash(queryUser.Password, requestData.Password)
  if err != nil {
    c.JSON(401, err.Error())
    return
  }
  c.JSON(http.StatusAccepted, "Login succesful")
}

func ReadUserData(c *gin.Context) {
  var requestData struct {Username string `form:"username" json:"username" binding:"required"`}
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from request")
    return
  }
  requestData.Username = models.Santize(requestData.Username)  
  
  //Connect to database
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  defer db.SQL.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  } 
  userRepo := repo.NewUserRepo(db.SQL)
  var queryUser models.User
  queryUser, err = userRepo.ReadUser(requestData.Username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err.Error())
    return
  }
  c.JSON(http.StatusOK, queryUser)
}

func UpdateUser(c *gin.Context) {
  var requestData models.User
  if c.ShouldBindJSON(&requestData) != nil {
    c.JSON(http.StatusBadRequest, "Cannot parse data from request")
    return
  }

  var err error
  requestData.Username = models.Santize(requestData.Username)
  requestData.Password = models.Santize(requestData.Password)
  requestData.Name = models.Santize(requestData.Name)
  requestData.Password, err = models.Hash(requestData.Password)
  requestData.Created_at = time.Now()
  
  if err != nil {
    c.JSON(401, "Cannot hash new password")
    return
  }

  //Connect to database
  db, err := driver.Connect(configs.Host, configs.Port, configs.User, configs.Password, configs.Name)
  defer db.SQL.Close()
  if err != nil {
    c.JSON(http.StatusInternalServerError, "Cannot connect to database")
    return
  } 
  userRepo := repo.NewUserRepo(db.SQL)
  err = userRepo.UpdateUser(&requestData) 
  if err != nil {
    c.JSON(401, err.Error())
    return
  }
  c.JSON(http.StatusOK, "Update succesful")
}
