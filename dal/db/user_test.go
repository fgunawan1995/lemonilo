package db

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgunawan1995/lemonilo/mocks"
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/fgunawan1995/lemonilo/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

func Test_impl_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
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
				rows := sqlmock.NewRows([]string{"id", "user_id", "email", "password", "address", "is_active"}).
					AddRow(1, "ABC", "abc", "aaa", "bbb", true)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: model.User{
				ID:       "1",
				UserID:   "ABC",
				Email:    "abc",
				Password: "aaa",
				Address:  "bbb",
				IsActive: true,
			},
		},
		{
			name: "error",
			mock: func() {
				mock.ExpectQuery("(.*?)").WillReturnError(errors.New("db error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: sqlx.NewDb(db, "sqlmock"),
			}
			tt.mock()
			got, err := dal.GetUserByID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetUserByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetUserByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_InsertUser(t *testing.T) {
	mockTx := new(mocks.Transaction)
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		tx   util.Transaction
		data model.InsertUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Times(1)
			},
		},
		{
			name: "error",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("db error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: tt.fields.db,
			}
			tt.mock()
			if err := dal.InsertUser(mockTx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_UpdateUser(t *testing.T) {
	mockTx := new(mocks.Transaction)
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		tx   util.Transaction
		data model.UpdateUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Times(1)
			},
		},
		{
			name: "error",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("db error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: tt.fields.db,
			}
			tt.mock()
			if err := dal.UpdateUser(mockTx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_DeleteUser(t *testing.T) {
	mockTx := new(mocks.Transaction)
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		tx   util.Transaction
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything).Return(nil, nil).Times(1)
			},
		},
		{
			name: "error",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything).Return(nil, errors.New("db error")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: tt.fields.db,
			}
			tt.mock()
			if err := dal.DeleteUser(mockTx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
