package models

type Destination_detail struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Name         string    	 `form:"name" json:"name" binding:"required"`
  Address      string    	 `form:"address" json:"address" binding:"required"`
  CityProvince string 	 	 `form:"cityProvince" json:"cityProvince" binding:"required"`
  Description  string     	 `form:"description" json:"description" binding:"required"`
  Items        []Item_detail     `form:"items" json:"items" binding:"required"`
  Images       []string  	 `form:"images" json:"images" binding:"required"`
  Related      []int             `form:"related" json:"related" binding:"required"`
}

type DestinationAPI struct {
  Id 	       int 	 `form:"id" json:"id" binding:"required"`
  Name         string    `form:"name" json:"name" binding:"required"`
  Address      string    `form:"address" json:"address" binding:"required"`
  CityProvince string 	 `form:"cityProvince" json:"cityProvince" binding:"required"`
  Description  string    `form:"description" json:"description" binding:"required"`
  Images       []string  `form:"images" json:"images" binding:"required"`
  AvgPrices    float64   `form:"avgPrices" json:"avgPrices" binding:"required"`
  AvgRatings   float64   `form:"avgRatings" json:"avgRatings" binding:"required"`
  CntComment   int       `form:"cntComment" json:"cntComment" binding:"required"`
  CntRating    int       `form:"cntRating" json:"cntRating" binding:"required"`
  Nearby       []string  `form:"nearby" json:"nearby" binding:"required"`
  Related      []int     `form:"related" json:"related" binding:"required"`
  Items        []Item_detail     `form:"items" json:"items" binding:"required"`
}
