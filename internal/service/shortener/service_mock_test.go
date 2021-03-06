// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/shortener/service.go

// Package shortener is a generated GoMock package.
package shortener

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/junnotantra/go-shortener/internal/types"
	reflect "reflect"
)

// MockResource is a mock of Resource interface
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
}

// MockResourceMockRecorder is the mock recorder for MockResource
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// SaveShortURL mocks base method
func (m *MockResource) SaveShortURL(shortURL types.ShortURL, allowUpdate bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveShortURL", shortURL, allowUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveShortURL indicates an expected call of SaveShortURL
func (mr *MockResourceMockRecorder) SaveShortURL(shortURL, allowUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveShortURL", reflect.TypeOf((*MockResource)(nil).SaveShortURL), shortURL, allowUpdate)
}

// GetShortURLInfo mocks base method
func (m *MockResource) GetShortURLInfo(uniqueString string) (types.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortURLInfo", uniqueString)
	ret0, _ := ret[0].(types.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShortURLInfo indicates an expected call of GetShortURLInfo
func (mr *MockResourceMockRecorder) GetShortURLInfo(uniqueString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortURLInfo", reflect.TypeOf((*MockResource)(nil).GetShortURLInfo), uniqueString)
}
