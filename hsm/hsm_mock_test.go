// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

//go:build hsm
// +build hsm

// Code generated by MockGen. DO NOT EDIT.
// Source: hsm/hsm.go

// Package hsm_test is a generated GoMock package.
package hsm_test

import (
	elliptic "crypto/elliptic"
	reflect "reflect"

	crypto11 "github.com/ThalesIgnite/crypto11"
	gomock "github.com/golang/mock/gomock"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// FindKeyPair mocks base method.
func (m *MockContext) FindKeyPair(id, label []byte) (crypto11.Signer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKeyPair", id, label)
	ret0, _ := ret[0].(crypto11.Signer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKeyPair indicates an expected call of FindKeyPair.
func (mr *MockContextMockRecorder) FindKeyPair(id, label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKeyPair", reflect.TypeOf((*MockContext)(nil).FindKeyPair), id, label)
}

// FindKeyPairs mocks base method.
func (m *MockContext) FindKeyPairs(id, label []byte) ([]crypto11.Signer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKeyPairs", id, label)
	ret0, _ := ret[0].([]crypto11.Signer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindKeyPairs indicates an expected call of FindKeyPairs.
func (mr *MockContextMockRecorder) FindKeyPairs(id, label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKeyPairs", reflect.TypeOf((*MockContext)(nil).FindKeyPairs), id, label)
}

// GenerateECDSAKeyPairWithAttributes mocks base method.
func (m *MockContext) GenerateECDSAKeyPairWithAttributes(public, private crypto11.AttributeSet, curve elliptic.Curve) (crypto11.Signer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateECDSAKeyPairWithAttributes", public, private, curve)
	ret0, _ := ret[0].(crypto11.Signer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateECDSAKeyPairWithAttributes indicates an expected call of GenerateECDSAKeyPairWithAttributes.
func (mr *MockContextMockRecorder) GenerateECDSAKeyPairWithAttributes(public, private, curve interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateECDSAKeyPairWithAttributes", reflect.TypeOf((*MockContext)(nil).GenerateECDSAKeyPairWithAttributes), public, private, curve)
}

// GenerateRSAKeyPairWithAttributes mocks base method.
func (m *MockContext) GenerateRSAKeyPairWithAttributes(public, private crypto11.AttributeSet, bits int) (crypto11.SignerDecrypter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRSAKeyPairWithAttributes", public, private, bits)
	ret0, _ := ret[0].(crypto11.SignerDecrypter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRSAKeyPairWithAttributes indicates an expected call of GenerateRSAKeyPairWithAttributes.
func (mr *MockContextMockRecorder) GenerateRSAKeyPairWithAttributes(public, private, bits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRSAKeyPairWithAttributes", reflect.TypeOf((*MockContext)(nil).GenerateRSAKeyPairWithAttributes), public, private, bits)
}

// GetAttribute mocks base method.
func (m *MockContext) GetAttribute(key interface{}, attribute crypto11.AttributeType) (*crypto11.Attribute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttribute", key, attribute)
	ret0, _ := ret[0].(*crypto11.Attribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttribute indicates an expected call of GetAttribute.
func (mr *MockContextMockRecorder) GetAttribute(key, attribute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttribute", reflect.TypeOf((*MockContext)(nil).GetAttribute), key, attribute)
}
