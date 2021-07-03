package db

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNew(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want DBDAL
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.db)
		})
	}
}

func Test_impl_GetDB(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   *sqlx.DB
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: tt.fields.db,
			}
			if got := dal.GetDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
