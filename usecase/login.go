package usecase

import (
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/fgunawan1995/lemonilo/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//Login login process, check user from db, and save token
func (u *impl) Login(data model.Login) (string, error) {
	var token string
	user, err := u.dbDAL.GetUserByUserID(data.UserIDOrEmail)
	if err != nil {
		user, err = u.dbDAL.GetUserByEmail(data.UserIDOrEmail)
		if err != nil {
			return token, errors.WithStack(err)
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return token, errors.WithStack(err)
	}
	token, err = util.GenerateToken(u.cfg.Server.Secret)
	if err != nil {
		return token, errors.WithStack(err)
	}
	err = u.cacheDAL.SetUserToken(user.ID, token)
	if err != nil {
		return token, errors.WithStack(err)
	}
	return token, nil
}
