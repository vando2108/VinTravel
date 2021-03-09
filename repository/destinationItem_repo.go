package repository

import "vintravel/models"

type DestinationItemRepo interface {
  CreateDestinationItem(int, models.Item) (error)
}
