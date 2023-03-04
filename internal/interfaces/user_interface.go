package interfaces

import (
	"net/http"

	"template/internal/model"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type UserService interface {
	CreateUser(name, email, password string) (*model.User, error)
	GetUser(id int) (*model.User, error)
}

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUser(id int) (*model.User, error)
}
