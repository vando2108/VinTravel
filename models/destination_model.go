package models

type Destination struct {
  Id 	       		int 	       	 `form:"id" json:"id"`
  Name         		string    	 `form:"name" json:"name" binding:"required"`
  Address      		string    	 `form:"address" json:"address" binding:"required"`
  CityProvince 		string 	 	 `form:"cityProvince" json:"cityProvince" binding:"required"`
  Description  		string     	 `form:"description" json:"description" binding:"required"`
  Coordinate   		string     	 `form:"coordinate" json:"coordinate" binding:"required"`
  Items        		[]Item           `form:"items" json:"items" binding:"required"`
  Images       		[]string  	 `form:"images" json:"images" binding:"required"`
  Nearby       		[]string  	 `form:"nearby" json:"nearby" binding:"required"`
  RelatedNearby         []string         `form:"relatedNearby" json:"relatedNearby" binding:"required"`
  Related      		[]string         `form:"related" json:"related" binding:"required"`
  Tags         		[]string 	 `form:"tags" json:"tags" binding:"required"`
  Types 		[]string 	 `form:"types" json:"types" binding:"required"`
  Functionalities 	[]string         `form:"functionalities" json:"functionalities" binding:"required"`   
}

type Destination_detail struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Name         string    	 `form:"name" json:"name" binding:"required"`
  Address      string    	 `form:"address" json:"address" binding:"required"`
  CityProvince string 	 	 `form:"cityprovince" json:"cityprovince" binding:"required"`
  Description  string     	 `form:"description" json:"description" binding:"required"`
  Coordinate   string     	 `form:"coordinate" json:"coordinate" binding:"required"`
}

type Destination_item struct {
  Id 	       int 	       	 `form:"id" json:"id"`
  Parent_id    int 		 `form:"parent_id" json:"parent_id" binding:"required"`
  Name         string    	 `form:"name" json:"name" binding:"required"`
  Price        float64  	 `form:"price" json:"price" binding:"required"`
}

type DestinationAPI struct {
  Id 	       		int 	 	`form:"id" json:"id" binding:"required"`
  Name         		string    	`form:"name" json:"name" binding:"required"`
  Address      		string    	`form:"address" json:"address" binding:"required"`
  CityProvince 		string 	 	`form:"cityProvince" json:"cityProvince" binding:"required"`
  Description  		string    	`form:"description" json:"description" binding:"required"`
  Coordinate   		string     	`form:"point" json:"point" binding:"required"`
  MaxPrice    		float64   	`form:"maxPrice" json:"maxPrice" binding:"required"`
  MinPrice    		float64   	`form:"minPrice" json:"minPrice" binding:"required"`
  AvgRatings   		float64   	`form:"avgRatings" json:"avgRatings" binding:"required"`
  Rating                []Rating        `form:"rating" json:"rating" binding:"required"`
  Images       		[]string  	`form:"images" json:"images" binding:"required"`
  Nearby       		[]string  	`form:"nearby" json:"nearby" binding:"required"`
  RelatedNearby         []string        `form:"relatedNearby" json:"relatedNearby" binding:"required"`
  Related      		[]string     	`form:"related" json:"related" binding:"required"`
  Items        		[]Item 	        `form:"items" json:"items" binding:"required"`
  Tags         		[]string 	`form:"tags" json:"tags" binding:"required"`
  Types 		[]string 	`form:"types" json:"types" binding:"required"`
  Functionalities 	[]string        `form:"functionalities" json:"functionalities" binding:"required"`   
}
