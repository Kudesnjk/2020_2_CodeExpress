// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_CodeExpress/internal/album (interfaces: AlbumUsecase,AlbumRep)

// Package mock_album is a generated GoMock package.
package mock_album

import (
	models "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/models"
	error_response "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/error_response"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAlbumUsecase is a mock of AlbumUsecase interface
type MockAlbumUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumUsecaseMockRecorder
}

// MockAlbumUsecaseMockRecorder is the mock recorder for MockAlbumUsecase
type MockAlbumUsecaseMockRecorder struct {
	mock *MockAlbumUsecase
}

// NewMockAlbumUsecase creates a new mock instance
func NewMockAlbumUsecase(ctrl *gomock.Controller) *MockAlbumUsecase {
	mock := &MockAlbumUsecase{ctrl: ctrl}
	mock.recorder = &MockAlbumUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlbumUsecase) EXPECT() *MockAlbumUsecaseMockRecorder {
	return m.recorder
}

// CreateAlbum mocks base method
func (m *MockAlbumUsecase) CreateAlbum(arg0 *models.Album) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlbum", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// CreateAlbum indicates an expected call of CreateAlbum
func (mr *MockAlbumUsecaseMockRecorder) CreateAlbum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlbum", reflect.TypeOf((*MockAlbumUsecase)(nil).CreateAlbum), arg0)
}

// DeleteAlbum mocks base method
func (m *MockAlbumUsecase) DeleteAlbum(arg0 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlbum", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// DeleteAlbum indicates an expected call of DeleteAlbum
func (mr *MockAlbumUsecaseMockRecorder) DeleteAlbum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlbum", reflect.TypeOf((*MockAlbumUsecase)(nil).DeleteAlbum), arg0)
}

// GetByArtistID mocks base method
func (m *MockAlbumUsecase) GetByArtistID(arg0 uint64) ([]*models.Album, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByArtistID", arg0)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByArtistID indicates an expected call of GetByArtistID
func (mr *MockAlbumUsecaseMockRecorder) GetByArtistID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByArtistID", reflect.TypeOf((*MockAlbumUsecase)(nil).GetByArtistID), arg0)
}

// GetByID mocks base method
func (m *MockAlbumUsecase) GetByID(arg0 uint64) (*models.Album, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.Album)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockAlbumUsecaseMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAlbumUsecase)(nil).GetByID), arg0)
}

// GetByParams mocks base method
func (m *MockAlbumUsecase) GetByParams(arg0, arg1 uint64) ([]*models.Album, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParams", arg0, arg1)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByParams indicates an expected call of GetByParams
func (mr *MockAlbumUsecaseMockRecorder) GetByParams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParams", reflect.TypeOf((*MockAlbumUsecase)(nil).GetByParams), arg0, arg1)
}

// GetTopByParams mocks base method
func (m *MockAlbumUsecase) GetTopByParams(arg0, arg1 uint64) ([]*models.Album, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopByParams", arg0, arg1)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetTopByParams indicates an expected call of GetTopByParams
func (mr *MockAlbumUsecaseMockRecorder) GetTopByParams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopByParams", reflect.TypeOf((*MockAlbumUsecase)(nil).GetTopByParams), arg0, arg1)
}

// UpdateAlbum mocks base method
func (m *MockAlbumUsecase) UpdateAlbum(arg0 *models.Album) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlbum", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// UpdateAlbum indicates an expected call of UpdateAlbum
func (mr *MockAlbumUsecaseMockRecorder) UpdateAlbum(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlbum", reflect.TypeOf((*MockAlbumUsecase)(nil).UpdateAlbum), arg0)
}

// UpdateAlbumPoster mocks base method
func (m *MockAlbumUsecase) UpdateAlbumPoster(arg0 *models.Album) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlbumPoster", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// UpdateAlbumPoster indicates an expected call of UpdateAlbumPoster
func (mr *MockAlbumUsecaseMockRecorder) UpdateAlbumPoster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlbumPoster", reflect.TypeOf((*MockAlbumUsecase)(nil).UpdateAlbumPoster), arg0)
}

// MockAlbumRep is a mock of AlbumRep interface
type MockAlbumRep struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumRepMockRecorder
}

// MockAlbumRepMockRecorder is the mock recorder for MockAlbumRep
type MockAlbumRepMockRecorder struct {
	mock *MockAlbumRep
}

// NewMockAlbumRep creates a new mock instance
func NewMockAlbumRep(ctrl *gomock.Controller) *MockAlbumRep {
	mock := &MockAlbumRep{ctrl: ctrl}
	mock.recorder = &MockAlbumRepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlbumRep) EXPECT() *MockAlbumRepMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockAlbumRep) Delete(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAlbumRepMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAlbumRep)(nil).Delete), arg0)
}

// Insert mocks base method
func (m *MockAlbumRep) Insert(arg0 *models.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockAlbumRepMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAlbumRep)(nil).Insert), arg0)
}

// SelectByArtistID mocks base method
func (m *MockAlbumRep) SelectByArtistID(arg0 uint64) ([]*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByArtistID", arg0)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByArtistID indicates an expected call of SelectByArtistID
func (mr *MockAlbumRepMockRecorder) SelectByArtistID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByArtistID", reflect.TypeOf((*MockAlbumRep)(nil).SelectByArtistID), arg0)
}

// SelectByID mocks base method
func (m *MockAlbumRep) SelectByID(arg0 uint64) (*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByID", arg0)
	ret0, _ := ret[0].(*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockAlbumRepMockRecorder) SelectByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockAlbumRep)(nil).SelectByID), arg0)
}

// SelectByParam mocks base method
func (m *MockAlbumRep) SelectByParam(arg0, arg1 uint64) ([]*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByParam", arg0, arg1)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByParam indicates an expected call of SelectByParam
func (mr *MockAlbumRepMockRecorder) SelectByParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByParam", reflect.TypeOf((*MockAlbumRep)(nil).SelectByParam), arg0, arg1)
}

// SelectTopByParam mocks base method
func (m *MockAlbumRep) SelectTopByParam(arg0, arg1 uint64) ([]*models.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTopByParam", arg0, arg1)
	ret0, _ := ret[0].([]*models.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTopByParam indicates an expected call of SelectTopByParam
func (mr *MockAlbumRepMockRecorder) SelectTopByParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTopByParam", reflect.TypeOf((*MockAlbumRep)(nil).SelectTopByParam), arg0, arg1)
}

// Update mocks base method
func (m *MockAlbumRep) Update(arg0 *models.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockAlbumRepMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAlbumRep)(nil).Update), arg0)
}

// UpdatePoster mocks base method
func (m *MockAlbumRep) UpdatePoster(arg0 *models.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePoster", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePoster indicates an expected call of UpdatePoster
func (mr *MockAlbumRepMockRecorder) UpdatePoster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePoster", reflect.TypeOf((*MockAlbumRep)(nil).UpdatePoster), arg0)
}
