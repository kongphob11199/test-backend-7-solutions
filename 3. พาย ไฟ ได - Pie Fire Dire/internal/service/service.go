package service

import "beef/internal/repository"

type Service struct {
	Beef BeefService
}

type Deps struct {
	Repository *repository.Repositories
}

func NewService(deps Deps) *Service {
	return &Service{
		Beef: NewServiceBeef(deps.Repository.Beef),
	}
}
