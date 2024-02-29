package postgres

import (
	pb "go-exam/product-service/genproto/product"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type proRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewProductRepo(db *sqlx.DB) *proRepo {
	return &proRepo{db: db}
}

func (r *proRepo) Create(product *pb.Product) (*pb.Product, error) {

	id := uuid.NewString()
	product.Id = id


	var res pb.Product

	query := `INSERT INTO product(
		id, 
		product_name, 
		product_price, 
		product_about) 
		VALUES ($1, $2, $3, $4) 
		RETURNING 
		id, 
		product_name, 
		product_price, 
		product_about,
		created_at,
		updeted_at,
		deleted_at `
	err := r.db.QueryRow(query, product.Id, product.ProductName, product.ProductPrice, product.ProductAbout).Scan(
		&res.Id,
		&res.ProductName,
		&res.ProductPrice,
		&res.ProductAbout,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *proRepo) Get(id *pb.GetProductRequest) (*pb.Product, error) {
	var product pb.Product

	query := `SELECT
	 	id, 
		product_name,
		product_price,
		product_about,
		created_at,
		updeted_at,
		deleted_at
		from product where id=$1`

	err := r.db.QueryRow(query, id.Id).Scan(
		&product.Id,
		&product.ProductName,
		&product.ProductPrice,
		&product.ProductAbout,
		&product.CreatedAt,
		&product.UpdetedAt,
		&product.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *proRepo) GetAll(product *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	var allProduct pb.GetAllResponse
	query := `
	SELECT
		id,
		product_name,
		product_price,
		product_about,
		created_at,
		updeted_at,
		deleted_at
	FROM 
		product
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`

	offset := product.Limit * (product.Page - 1)

	rows, err := r.db.Query(query, product.Limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pr pb.Product

		err := rows.Scan(
			&pr.Id,
			&pr.ProductName,
			&pr.ProductPrice,
			&pr.ProductAbout,
			&pr.CreatedAt,
			&pr.UpdetedAt,
			&pr.DeletedAt)

		if err != nil {
			return nil, err
		}

		allProduct.Products = append(allProduct.Products, &pr)
	}
	return &allProduct, nil
}

func (r *proRepo) Update(product *pb.Product) (*pb.Product, error) {
	var res pb.Product

	query := `
	UPDATE
		product
	SET
		product_name=$1,
		product_price=$2,
		product_about=$3,
		updeted_at=CURRENT_TIMESTAMP
	WHERE
		id=$4
	returning
		id, 
		product_name,
		product_price,
		product_about,
		created_at,
		updeted_at,
		deleted_at
	`
	err := r.db.QueryRow(
		query,
		product.ProductName,
		product.ProductPrice,
		product.ProductAbout,
		product.Id).Scan(
		&res.Id,
		&res.ProductName,
		&res.ProductPrice,
		&res.ProductAbout,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *proRepo) Delete(pr *pb.GetProductRequest) (*pb.Product, error) {
	
	var res pb.Product

	query := `
	UPDATE
		product
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	`
	err := r.db.QueryRow(query, pr.Id).Scan(
		&res.Id,
		&res.ProductName,
		&res.ProductPrice,
		&res.ProductAbout,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}
