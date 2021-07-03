package db

import (
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/fgunawan1995/lemonilo/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//GetUserByID from db and return single data of user
func (dal *impl) GetUserByID(userID string) (model.User, error) {
	var result model.User
	err := dal.db.Get(&result, getUserByID, userID)
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil
}

//InsertUser insert single data to table 'users'
func (dal *impl) InsertUser(tx util.Transaction, data model.InsertUser) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "hash password failed")
	}
	_, err = tx.Exec(insertUser, data.UserID, data.Email, data.Address, hashedPassword)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//UpdateUser Update single data to table 'users'
func (dal *impl) UpdateUser(tx util.Transaction, data model.UpdateUser) error {
	_, err := tx.Exec(updateUser, data.ID, data.Email, data.Address)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//DeleteUser non-activate single user to table 'users'
func (dal *impl) DeleteUser(tx util.Transaction, userID string) error {
	_, err := tx.Exec(deleteUser, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
