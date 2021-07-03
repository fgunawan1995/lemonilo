package model

import (
	"errors"
	"reflect"
	"testing"
)

func TestBuildAPIResponseError(t *testing.T) {
	type args struct {
		statusCode int
		err        error
	}
	tests := []struct {
		name string
		args args
		want GeneralAPIResponse
	}{
		{
			name: "success",
			args: args{
				err: errors.New("aaa"),
			},
			want: GeneralAPIResponse{
				Error: "aaa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAPIResponseError(tt.args.statusCode, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAPIResponseError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildAPIResponseSuccess(t *testing.T) {
	type args struct {
		statusCode int
		data       interface{}
	}
	tests := []struct {
		name string
		args args
		want GeneralAPIResponse
	}{
		{
			name: "success",
			args: args{
				data: "aaa",
			},
			want: GeneralAPIResponse{
				Data: "aaa",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAPIResponseSuccess(tt.args.statusCode, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAPIResponseSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
