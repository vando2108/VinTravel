package repository

import "vintravel/models"

type DestinationRepo interface {
  CreateDestination(models.Destination) (error)
}
