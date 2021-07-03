package model

import "testing"

func TestLogin_Validate(t *testing.T) {
	type fields struct {
		UserIDOrEmail string
		Password      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				UserIDOrEmail: "aaa",
				Password:      "aaa",
			},
		},
		{
			name: "error required",
			fields: fields{
				UserIDOrEmail: "aaa",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Login{
				UserIDOrEmail: tt.fields.UserIDOrEmail,
				Password:      tt.fields.Password,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Login.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
