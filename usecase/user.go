package usecase

import (
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/fgunawan1995/lemonilo/util"
	"github.com/pkg/errors"
)

//GetUserByID get single data of user
func (u *impl) GetUserByID(userID string) (model.User, error) {
	user, err := u.dbDAL.GetUserByID(userID)
	if err != nil {
		return user, errors.WithStack(err)
	}
	return user, nil
}

//InsertUser insert single data to table 'users'
func (u *impl) InsertUser(data model.InsertUser) error {
	err := data.Validate()
	if err != nil {
		return errors.WithStack(err)
	}
	err = util.WithTransaction(u.dbDAL.GetDB(), func(tx util.Transaction) error {
		err := u.dbDAL.InsertUser(tx, data)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//UpdateUser Update single data to table 'users'
func (u *impl) UpdateUser(data model.UpdateUser) error {
	err := data.Validate()
	if err != nil {
		return errors.WithStack(err)
	}
	err = util.WithTransaction(u.dbDAL.GetDB(), func(tx util.Transaction) error {
		err := u.dbDAL.UpdateUser(tx, data)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//DeleteUser non-activate single user to table 'users'
func (u *impl) DeleteUser(userID string) error {
	return util.WithTransaction(u.dbDAL.GetDB(), func(tx util.Transaction) error {
		err := u.dbDAL.DeleteUser(tx, userID)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}
