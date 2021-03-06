package models

type Item_detail struct {
  Name string `form:"name" json:"name" binding:"required"`
  Price float64 `form:"price" json:"price" binding:"required"`
}
