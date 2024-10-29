package repository

import "beef/internal/models"

type repositoryBeef struct {
}

func NewRepositoryBeef() *repositoryBeef {
	return &repositoryBeef{}
}

func (r *repositoryBeef) FindBeef(input map[string]uint32) (*models.ModelBeef, error) {

	beef := &models.ModelBeef{
		Beef: input,
	}

	return beef, nil
}
