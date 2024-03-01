package postgres

import (
	"fmt"
	pbu "go-exam/user-service/genproto/user"
	"log"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pbu.User) (*pbu.User, error) {

	var res pbu.User
	query := `
		INSERT INTO users(
			first_name, 
			last_name, 
			email,
			password
		)
		VALUES ($1, $2, $3, $4) 
		RETURNING 
			id, 
			first_name, 
			last_name,
			email,
			password,
			created_at,
			updeted_at
	`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Password,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		fmt.Println("err 2")
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) Get(id *pbu.GetUserRequest) (*pbu.User, error) {
	var user pbu.User
	query := `
	SELECT
		id,
		first_name,
		last_name,
		email,
		password,
		created_at,
		updeted_at
	FROM 
		users
	WHERE
		id=$1 
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, id.UserId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Update(user *pbu.User) (*pbu.User, error) {
	var res pbu.User
	query := `
	UPDATE
		users
	SET
		first_name=$1,
		last_name=$2,
		updeted_at=CURRENT_TIMESTAMP
	WHERE
		id=$3
	returning
		id, 
		first_name,
		last_name,
		created_at,
		updeted_at
	`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) Delete(user *pbu.GetUserRequest) (*pbu.User, error) {
	var res pbu.User
	query := `
	UPDATE
		users
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	RETURNING
		id, 
		first_name, 
		last_name,
		created_at,
		updeted_at,
		deleted_at
	`
	err := r.db.QueryRow(query, user.UserId).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (r *userRepo) GetAll(user *pbu.GetAllRequest) (*pbu.GetAllResponse, error) {
	var allUser pbu.GetAllResponse
	query := `
	SELECT
		id,
		first_name,
		last_name,
		created_at,
		updeted_at
	FROM 
		users 
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`
	offset := user.Limit * (user.Page - 1)
	rows, err := r.db.Query(query, user.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user pbu.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&user.UpdetedAt)
		if err != nil {
			return nil, err
		}
		allUser.Users = append(allUser.Users, &user)
	}
	return &allUser, nil
}

func (s *userRepo) CheckUniqueness(req *pbu.CheckUniquenessRequest) (*pbu.CheckUniquenessResponse, error) {
	var email int

	fmt.Println(req.Field, req.Value)

	query := fmt.Sprintf("SELECT count(1) from users WHERE %s = $1 ", req.Field)
	err := s.db.QueryRow(query, req.Value).Scan(&email)
	if err != nil {
		log.Fatal("error while checking!!!", err.Error())
	}
	if email == 1 {

		return &pbu.CheckUniquenessResponse{
			Result: true,
		}, nil
	}

	return &pbu.CheckUniquenessResponse{
		Result: false,
	}, nil
}

func SubVerification(db *sqlx.DB, user *pbu.UserDetail) (*pbu.User, error) {

	var res pbu.User

	query := `
		INSERT INTO users(
			first_name, 
			last_name, 
			email,
			password
		)
		VALUES ($1, $2, $3, $4) 
		RETURNING 
			id, 
			first_name, 
			last_name,
			email,
			password,
			created_at,
			updeted_at
	`
	err := db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Password,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		fmt.Println("error subvalidation")
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) Login(id *pbu.LoginRequest) (*pbu.User, error) {
	var user pbu.User
	query := `
	SELECT
		id,
		first_name,
		last_name,
		email,
		password,
		created_at,
		updeted_at
	FROM 
		users
	WHERE
		email=$1
	AND 
		password = $2
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, id.Email, id.Password).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
