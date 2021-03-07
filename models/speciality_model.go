package models

type Speciality_detail struct {
  Id int `json:"id"`
  Origin string `json:"origin"`
  Name string `json:"name"`
  Voice string `json:"voice"`
  Description string `json:"description"`
}

type SpecialityApi struct {
  Name string `form:"name" json:"name" binding:"required"`
  Origin string `form:"origin" json:"origin" binding:"required"`
  Voice string `form:"voice" json:"voice" binding:"required"`
  Description string `form:"description" json:"description" binding:"required"`
  Categories []string `form:"categories" json:"categories" binding:"required"`
  Related []string `form:"related" json:"related" binding:"required"`
  Images []string `form:"images" json:"images" binding:"required"`
}

type Speciality_image struct {
  Id int `form:"id" json:"id" binding:"required"`
  Speciality_parent_id int `form:"speciality_parent_id" json:"speciality_parent_id" binding:"required"`
  Url string `form:"url" json:"url" binding:"required"`
}

type Speciality_related struct {
  Id int `form:"id" json:"id" binding:"required"`
  Speciality_parent_id int `form:"speciality_parent_id" json:"speciality_parent_id" binding:"required"`
  Name string `form:"name" json:"name" binding:"required"`
}

type Speciality_categories struct {
  Id int `form:"id" json:"id" binding:"required"`
  Speciality_parent_id int `form:"speciality_parent_id" json:"speciality_parent_id" binding:"required"`
  Cate string `form:"cate" json:"cate" binding:"required"`
}
