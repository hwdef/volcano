/*
Copyright 2024 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCNIPluginsConfHandler is a mock of CNIPluginsConfHandler interface.
type MockCNIPluginsConfHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCNIPluginsConfHandlerMockRecorder
}

// MockCNIPluginsConfHandlerMockRecorder is the mock recorder for MockCNIPluginsConfHandler.
type MockCNIPluginsConfHandlerMockRecorder struct {
	mock *MockCNIPluginsConfHandler
}

// NewMockCNIPluginsConfHandler creates a new mock instance.
func NewMockCNIPluginsConfHandler(ctrl *gomock.Controller) *MockCNIPluginsConfHandler {
	mock := &MockCNIPluginsConfHandler{ctrl: ctrl}
	mock.recorder = &MockCNIPluginsConfHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCNIPluginsConfHandler) EXPECT() *MockCNIPluginsConfHandlerMockRecorder {
	return m.recorder
}

// AddOrUpdateCniPluginToConfList mocks base method.
func (m *MockCNIPluginsConfHandler) AddOrUpdateCniPluginToConfList(configPath, name string, plugin map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateCniPluginToConfList", configPath, name, plugin)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrUpdateCniPluginToConfList indicates an expected call of AddOrUpdateCniPluginToConfList.
func (mr *MockCNIPluginsConfHandlerMockRecorder) AddOrUpdateCniPluginToConfList(configPath, name, plugin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateCniPluginToConfList", reflect.TypeOf((*MockCNIPluginsConfHandler)(nil).AddOrUpdateCniPluginToConfList), configPath, name, plugin)
}

// DeleteCniPluginFromConfList mocks base method.
func (m *MockCNIPluginsConfHandler) DeleteCniPluginFromConfList(configPath, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCniPluginFromConfList", configPath, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCniPluginFromConfList indicates an expected call of DeleteCniPluginFromConfList.
func (mr *MockCNIPluginsConfHandlerMockRecorder) DeleteCniPluginFromConfList(configPath, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCniPluginFromConfList", reflect.TypeOf((*MockCNIPluginsConfHandler)(nil).DeleteCniPluginFromConfList), configPath, name)
}