package models

type Product struct {
	Id           string `json:"id"`
	ProductName  string `json:"product_name"`
	ProductPrice int64  `json:"product_price"`
	ProductAbout string `json:"product_about"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
