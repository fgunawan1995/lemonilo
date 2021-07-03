package db

import (
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/pkg/errors"
)

//GetUserByUserID used for login process, get user by user id
func (dal *impl) GetUserByUserID(userID string) (model.User, error) {
	var result model.User
	err := dal.db.Get(&result, getUserByUserID, userID)
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}

//GetUserByEmail used for login process, get user by email
func (dal *impl) GetUserByEmail(email string) (model.User, error) {
	var result model.User
	err := dal.db.Get(&result, getUserByEmail, email)
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}
