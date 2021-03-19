package models

type Tag_table struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int               `form:"parent_id" json:"parent_id" binding:"required"`
  Tag 	       string 		 `form:"tag" json:"tag" binding:"required"`
}


type Type_table struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int               `form:"parent_id" json:"parent_id" binding:"required"`
  Name         string 		 `form:"related" json:"related" binding:"required"`
}

type Functionality_table struct {
  Id 	       	int 	       	 `form:"id" json:"id"`
  Parent_id    	int              `form:"parent_id" json:"parent_id" binding:"required"`
  Functionality string 		 `form:"functionality" json:"functionality" binding:"required"`
}

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

type Nearby_table struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int               `form:"parent_id" json:"parent_id" binding:"required"`
  Name         string 		 `form:"name" json:"name" binding:"required"`
}

type RelatedNearby_table struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int               `form:"parent_id" json:"parent_id" binding:"required"`
  Name         string 		 `form:"related" json:"related" binding:"required"`
}
