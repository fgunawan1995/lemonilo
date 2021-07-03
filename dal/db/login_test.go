package db

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgunawan1995/lemonilo/model"
	"github.com/jmoiron/sqlx"
)

func Test_impl_GetUserByUserID(t *testing.T) {
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
			got, err := dal.GetUserByUserID(tt.args.userID)
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

func Test_impl_GetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type args struct {
		email string
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
			got, err := dal.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
