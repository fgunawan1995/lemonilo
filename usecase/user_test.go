package usecase

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgunawan1995/lemonilo/mocks"
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func Test_impl_GetUserByID(t *testing.T) {
	mDB := new(mocks.DBDAL)
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    model.User
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				mDB.On("GetUserByID", mock.Anything).Return(model.User{
					UserID: "tester",
				}, nil).Times(1)
			},
			want: model.User{
				UserID: "tester",
			},
		},
		{
			name: "error",
			mock: func() {
				mDB.On("GetUserByID", mock.Anything).Return(model.User{}, errors.New("db error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				dbDAL: mDB,
			}
			tt.mock()
			got, err := u.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_InsertUser(t *testing.T) {
	mDB := new(mocks.DBDAL)
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		data model.InsertUser
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				data: model.InsertUser{
					UserID:   "abcde",
					Email:    "aaa@gmail.com",
					Password: "abcdefghij",
					Address:  "aaa bbb",
				},
			},
			mock: func() {
				mDB.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				mDB.On("InsertUser", mock.Anything, mock.Anything).Return(nil).Times(1)
				mockSQL.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				data: model.InsertUser{
					UserID:   "abcde",
					Email:    "aaa@gmail.com",
					Password: "abcdefghij",
					Address:  "aaa bbb",
				},
			},
			mock: func() {
				mDB.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				mDB.On("InsertUser", mock.Anything, mock.Anything).Return(errors.New("aaa")).Times(1)
				mockSQL.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "error validate",
			mock: func() {
				return
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				dbDAL: mDB,
			}
			tt.mock()
			if err := u.InsertUser(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_UpdateUser(t *testing.T) {
	mDB := new(mocks.DBDAL)
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		data model.UpdateUser
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				data: model.UpdateUser{
					ID:      "1",
					Email:   "aaa@gmail.com",
					Address: "aaa bbb",
				},
			},
			mock: func() {
				mDB.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				mDB.On("UpdateUser", mock.Anything, mock.Anything).Return(nil).Times(1)
				mockSQL.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				data: model.UpdateUser{
					ID:      "1",
					Email:   "aaa@gmail.com",
					Address: "aaa bbb",
				},
			},
			mock: func() {
				mDB.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				mDB.On("UpdateUser", mock.Anything, mock.Anything).Return(errors.New("aaa")).Times(1)
				mockSQL.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "error validate",
			mock: func() {
				return
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				dbDAL: mDB,
			}
			tt.mock()
			if err := u.UpdateUser(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_DeleteUser(t *testing.T) {
	mDB := new(mocks.DBDAL)
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				data: "1",
			},
			mock: func() {
				mDB.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				mDB.On("DeleteUser", mock.Anything, mock.Anything).Return(nil).Times(1)
				mockSQL.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				data: "1",
			},
			mock: func() {
				mDB.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				mDB.On("DeleteUser", mock.Anything, mock.Anything).Return(errors.New("aaa")).Times(1)
				mockSQL.ExpectRollback()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				dbDAL: mDB,
			}
			tt.mock()
			if err := u.DeleteUser(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
