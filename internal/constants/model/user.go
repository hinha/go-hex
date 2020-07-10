package model

// User models
type User struct {
	ID        string `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"hash_password" json:"password,omitempty"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	LastLogin int64  `db:"last_login" json:"last_login"`
	Status    int8   `db:"status" json:"status"`
}

type Token struct {
	UniqueToken string `db:"unique_token" json:"unique_token"`
	TimeAt		string `db:"time_at" json:"time_at"`
}

// RequestRegister for error handler
type RequestRegister struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}