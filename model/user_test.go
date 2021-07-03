package model

import (
	"testing"
)

func TestInsertUser_Validate(t *testing.T) {
	type fields struct {
		UserID   string
		Email    string
		Address  string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				UserID:   "aaabbb",
				Email:    "aaa@gmail.com",
				Address:  "aaa",
				Password: "aaabbbccc",
			},
		},
		{
			name: "error required",
			fields: fields{
				UserID: "aaa",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := InsertUser{
				UserID:   tt.fields.UserID,
				Email:    tt.fields.Email,
				Address:  tt.fields.Address,
				Password: tt.fields.Password,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUser_Validate(t *testing.T) {
	type fields struct {
		ID      string
		Email   string
		Address string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ID:      "1",
				Email:   "aaa@gmail.com",
				Address: "aaa",
			},
		},
		{
			name: "error required",
			fields: fields{
				Email: "aaa@gmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := UpdateUser{
				ID:      tt.fields.ID,
				Email:   tt.fields.Email,
				Address: tt.fields.Address,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
