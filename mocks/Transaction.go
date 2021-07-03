// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// Exec provides a mock function with given fields: query, args
func (_m *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(string, ...interface{}) sql.Result); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}