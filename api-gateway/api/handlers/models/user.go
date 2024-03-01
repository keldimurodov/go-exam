package models

// type User struct {
// 	id string  `json:"id"`
// 	first_name string `json:"first_name"`
// 	last_name string `json:"last_name"`
// 	picture string `json:"picture"`
// 	bio string `json:"bio"`
// 	email string `json:"email"`
// 	password string `json:"password"`
// 	created_at string `json:"created_at"`
// 	updeted_at string `json:"updated_at"`
// 	deleted_at string `json:"deleted_at"`
// }

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	Users []*User `json:"users"`
}
