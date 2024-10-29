package repository

import (
	"beef/internal/models"
)

type BeefRepository interface {
	FindBeef(input map[string]uint32) (*models.ModelBeef, error)
}
