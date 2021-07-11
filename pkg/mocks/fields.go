// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cheebo/go-config (interfaces: Fields)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	go_config "github.com/cheebo/go-config"
	gomock "github.com/golang/mock/gomock"
)

// MockFields is a mock of Fields interface.
type MockFields struct {
	ctrl     *gomock.Controller
	recorder *MockFieldsMockRecorder
}

// MockFieldsMockRecorder is the mock recorder for MockFields.
type MockFieldsMockRecorder struct {
	mock *MockFields
}

// NewMockFields creates a new mock instance.
func NewMockFields(ctrl *gomock.Controller) *MockFields {
	mock := &MockFields{ctrl: ctrl}
	mock.recorder = &MockFieldsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFields) EXPECT() *MockFieldsMockRecorder {
	return m.recorder
}

// Bool mocks base method.
func (m *MockFields) Bool(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bool", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Bool indicates an expected call of Bool.
func (mr *MockFieldsMockRecorder) Bool(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockFields)(nil).Bool), arg0)
}

// Float mocks base method.
func (m *MockFields) Float(arg0 string) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Float indicates an expected call of Float.
func (mr *MockFieldsMockRecorder) Float(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockFields)(nil).Float), arg0)
}

// Get mocks base method.
func (m *MockFields) Get(arg0 string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockFieldsMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFields)(nil).Get), arg0)
}

// Int mocks base method.
func (m *MockFields) Int(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Int indicates an expected call of Int.
func (mr *MockFieldsMockRecorder) Int(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockFields)(nil).Int), arg0)
}

// Int32 mocks base method.
func (m *MockFields) Int32(arg0 string) int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int32", arg0)
	ret0, _ := ret[0].(int32)
	return ret0
}

// Int32 indicates an expected call of Int32.
func (mr *MockFieldsMockRecorder) Int32(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int32", reflect.TypeOf((*MockFields)(nil).Int32), arg0)
}

// Int64 mocks base method.
func (m *MockFields) Int64(arg0 string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// Int64 indicates an expected call of Int64.
func (mr *MockFieldsMockRecorder) Int64(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64", reflect.TypeOf((*MockFields)(nil).Int64), arg0)
}

// Int8 mocks base method.
func (m *MockFields) Int8(arg0 string) int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int8", arg0)
	ret0, _ := ret[0].(int8)
	return ret0
}

// Int8 indicates an expected call of Int8.
func (mr *MockFieldsMockRecorder) Int8(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int8", reflect.TypeOf((*MockFields)(nil).Int8), arg0)
}

// IsSet mocks base method.
func (m *MockFields) IsSet(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSet", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSet indicates an expected call of IsSet.
func (mr *MockFieldsMockRecorder) IsSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSet", reflect.TypeOf((*MockFields)(nil).IsSet), arg0)
}

// Slice mocks base method.
func (m *MockFields) Slice(arg0 string) []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slice", arg0)
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// Slice indicates an expected call of Slice.
func (mr *MockFieldsMockRecorder) Slice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slice", reflect.TypeOf((*MockFields)(nil).Slice), arg0)
}

// SliceInt mocks base method.
func (m *MockFields) SliceInt(arg0 string) []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SliceInt", arg0)
	ret0, _ := ret[0].([]int)
	return ret0
}

// SliceInt indicates an expected call of SliceInt.
func (mr *MockFieldsMockRecorder) SliceInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SliceInt", reflect.TypeOf((*MockFields)(nil).SliceInt), arg0)
}

// SliceString mocks base method.
func (m *MockFields) SliceString(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SliceString", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// SliceString indicates an expected call of SliceString.
func (mr *MockFieldsMockRecorder) SliceString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SliceString", reflect.TypeOf((*MockFields)(nil).SliceString), arg0)
}

// String mocks base method.
func (m *MockFields) String(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockFieldsMockRecorder) String(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockFields)(nil).String), arg0)
}

// StringMap mocks base method.
func (m *MockFields) StringMap(arg0 string) map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringMap", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// StringMap indicates an expected call of StringMap.
func (mr *MockFieldsMockRecorder) StringMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringMap", reflect.TypeOf((*MockFields)(nil).StringMap), arg0)
}

// StringMapInt mocks base method.
func (m *MockFields) StringMapInt(arg0 string) map[string]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringMapInt", arg0)
	ret0, _ := ret[0].(map[string]int)
	return ret0
}

// StringMapInt indicates an expected call of StringMapInt.
func (mr *MockFieldsMockRecorder) StringMapInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringMapInt", reflect.TypeOf((*MockFields)(nil).StringMapInt), arg0)
}

// StringMapSliceString mocks base method.
func (m *MockFields) StringMapSliceString(arg0 string) map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringMapSliceString", arg0)
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// StringMapSliceString indicates an expected call of StringMapSliceString.
func (mr *MockFieldsMockRecorder) StringMapSliceString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringMapSliceString", reflect.TypeOf((*MockFields)(nil).StringMapSliceString), arg0)
}

// StringMapString mocks base method.
func (m *MockFields) StringMapString(arg0 string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringMapString", arg0)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// StringMapString indicates an expected call of StringMapString.
func (mr *MockFieldsMockRecorder) StringMapString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringMapString", reflect.TypeOf((*MockFields)(nil).StringMapString), arg0)
}

// Sub mocks base method.
func (m *MockFields) Sub(arg0 string) go_config.Fields {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sub", arg0)
	ret0, _ := ret[0].(go_config.Fields)
	return ret0
}

// Sub indicates an expected call of Sub.
func (mr *MockFieldsMockRecorder) Sub(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockFields)(nil).Sub), arg0)
}

// UInt mocks base method.
func (m *MockFields) UInt(arg0 string) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UInt", arg0)
	ret0, _ := ret[0].(uint)
	return ret0
}

// UInt indicates an expected call of UInt.
func (mr *MockFieldsMockRecorder) UInt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UInt", reflect.TypeOf((*MockFields)(nil).UInt), arg0)
}

// UInt32 mocks base method.
func (m *MockFields) UInt32(arg0 string) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UInt32", arg0)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// UInt32 indicates an expected call of UInt32.
func (mr *MockFieldsMockRecorder) UInt32(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UInt32", reflect.TypeOf((*MockFields)(nil).UInt32), arg0)
}

// UInt64 mocks base method.
func (m *MockFields) UInt64(arg0 string) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UInt64", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// UInt64 indicates an expected call of UInt64.
func (mr *MockFieldsMockRecorder) UInt64(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UInt64", reflect.TypeOf((*MockFields)(nil).UInt64), arg0)
}

// Unmarshal mocks base method.
func (m *MockFields) Unmarshal(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockFieldsMockRecorder) Unmarshal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockFields)(nil).Unmarshal), arg0, arg1)
}