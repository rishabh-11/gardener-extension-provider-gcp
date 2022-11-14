// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client (interfaces: Factory,DNSClient,ComputeClient)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	client "github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client"
	gomock "github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	v1 "k8s.io/api/core/v1"
	client0 "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewComputeClient mocks base method.
func (m *MockFactory) NewComputeClient(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.ComputeClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewComputeClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.ComputeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewComputeClient indicates an expected call of NewComputeClient.
func (mr *MockFactoryMockRecorder) NewComputeClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewComputeClient", reflect.TypeOf((*MockFactory)(nil).NewComputeClient), arg0, arg1, arg2)
}

// NewDNSClient mocks base method.
func (m *MockFactory) NewDNSClient(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.DNSClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDNSClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.DNSClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDNSClient indicates an expected call of NewDNSClient.
func (mr *MockFactoryMockRecorder) NewDNSClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDNSClient", reflect.TypeOf((*MockFactory)(nil).NewDNSClient), arg0, arg1, arg2)
}

// NewIAMClient mocks base method.
func (m *MockFactory) NewIAMClient(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.IAMClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewIAMClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.IAMClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewIAMClient indicates an expected call of NewIAMClient.
func (mr *MockFactoryMockRecorder) NewIAMClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIAMClient", reflect.TypeOf((*MockFactory)(nil).NewIAMClient), arg0, arg1, arg2)
}

// NewStorageClient mocks base method.
func (m *MockFactory) NewStorageClient(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.StorageClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStorageClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.StorageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStorageClient indicates an expected call of NewStorageClient.
func (mr *MockFactoryMockRecorder) NewStorageClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStorageClient", reflect.TypeOf((*MockFactory)(nil).NewStorageClient), arg0, arg1, arg2)
}

// MockDNSClient is a mock of DNSClient interface.
type MockDNSClient struct {
	ctrl     *gomock.Controller
	recorder *MockDNSClientMockRecorder
}

// MockDNSClientMockRecorder is the mock recorder for MockDNSClient.
type MockDNSClientMockRecorder struct {
	mock *MockDNSClient
}

// NewMockDNSClient creates a new mock instance.
func NewMockDNSClient(ctrl *gomock.Controller) *MockDNSClient {
	mock := &MockDNSClient{ctrl: ctrl}
	mock.recorder = &MockDNSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNSClient) EXPECT() *MockDNSClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateRecordSet mocks base method.
func (m *MockDNSClient) CreateOrUpdateRecordSet(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string, arg5 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateRecordSet", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateRecordSet indicates an expected call of CreateOrUpdateRecordSet.
func (mr *MockDNSClientMockRecorder) CreateOrUpdateRecordSet(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateRecordSet", reflect.TypeOf((*MockDNSClient)(nil).CreateOrUpdateRecordSet), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DeleteRecordSet mocks base method.
func (m *MockDNSClient) DeleteRecordSet(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecordSet indicates an expected call of DeleteRecordSet.
func (mr *MockDNSClientMockRecorder) DeleteRecordSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordSet", reflect.TypeOf((*MockDNSClient)(nil).DeleteRecordSet), arg0, arg1, arg2, arg3)
}

// GetManagedZones mocks base method.
func (m *MockDNSClient) GetManagedZones(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedZones", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedZones indicates an expected call of GetManagedZones.
func (mr *MockDNSClientMockRecorder) GetManagedZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedZones", reflect.TypeOf((*MockDNSClient)(nil).GetManagedZones), arg0)
}

// MockComputeClient is a mock of ComputeClient interface.
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient.
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance.
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// GetExternalAddresses mocks base method.
func (m *MockComputeClient) GetExternalAddresses(arg0 context.Context, arg1 string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalAddresses", arg0, arg1)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalAddresses indicates an expected call of GetExternalAddresses.
func (mr *MockComputeClientMockRecorder) GetExternalAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalAddresses", reflect.TypeOf((*MockComputeClient)(nil).GetExternalAddresses), arg0, arg1)
}

// GetSubnet mocks base method.
func (m *MockComputeClient) GetSubnet(arg0 context.Context, arg1, arg2 string) (*compute.Subnetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*compute.Subnetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockComputeClientMockRecorder) GetSubnet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockComputeClient)(nil).GetSubnet), arg0, arg1, arg2)
}

// GetVPC mocks base method.
func (m *MockComputeClient) GetVPC(arg0 context.Context, arg1 string) (*compute.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVPC", arg0, arg1)
	ret0, _ := ret[0].(*compute.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVPC indicates an expected call of GetVPC.
func (mr *MockComputeClientMockRecorder) GetVPC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVPC", reflect.TypeOf((*MockComputeClient)(nil).GetVPC), arg0, arg1)
}
