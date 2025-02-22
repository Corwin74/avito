// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/transaction/interface.go

// Package transaction is a generated GoMock package.
package transaction

import (
	context "context"
	reflect "reflect"
	transaction "shop/pkg/transaction"

	gomock "github.com/golang/mock/gomock"
)

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTransaction) Commit(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionMockRecorder) Commit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit), ctx)
}

// Rollback mocks base method.
func (m *MockTransaction) Rollback(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionMockRecorder) Rollback(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransaction)(nil).Rollback), ctx)
}

// MockFabric is a mock of Fabric interface.
type MockFabric struct {
	ctrl     *gomock.Controller
	recorder *MockFabricMockRecorder
}

// MockFabricMockRecorder is the mock recorder for MockFabric.
type MockFabricMockRecorder struct {
	mock *MockFabric
}

// NewMockFabric creates a new mock instance.
func NewMockFabric(ctrl *gomock.Controller) *MockFabric {
	mock := &MockFabric{ctrl: ctrl}
	mock.recorder = &MockFabricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFabric) EXPECT() *MockFabricMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockFabric) Begin(ctx context.Context) (context.Context, transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", ctx)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(transaction.Transaction)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Begin indicates an expected call of Begin.
func (mr *MockFabricMockRecorder) Begin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockFabric)(nil).Begin), ctx)
}
