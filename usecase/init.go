package usecase

import (
	"github.com/fgunawan1995/lemonilo/config"
	cachedal "github.com/fgunawan1995/lemonilo/dal/cache"
	dbdal "github.com/fgunawan1995/lemonilo/dal/db"
	"github.com/fgunawan1995/lemonilo/model"
)

type impl struct {
	cfg      *config.Config
	dbDAL    dbdal.DBDAL
	cacheDAL cachedal.CacheDAL
}

type Usecase interface {
	//GetUserByID get single data of user
	GetUserByID(userID string) (model.User, error)
	//InsertUser insert single data to table 'users'
	InsertUser(data model.InsertUser) error
	//UpdateUser Update single data to table 'users'
	UpdateUser(data model.UpdateUser) error
	//DeleteUser non-activate single user to table 'users'
	DeleteUser(userID string) error
	//Login login process, check user from db, and save token
	Login(data model.Login) (string, error)
}

func New(cfg *config.Config, db dbdal.DBDAL, cache cachedal.CacheDAL) Usecase {
	return &impl{
		cfg:      cfg,
		dbDAL:    db,
		cacheDAL: cache,
	}
}
