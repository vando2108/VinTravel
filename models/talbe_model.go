package models

type Image_table struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int               `form:"parent_id" json:"parent_id" binding:"required"`
  Url 	       string            `form:"url" json:"url" binding:"required"`
}

type Related_table struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int               `form:"parent_id" json:"parent_id" binding:"required"`
  Name         string 		 `form:"related" json:"related" binding:"required"`
}
