// Code generated by MockGen. DO NOT EDIT.
// Source: controllers/refresolver/refresolver.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=controllers/mocks/refresolver_mock.go -source=controllers/refresolver/refresolver.go GslbRefResolver
//

// Package mocks is a generated GoMock package.
package mocks

/*
Copyright 2022 The k8gb Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	reflect "reflect"

	v1beta1 "github.com/k8gb-io/k8gb/api/v1beta1"
	utils "github.com/k8gb-io/k8gb/controllers/utils"
	gomock "go.uber.org/mock/gomock"
)

// MockGslbReferenceResolver is a mock of GslbReferenceResolver interface.
type MockGslbReferenceResolver struct {
	ctrl     *gomock.Controller
	recorder *MockGslbReferenceResolverMockRecorder
}

// MockGslbReferenceResolverMockRecorder is the mock recorder for MockGslbReferenceResolver.
type MockGslbReferenceResolverMockRecorder struct {
	mock *MockGslbReferenceResolver
}

// NewMockGslbReferenceResolver creates a new mock instance.
func NewMockGslbReferenceResolver(ctrl *gomock.Controller) *MockGslbReferenceResolver {
	mock := &MockGslbReferenceResolver{ctrl: ctrl}
	mock.recorder = &MockGslbReferenceResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGslbReferenceResolver) EXPECT() *MockGslbReferenceResolverMockRecorder {
	return m.recorder
}

// GetGslbExposedIPs mocks base method.
func (m *MockGslbReferenceResolver) GetGslbExposedIPs(arg0 utils.DNSList) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGslbExposedIPs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGslbExposedIPs indicates an expected call of GetGslbExposedIPs.
func (mr *MockGslbReferenceResolverMockRecorder) GetGslbExposedIPs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGslbExposedIPs", reflect.TypeOf((*MockGslbReferenceResolver)(nil).GetGslbExposedIPs), arg0)
}

// GetServers mocks base method.
func (m *MockGslbReferenceResolver) GetServers() ([]*v1beta1.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServers")
	ret0, _ := ret[0].([]*v1beta1.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServers indicates an expected call of GetServers.
func (mr *MockGslbReferenceResolverMockRecorder) GetServers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServers", reflect.TypeOf((*MockGslbReferenceResolver)(nil).GetServers))
}
