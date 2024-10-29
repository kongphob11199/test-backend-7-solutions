package service

import "beef/internal/models"

type BeefService interface {
	FindBeef() (*models.ModelBeef, error)
}
