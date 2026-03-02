package repository

type repository interface {
	Create()
	Get()
	GetByID()
	Delete()
	Add()
}
