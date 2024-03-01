package repo

import (
	pbu "go-exam/user-service/genproto/user"
)

// UserStorageI ...
type UserStoragePostgresI interface {
	Create(*pbu.User) (*pbu.User, error)
	Get(user *pbu.GetUserRequest) (*pbu.User, error)
	GetAll(req *pbu.GetAllRequest) (*pbu.GetAllResponse, error)
	Delete(user *pbu.GetUserRequest) (*pbu.User, error)
	Update(user *pbu.User) (*pbu.User, error)
	CheckUniqueness(req *pbu.CheckUniquenessRequest) (*pbu.CheckUniquenessResponse, error)
	Login(req *pbu.LoginRequest) (*pbu.User, error)
}
