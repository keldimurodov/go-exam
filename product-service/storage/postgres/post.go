package postgres

import (
	pb "go-exam/product-service/genproto/product"
	"log"

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
		updeted_at`
	err := r.db.QueryRow(query, product.Id, product.ProductName, product.ProductPrice, product.ProductAbout).Scan(
		&res.Id,
		&res.ProductName,
		&res.ProductPrice,
		&res.ProductAbout,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *proRepo) Get(produ *pb.GetRequest) (*pb.Product, error) {

	var res pb.Product

	query := `SELECT
		id, 
		product_name, 
		product_price, 
		product_about,
		created_at,
		updeted_at, 
		deleted_at 
		FROM product 
		WHERE id = $1`
	err := r.db.QueryRow(query, produ.Id).Scan(
		&res.Id,
		&res.ProductName,
		&res.ProductPrice,
		&res.ProductAbout,
		&res.CreatedAt,
		&res.UpdetedAt,
		)
	if err != nil {
		return nil, err
	}

	return &res, nil
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
		updeted_at
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
			&pr.UpdetedAt)

		if err != nil {
			return nil, err
		}

		allProduct.Products = append(allProduct.Products, &pr)
	}
	return &allProduct, nil
}

func (r *proRepo) Update(prr *pb.Product) (*pb.Product, error) {
	query := `
	UPDATE
		product
	SET
		product_name = $1,
		product_price = $2,
		product_about = $3,
		updeted_at = CURRENT_TIMESTAMP
	WHERE
		id = $4
	RETURNING
	id, 
	product_name, 
	product_price, 
	product_about,
	created_at,
	updeted_at`

	var respUser pb.Product
	err := r.db.QueryRow(query, prr.ProductName, prr.ProductPrice, prr.ProductAbout, prr.Id).Scan(
        &respUser.Id,
        &respUser.ProductName,
        &respUser.ProductPrice,
        &respUser.ProductAbout,
		&respUser.CreatedAt,
        &respUser.UpdetedAt,
	)

	if err != nil {
		log.Println("Error updating user in postgres")
		return nil, err
	}
	return &respUser, nil
}

func (r *proRepo) Delete(pr *pb.GetRequest) (*pb.Product, error) {
	
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
