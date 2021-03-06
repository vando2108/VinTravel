package repository

import "vintravel/models"

type DescriptionRepo interface {
  CreateDestination(models.Destination_detail) (error)
}
