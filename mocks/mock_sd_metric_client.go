// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/ts-bridge/stackdriver (interfaces: MetricClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	monitoring "cloud.google.com/go/monitoring/apiv3"
	context "context"
	gomock "github.com/golang/mock/gomock"
	gax "github.com/googleapis/gax-go/v2"
	metric "google.golang.org/genproto/googleapis/api/metric"
	monitoring0 "google.golang.org/genproto/googleapis/monitoring/v3"
	reflect "reflect"
)

// MockMetricClient is a mock of MetricClient interface
type MockMetricClient struct {
	ctrl     *gomock.Controller
	recorder *MockMetricClientMockRecorder
}

// MockMetricClientMockRecorder is the mock recorder for MockMetricClient
type MockMetricClientMockRecorder struct {
	mock *MockMetricClient
}

// NewMockMetricClient creates a new mock instance
func NewMockMetricClient(ctrl *gomock.Controller) *MockMetricClient {
	mock := &MockMetricClient{ctrl: ctrl}
	mock.recorder = &MockMetricClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetricClient) EXPECT() *MockMetricClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMetricClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMetricClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetricClient)(nil).Close))
}

// CreateMetricDescriptor mocks base method
func (m *MockMetricClient) CreateMetricDescriptor(arg0 context.Context, arg1 *monitoring0.CreateMetricDescriptorRequest, arg2 ...gax.CallOption) (*metric.MetricDescriptor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMetricDescriptor", varargs...)
	ret0, _ := ret[0].(*metric.MetricDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMetricDescriptor indicates an expected call of CreateMetricDescriptor
func (mr *MockMetricClientMockRecorder) CreateMetricDescriptor(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMetricDescriptor", reflect.TypeOf((*MockMetricClient)(nil).CreateMetricDescriptor), varargs...)
}

// CreateTimeSeries mocks base method
func (m *MockMetricClient) CreateTimeSeries(arg0 context.Context, arg1 *monitoring0.CreateTimeSeriesRequest, arg2 ...gax.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTimeSeries", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTimeSeries indicates an expected call of CreateTimeSeries
func (mr *MockMetricClientMockRecorder) CreateTimeSeries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTimeSeries", reflect.TypeOf((*MockMetricClient)(nil).CreateTimeSeries), varargs...)
}

// DeleteMetricDescriptor mocks base method
func (m *MockMetricClient) DeleteMetricDescriptor(arg0 context.Context, arg1 *monitoring0.DeleteMetricDescriptorRequest, arg2 ...gax.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMetricDescriptor", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMetricDescriptor indicates an expected call of DeleteMetricDescriptor
func (mr *MockMetricClientMockRecorder) DeleteMetricDescriptor(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMetricDescriptor", reflect.TypeOf((*MockMetricClient)(nil).DeleteMetricDescriptor), varargs...)
}

// GetMetricDescriptor mocks base method
func (m *MockMetricClient) GetMetricDescriptor(arg0 context.Context, arg1 *monitoring0.GetMetricDescriptorRequest, arg2 ...gax.CallOption) (*metric.MetricDescriptor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetricDescriptor", varargs...)
	ret0, _ := ret[0].(*metric.MetricDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetricDescriptor indicates an expected call of GetMetricDescriptor
func (mr *MockMetricClientMockRecorder) GetMetricDescriptor(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricDescriptor", reflect.TypeOf((*MockMetricClient)(nil).GetMetricDescriptor), varargs...)
}

// ListTimeSeries mocks base method
func (m *MockMetricClient) ListTimeSeries(arg0 context.Context, arg1 *monitoring0.ListTimeSeriesRequest, arg2 ...gax.CallOption) *monitoring.TimeSeriesIterator {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTimeSeries", varargs...)
	ret0, _ := ret[0].(*monitoring.TimeSeriesIterator)
	return ret0
}

// ListTimeSeries indicates an expected call of ListTimeSeries
func (mr *MockMetricClientMockRecorder) ListTimeSeries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTimeSeries", reflect.TypeOf((*MockMetricClient)(nil).ListTimeSeries), varargs...)
}
