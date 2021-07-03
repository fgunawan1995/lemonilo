package usecase

import (
	"errors"
	"testing"

	"github.com/fgunawan1995/lemonilo/config"
	"github.com/fgunawan1995/lemonilo/mocks"
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/stretchr/testify/mock"
)

func Test_impl_Login(t *testing.T) {
	mCache := new(mocks.CacheDAL)
	mDB := new(mocks.DBDAL)
	mCfg := &config.Config{
		Server: config.ServerConfig{
			Secret: "secret_sample",
		},
	}
	type args struct {
		data model.Login
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success, by user_id",
			args: args{
				data: model.Login{
					UserIDOrEmail: "tester",
					Password:      "testpassword",
				},
			},
			mock: func() {
				mDB.On("GetUserByUserID", mock.Anything).Return(model.User{
					UserID:   "tester",
					Password: "$2a$10$LjIQStxh0z/7o1JcGJrJ6eAQXvmEDYPGqS9TBdXxwhQuK/VTrXs7a",
				}, nil).Times(1)
				mCache.On("SetUserToken", mock.Anything, mock.Anything).Return(nil).Times(1)
			},
		},
		{
			name: "success, by email",
			args: args{
				data: model.Login{
					UserIDOrEmail: "tester",
					Password:      "testpassword",
				},
			},
			mock: func() {
				mDB.On("GetUserByUserID", mock.Anything).Return(model.User{}, errors.New("no rows")).Times(1)
				mDB.On("GetUserByEmail", mock.Anything).Return(model.User{
					UserID:   "tester",
					Password: "$2a$10$LjIQStxh0z/7o1JcGJrJ6eAQXvmEDYPGqS9TBdXxwhQuK/VTrXs7a",
				}, nil).Times(1)
				mCache.On("SetUserToken", mock.Anything, mock.Anything).Return(nil).Times(1)
			},
		},
		{
			name: "error not found",
			args: args{
				data: model.Login{
					UserIDOrEmail: "tester",
					Password:      "testpassword",
				},
			},
			mock: func() {
				mDB.On("GetUserByUserID", mock.Anything).Return(model.User{}, errors.New("no rows")).Times(1)
				mDB.On("GetUserByEmail", mock.Anything).Return(model.User{}, errors.New("no rows")).Times(1)
				mCache.On("SetUserToken", mock.Anything, mock.Anything).Return(nil).Times(1)
			},
			wantErr: true,
		},
		{
			name: "error invalid password",
			args: args{
				data: model.Login{
					UserIDOrEmail: "tester",
					Password:      "testpasswor",
				},
			},
			mock: func() {
				mDB.On("GetUserByUserID", mock.Anything).Return(model.User{
					UserID:   "tester",
					Password: "$2a$10$LjIQStxh0z/7o1JcGJrJ6eAQXvmEDYPGqS9TBdXxwhQuK/VTrXs7a",
				}, nil).Times(1)
				mCache.On("SetUserToken", mock.Anything, mock.Anything).Return(nil).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				dbDAL:    mDB,
				cacheDAL: mCache,
				cfg:      mCfg,
			}
			tt.mock()
			_, err := u.Login(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
