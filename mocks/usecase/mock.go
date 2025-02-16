// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/user/interface.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"
	models "shop/internal/models"
	item "shop/internal/repository/item"
	transferhistory "shop/internal/repository/transferhistory"
	transferhistoryname "shop/internal/repository/transferhistoryname"
	user "shop/internal/repository/user"
	useritem "shop/internal/repository/useritem"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepo) Create(ctx context.Context, user models.User) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepoMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepo)(nil).Create), ctx, user)
}

// Get mocks base method.
func (m *MockUserRepo) Get(ctx context.Context, filter user.Filter, opts user.GetOptions) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter, opts)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserRepoMockRecorder) Get(ctx, filter, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRepo)(nil).Get), ctx, filter, opts)
}

// IsAuth mocks base method.
func (m *MockUserRepo) IsAuth(ctx context.Context) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuth", ctx)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuth indicates an expected call of IsAuth.
func (mr *MockUserRepoMockRecorder) IsAuth(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuth", reflect.TypeOf((*MockUserRepo)(nil).IsAuth), ctx)
}

// Update mocks base method.
func (m *MockUserRepo) Update(ctx context.Context, update user.Update, filter user.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, update, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepoMockRecorder) Update(ctx, update, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepo)(nil).Update), ctx, update, filter)
}

// MockItemRepo is a mock of ItemRepo interface.
type MockItemRepo struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepoMockRecorder
}

// MockItemRepoMockRecorder is the mock recorder for MockItemRepo.
type MockItemRepoMockRecorder struct {
	mock *MockItemRepo
}

// NewMockItemRepo creates a new mock instance.
func NewMockItemRepo(ctrl *gomock.Controller) *MockItemRepo {
	mock := &MockItemRepo{ctrl: ctrl}
	mock.recorder = &MockItemRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepo) EXPECT() *MockItemRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockItemRepo) Get(ctx context.Context, filter item.Filter) (models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter)
	ret0, _ := ret[0].(models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockItemRepoMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockItemRepo)(nil).Get), ctx, filter)
}

// GetMany mocks base method.
func (m *MockItemRepo) GetMany(ctx context.Context, filter item.Filter) ([]models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, filter)
	ret0, _ := ret[0].([]models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockItemRepoMockRecorder) GetMany(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockItemRepo)(nil).GetMany), ctx, filter)
}

// MockTransferHistory is a mock of TransferHistory interface.
type MockTransferHistory struct {
	ctrl     *gomock.Controller
	recorder *MockTransferHistoryMockRecorder
}

// MockTransferHistoryMockRecorder is the mock recorder for MockTransferHistory.
type MockTransferHistoryMockRecorder struct {
	mock *MockTransferHistory
}

// NewMockTransferHistory creates a new mock instance.
func NewMockTransferHistory(ctrl *gomock.Controller) *MockTransferHistory {
	mock := &MockTransferHistory{ctrl: ctrl}
	mock.recorder = &MockTransferHistoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferHistory) EXPECT() *MockTransferHistoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransferHistory) Create(ctx context.Context, th models.TransferHistory) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, th)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTransferHistoryMockRecorder) Create(ctx, th interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransferHistory)(nil).Create), ctx, th)
}

// GetMany mocks base method.
func (m *MockTransferHistory) GetMany(ctx context.Context, filter transferhistory.Filter) ([]models.TransferHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, filter)
	ret0, _ := ret[0].([]models.TransferHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockTransferHistoryMockRecorder) GetMany(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockTransferHistory)(nil).GetMany), ctx, filter)
}

// MockTransferHistoryName is a mock of TransferHistoryName interface.
type MockTransferHistoryName struct {
	ctrl     *gomock.Controller
	recorder *MockTransferHistoryNameMockRecorder
}

// MockTransferHistoryNameMockRecorder is the mock recorder for MockTransferHistoryName.
type MockTransferHistoryNameMockRecorder struct {
	mock *MockTransferHistoryName
}

// NewMockTransferHistoryName creates a new mock instance.
func NewMockTransferHistoryName(ctrl *gomock.Controller) *MockTransferHistoryName {
	mock := &MockTransferHistoryName{ctrl: ctrl}
	mock.recorder = &MockTransferHistoryNameMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferHistoryName) EXPECT() *MockTransferHistoryNameMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTransferHistoryName) Create(ctx context.Context, th models.TransferHistoryName) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, th)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTransferHistoryNameMockRecorder) Create(ctx, th interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTransferHistoryName)(nil).Create), ctx, th)
}

// GetMany mocks base method.
func (m *MockTransferHistoryName) GetMany(ctx context.Context, filter transferhistoryname.Filter) ([]models.TransferHistoryName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, filter)
	ret0, _ := ret[0].([]models.TransferHistoryName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockTransferHistoryNameMockRecorder) GetMany(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockTransferHistoryName)(nil).GetMany), ctx, filter)
}

// MockUserItemRepo is a mock of UserItemRepo interface.
type MockUserItemRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserItemRepoMockRecorder
}

// MockUserItemRepoMockRecorder is the mock recorder for MockUserItemRepo.
type MockUserItemRepoMockRecorder struct {
	mock *MockUserItemRepo
}

// NewMockUserItemRepo creates a new mock instance.
func NewMockUserItemRepo(ctrl *gomock.Controller) *MockUserItemRepo {
	mock := &MockUserItemRepo{ctrl: ctrl}
	mock.recorder = &MockUserItemRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserItemRepo) EXPECT() *MockUserItemRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserItemRepo) Create(ctx context.Context, md models.UserItem) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, md)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserItemRepoMockRecorder) Create(ctx, md interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserItemRepo)(nil).Create), ctx, md)
}

// Get mocks base method.
func (m *MockUserItemRepo) Get(ctx context.Context, filter useritem.Filter) (models.UserItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter)
	ret0, _ := ret[0].(models.UserItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserItemRepoMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserItemRepo)(nil).Get), ctx, filter)
}

// GetUserItemsAmount mocks base method.
func (m *MockUserItemRepo) GetUserItemsAmount(ctx context.Context, filter useritem.Filter) ([]models.UserItemsAmount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserItemsAmount", ctx, filter)
	ret0, _ := ret[0].([]models.UserItemsAmount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserItemsAmount indicates an expected call of GetUserItemsAmount.
func (mr *MockUserItemRepoMockRecorder) GetUserItemsAmount(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserItemsAmount", reflect.TypeOf((*MockUserItemRepo)(nil).GetUserItemsAmount), ctx, filter)
}
