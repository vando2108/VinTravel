package models

type Rating struct {
  Id 		 int 	   `json:"id" form:"id"`
  Username 	 string    `json:"user_id" form:"user_id" binding:"required"`
  Destination_id int       `json:"destination_id" form:"destination_id" binding:"required"` 
  Star 		 int 	   `json:"star" form:"star" binding:"required"`
  Cmt 		 string    `json:"cmt" form:"cmt" binding:"required"`
}
