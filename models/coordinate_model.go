package models

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
  Lat float64 `json:"lat" form:"lat" binding:"required"`
  Lng float64 `json:"lng" form:"lng" binding:"required"`
}

func StrToCoordinate(s string) (float64, float64, error) {
  temp := strings.Split(s, ",")  
  var lat, lng float64
  var err error
  if lat, err = strconv.ParseFloat(temp[0], 64); err != nil {
    return 0, 0, err
  }
  if lng, err = strconv.ParseFloat(temp[1], 64); err != nil {
    return 0, 0, err
  }
  return lat, lng, nil
}

func rad(x float64) (float64) {
  return x * math.Pi / 180
}

func GetDistance(p1 Coordinate, p2 Coordinate) (float64) {
  r := float64(6378137)
  dLat := rad(p2.Lat - p1.Lat)
  dLong := rad(p2.Lng - p1.Lng)
  a := math.Sin(dLat / 2) * math.Sin(dLat / 2) + 
       math.Cos(rad(p1.Lat)) * math.Cos(rad(p2.Lat)) *
       math.Sin(dLong / 2) * math.Sin(dLong / 2)
  c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))
  d := r * c
  return d / 1000
}
