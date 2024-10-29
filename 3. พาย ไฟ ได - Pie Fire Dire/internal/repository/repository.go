package repository

type Repositories struct {
	Beef BeefRepository
}

func NewRepository() *Repositories {
	return &Repositories{
		Beef: NewRepositoryBeef(),
	}
}
