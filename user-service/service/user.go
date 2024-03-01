package service

import (
	"context"
	"database/sql"
	"fmt"
	pbu "go-exam/user-service/genproto/user"
	l "go-exam/user-service/pkg/logger"
	"go-exam/user-service/storage"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// UserService ...
type UserService struct {
	storage      storage.IStorage
	storageRedis storage.IStorageRedis
	logger       l.Logger
	db           *sql.DB
	rdb          *redis.Client
	// rdb redis.Client
}

// NewUserService ...
func NewUserService(db *sqlx.DB, rdb *redis.Client, log l.Logger) *UserService {
	return &UserService{
		storage:      storage.NewStoragePg(db),
		storageRedis: storage.NewStorageRedis(rdb),
		logger:       log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *pbu.GetUserRequest) (*pbu.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().Update(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pbu.GetUserRequest) (*pbu.User, error) {
	user, err := s.storage.User().Delete(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, req *pbu.GetAllRequest) (*pbu.GetAllResponse, error) {
	users, err := s.storage.User().GetAll(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CheckUniqueness(ctx context.Context, req *pbu.CheckUniquenessRequest) (*pbu.CheckUniquenessResponse, error) {
	user, err := s.storage.User().CheckUniqueness(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Sign(ctx context.Context, user *pbu.UserDetail) (*pbu.ResponseMessage, error) {
	req, err := s.storageRedis.UserRedis().Sign(user)
	if err != nil {
		fmt.Println("big service error")
		return nil, err

	}
	return req, nil
}

func (s *UserService) Verification(ctx context.Context, req *pbu.VerificationUserRequest) (*pbu.User, error) {
	user, err := s.storageRedis.UserRedis().Verification(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, req *pbu.LoginRequest) (*pbu.User, error) {
	user, err := s.storage.User().Login(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
