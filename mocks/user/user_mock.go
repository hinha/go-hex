package user

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testHEX/internal/constants/model"
)

// MockPersistence is a mock of Persistence interface
type MockPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceMockRecorder
}

// MockPersistenceMockRecorder is the mock recorder for MockPersistence
type MockPersistenceMockRecorder struct {
	mock *MockPersistence
}

// NewMockPersistence creates a new mock instance
func NewMockPersistence(ctrl *gomock.Controller) *MockPersistence {
	mock := &MockPersistence{ctrl: ctrl}
	mock.recorder = &MockPersistenceMockRecorder{mock}
	return mock
}

func (m *MockPersistence) EXPECT() *MockPersistenceMockRecorder {
	return m.recorder
}

func (m *MockPersistence) Create(user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPersistenceMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPersistence)(nil).Create), user)
}

// Find mocks base method
func (m *MockPersistence) Find(email, password string) (*model.User, *model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", email, password)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(*model.Token)
	ret2, _ := ret[2].(error)

	return ret0, ret1, ret2
}

// Find indicates an expected call of Find
func (mr *MockPersistenceMockRecorder) Find(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPersistence)(nil).Find), email, password)
}

func (m *MockPersistence) FindByEmail(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockPersistenceMockRecorder) FindByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockPersistence)(nil).FindByEmail), email)
}

type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

func (m *MockUsecase) Login(email, password string) (string, string, error) {
	panic("implement me")
}

func (m *MockPersistence) FindByID(userID int64) (*model.User, error) {
	panic("implement me")
}


type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockUsecase) Register(user *model.User) error {
	m.ctrl.T.Helper()
	user.Password = ""
	ret := m.ctrl.Call(m, "Register", user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUsecaseMockRecorder) Register(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUsecase)(nil).Register), user)
}

// MockCaching is a mock of Caching interface
type MockCaching struct {
	ctrl     *gomock.Controller
	recorder *MockCachingMockRecorder
}

// MockCachingMockRecorder is the mock recorder for MockCaching
type MockCachingMockRecorder struct {
	mock *MockCaching
}

func NewMockCaching(ctrl *gomock.Controller) *MockCaching {
	mock := &MockCaching{ctrl: ctrl}
	mock.recorder = &MockCachingMockRecorder{mock}
	return mock
}

func (m *MockCaching) EXPECT() *MockCachingMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockCaching) Save(user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockCachingMockRecorder) Save(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCaching)(nil).Save), user)
}

// Get mocks base method
func (m *MockCaching) Get(userID string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCachingMockRecorder) Get(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCaching)(nil).Get), userID)
}

// Delete mocks base method
func (m *MockCaching) Delete(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCachingMockRecorder) Delete(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCaching)(nil).Delete), userID)
}


// SaveToken mocks base method
func (m *MockCaching) SaveToken(token *model.Token, user *model.User) error {
	m.ctrl.T.Helper()

	token.UniqueToken = "TGrGeaK5iXSFa_GSyeBtIVdpHBHiofg3lRTmUSVgK-alepTHC-25aFInYg=="
	token.TimeAt = "1594374962"

	ret := m.ctrl.Call(m, "SaveToken", token, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveToken indicates an expected call of Save
func (mr *MockCachingMockRecorder) SaveToken(token, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveToken", reflect.TypeOf((*MockCaching)(nil).SaveToken), token, user)
}

// Get mocks base method
func (m *MockCaching) GetToken(userID string) (*model.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", userID)
	ret0, _ := ret[0].(*model.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCachingMockRecorder) GetToken(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockCaching)(nil).GetToken), userID)
}

// Delete mocks base method
func (m *MockCaching) DeleteToken(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToken", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCachingMockRecorder) DeleteToken(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockCaching)(nil).DeleteToken), userID)
}