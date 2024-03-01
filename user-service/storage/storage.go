package storage

import (
	"go-exam/user-service/storage/postgres"
	"go-exam/user-service/storage/repo"
	"go-exam/user-service/storage/sredis"

	// "github.com/redis/go-redis/v9"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// IStorage ...
type IStorage interface {
	User() repo.UserStoragePostgresI
}

type IStorageRedis interface {
	UserRedis() repo.UserStorageRedisI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStoragePostgresI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStoragePostgresI {
	return s.userRepo
}

// Redis

// storageRedis struct implements IStorage for Redis
type storageRedis struct {
	db            *redis.Client
	userRepoRedis repo.UserStorageRedisI
}

// NewStorageRedis creates a new instance of storageRedis
func NewStorageRedis(db *redis.Client) *storageRedis {
	// Initialize Redis-specific fields here
	return &storageRedis{
		db:            db,
		userRepoRedis: sredis.NewRedisRepo(db), // Assuming you have a Redis-specific user repository
	}
}

// User method returns the Redis user repository
func (s storageRedis) UserRedis() repo.UserStorageRedisI {
	return s.userRepoRedis
}
