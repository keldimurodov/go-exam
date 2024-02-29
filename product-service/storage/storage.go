package storage

import (
	"go-exam/product-service/storage/postgres"
	"go-exam/product-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	db       *sqlx.DB
	productRepo repo.ProductStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		productRepo: postgres.NewProductRepo(db),
	}
}

func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
