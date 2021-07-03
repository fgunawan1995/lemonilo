package model

import "github.com/go-playground/validator"

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	ID       string `db:"id" json:"id"`
	UserID   string `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Address  string `db:"address" json:"address"`
	Password string `db:"password" json:"-"`
	IsActive bool   `db:"is_active" json:"-"`
}

type InsertUser struct {
	UserID   string `json:"user_id" validate:"required,min=5,max=40,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address" validate:"required"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}

func (p InsertUser) Validate() error {
	return validate.Struct(p)
}

type UpdateUser struct {
	ID      string `db:"id" json:"id" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (p UpdateUser) Validate() error {
	return validate.Struct(p)
}
