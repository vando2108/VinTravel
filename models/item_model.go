package models

type Item struct {
  Id int `form:"id" json:"id" binding:"required"`
  Name string `form:"name" json:"name" binding:"required"`
  Price float64 `form:"price" json:"price" binding:"required"`
  Images []string `form:"images" json:"images" binding:"required"`
}

type Item_detail struct {
  Id int `form:"id" json:"id" binding:"required"`
  Name string `form:"name" json:"name" binding:"required"`
  Price float64 `form:"price" json:"price" binding:"required"`
}
