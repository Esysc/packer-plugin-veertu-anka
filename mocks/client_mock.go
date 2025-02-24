// Code generated by MockGen. DO NOT EDIT.
// Source: client/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/veertuinc/packer-plugin-veertu-anka/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockClient) Clone(params client.CloneParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockClientMockRecorder) Clone(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockClient)(nil).Clone), params)
}

// Copy mocks base method.
func (m *MockClient) Copy(params client.CopyParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockClientMockRecorder) Copy(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockClient)(nil).Copy), params)
}

// Create mocks base method.
func (m *MockClient) Create(params client.CreateParams, outputStreamer chan string) (client.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", params, outputStreamer)
	ret0, _ := ret[0].(client.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClientMockRecorder) Create(params, outputStreamer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClient)(nil).Create), params, outputStreamer)
}

// Delete mocks base method.
func (m *MockClient) Delete(params client.DeleteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), params)
}

// Describe mocks base method.
func (m *MockClient) Describe(vmName string) (client.DescribeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Describe", vmName)
	ret0, _ := ret[0].(client.DescribeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Describe indicates an expected call of Describe.
func (mr *MockClientMockRecorder) Describe(vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockClient)(nil).Describe), vmName)
}

// Exists mocks base method.
func (m *MockClient) Exists(vmName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", vmName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockClientMockRecorder) Exists(vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockClient)(nil).Exists), vmName)
}

// FuseAvailable mocks base method.
func (m *MockClient) FuseAvailable(vmName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FuseAvailable", vmName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FuseAvailable indicates an expected call of FuseAvailable.
func (mr *MockClientMockRecorder) FuseAvailable(vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FuseAvailable", reflect.TypeOf((*MockClient)(nil).FuseAvailable), vmName)
}

// License mocks base method.
func (m *MockClient) License() (client.LicenseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "License")
	ret0, _ := ret[0].(client.LicenseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// License indicates an expected call of License.
func (mr *MockClientMockRecorder) License() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "License", reflect.TypeOf((*MockClient)(nil).License))
}

// Modify mocks base method.
func (m *MockClient) Modify(vmName, command, property string, flags ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{vmName, command, property}
	for _, a := range flags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Modify", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Modify indicates an expected call of Modify.
func (mr *MockClientMockRecorder) Modify(vmName, command, property interface{}, flags ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{vmName, command, property}, flags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockClient)(nil).Modify), varargs...)
}

// RegistryList mocks base method.
func (m *MockClient) RegistryList(registryParams client.RegistryParams) ([]client.RegistryListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryList", registryParams)
	ret0, _ := ret[0].([]client.RegistryListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegistryList indicates an expected call of RegistryList.
func (mr *MockClientMockRecorder) RegistryList(registryParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryList", reflect.TypeOf((*MockClient)(nil).RegistryList), registryParams)
}

// RegistryListRepos mocks base method.
func (m *MockClient) RegistryListRepos() (client.RegistryListReposResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryListRepos")
	ret0, _ := ret[0].(client.RegistryListReposResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegistryListRepos indicates an expected call of RegistryListRepos.
func (mr *MockClientMockRecorder) RegistryListRepos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryListRepos", reflect.TypeOf((*MockClient)(nil).RegistryListRepos))
}

// RegistryListReposArm64 mocks base method.
func (m *MockClient) RegistryListReposArm64() (client.RegistryListReposResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryListReposArm64")
	ret0, _ := ret[0].(client.RegistryListReposResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegistryListReposArm64 indicates an expected call of RegistryListReposArm64.
func (mr *MockClientMockRecorder) RegistryListReposArm64() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryListReposArm64", reflect.TypeOf((*MockClient)(nil).RegistryListReposArm64))
}

// RegistryPull mocks base method.
func (m *MockClient) RegistryPull(registryParams client.RegistryParams, pullParams client.RegistryPullParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryPull", registryParams, pullParams)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistryPull indicates an expected call of RegistryPull.
func (mr *MockClientMockRecorder) RegistryPull(registryParams, pullParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryPull", reflect.TypeOf((*MockClient)(nil).RegistryPull), registryParams, pullParams)
}

// RegistryPush mocks base method.
func (m *MockClient) RegistryPush(registryParams client.RegistryParams, pushParams client.RegistryPushParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryPush", registryParams, pushParams)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistryPush indicates an expected call of RegistryPush.
func (mr *MockClientMockRecorder) RegistryPush(registryParams, pushParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryPush", reflect.TypeOf((*MockClient)(nil).RegistryPush), registryParams, pushParams)
}

// RegistryRevert mocks base method.
func (m *MockClient) RegistryRevert(url, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegistryRevert", url, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegistryRevert indicates an expected call of RegistryRevert.
func (mr *MockClientMockRecorder) RegistryRevert(url, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistryRevert", reflect.TypeOf((*MockClient)(nil).RegistryRevert), url, id)
}

// Run mocks base method.
func (m *MockClient) Run(params client.RunParams) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", params)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockClientMockRecorder) Run(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockClient)(nil).Run), params)
}

// Show mocks base method.
func (m *MockClient) Show(vmName string) (client.ShowResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", vmName)
	ret0, _ := ret[0].(client.ShowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show.
func (mr *MockClientMockRecorder) Show(vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockClient)(nil).Show), vmName)
}

// Start mocks base method.
func (m *MockClient) Start(params client.StartParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockClientMockRecorder) Start(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClient)(nil).Start), params)
}

// Stop mocks base method.
func (m *MockClient) Stop(params client.StopParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockClientMockRecorder) Stop(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop), params)
}

// Suspend mocks base method.
func (m *MockClient) Suspend(params client.SuspendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspend", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Suspend indicates an expected call of Suspend.
func (mr *MockClientMockRecorder) Suspend(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspend", reflect.TypeOf((*MockClient)(nil).Suspend), params)
}

// UpdateAddons mocks base method.
func (m *MockClient) UpdateAddons(vmName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddons", vmName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddons indicates an expected call of UpdateAddons.
func (mr *MockClientMockRecorder) UpdateAddons(vmName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddons", reflect.TypeOf((*MockClient)(nil).UpdateAddons), vmName)
}

// Version mocks base method.
func (m *MockClient) Version() (client.VersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(client.VersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockClientMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockClient)(nil).Version))
}
