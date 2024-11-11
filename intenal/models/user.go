package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"_"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
}

type LoginUser struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type RegisterUser struct {
	Username  string `db:"username" json:"username"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Password  string `db:"password" json:"password"`
	Age       int    `db:"age" json:"age"`
	Email     string `db:"email" json:"email"`
}

type ResetPassword struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
