package user

import (
	"sirclo/gql/entities"
)

type RepositoryUser interface {
	Authenticate(name, password string) bool
	GetUsers() ([]entities.User, error)
	GetUser(id int) (entities.User, error)
	// GetUserIdByUsername(name string) (int, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	DeleteUser(user entities.User) (entities.User, error)
}
