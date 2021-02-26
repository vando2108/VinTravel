package models

type Product_details struct {
  Id int `json:"id"`
  Origin string `json:"origin"`
  Name string `json:"name"`
  Voice string `json:"voice"`
  Description string `json:"description"`
  Categories string `json:"categories"`
  Related string `json:"related"`
}

type ProductApi struct {
  Name string `form:"name" json:"name" binding:"required"`
  Origin string `form:"origin" json:"origin" binding:"required"`
  Voice string `form:"voice" json:"voice" binding:"required"`
  Description string `form:"description" json:"description" binding:"required"`
  Categories string `form:"categories" json:"categories" binding:"required"`
  Related string `form:"related" json:"related" binding:"required"`
  Images []string `form:"images" json:"images" binding:"required"`
}

type Product_image struct {
  Id int
  Url string
}
