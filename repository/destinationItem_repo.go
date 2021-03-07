package repository

import "vintravel/models"

type DestinationItemRepo interface {
  CreateDestinationItem(models.Item) (error)
}
