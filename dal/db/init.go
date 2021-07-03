package db

import (
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/fgunawan1995/lemonilo/util"
	"github.com/jmoiron/sqlx"
)

type DBDAL interface {
	//GetUserByID from db and return single data of user
	GetUserByID(userID string) (model.User, error)
	//InsertUser insert single data to table 'users'
	InsertUser(tx util.Transaction, data model.InsertUser) error
	//UpdateUser Update single data to table 'users'
	UpdateUser(tx util.Transaction, data model.UpdateUser) error
	//DeleteUser non-activate single user to table 'users'
	DeleteUser(tx util.Transaction, userID string) error
	//GetUserByUserID used for login process, get user by user id
	GetUserByUserID(userID string) (model.User, error)
	//GetUserByEmail used for login process, get user by email
	GetUserByEmail(email string) (model.User, error)
	//GetDB for used withing util.WithTransaction
	GetDB() *sqlx.DB
}

type impl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) DBDAL {
	return &impl{
		db: db,
	}
}

//GetDB for used withing util.WithTransaction
func (dal *impl) GetDB() *sqlx.DB {
	return dal.db
}
