// Code generated by MockGen. DO NOT EDIT.
// Source: movie/internal/controller/movie/controller.go

// Package repository is a generated GoMock package.
package repository

import (
        context "context"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
        pkg "movieexample.com/metadata/pkg/model"
        model "movieexample.com/rating/pkg/model"
)

// MockratingGateway is a mock of ratingGateway interface.
type MockratingGateway struct {
        ctrl     *gomock.Controller
        recorder *MockratingGatewayMockRecorder
}

// MockratingGatewayMockRecorder is the mock recorder for MockratingGateway.
type MockratingGatewayMockRecorder struct {
        mock *MockratingGateway
}

// NewMockratingGateway creates a new mock instance.
func NewMockratingGateway(ctrl *gomock.Controller) *MockratingGateway {
        mock := &MockratingGateway{ctrl: ctrl}
        mock.recorder = &MockratingGatewayMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockratingGateway) EXPECT() *MockratingGatewayMockRecorder {
        return m.recorder
}

// GetAggregatedRating mocks base method.
func (m *MockratingGateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAggregatedRating", ctx, recordID, recordType)
        ret0, _ := ret[0].(float64)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAggregatedRating indicates an expected call of GetAggregatedRating.
func (mr *MockratingGatewayMockRecorder) GetAggregatedRating(ctx, recordID, recordType interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAggregatedRating", reflect.TypeOf((*MockratingGateway)(nil).GetAggregatedRating), ctx, recordID, recordType)
}

// PutRating mocks base method.
func (m *MockratingGateway) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "PutRating", ctx, recordID, recordType, rating)
        ret0, _ := ret[0].(error)
        return ret0
}

// PutRating indicates an expected call of PutRating.
func (mr *MockratingGatewayMockRecorder) PutRating(ctx, recordID, recordType, rating interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRating", reflect.TypeOf((*MockratingGateway)(nil).PutRating), ctx, recordID, recordType, rating)
}

// MockMetaDataGateway is a mock of MetaDataGateway interface.
type MockMetaDataGateway struct {
        ctrl     *gomock.Controller
        recorder *MockMetaDataGatewayMockRecorder
}

// MockMetaDataGatewayMockRecorder is the mock recorder for MockMetaDataGateway.
type MockMetaDataGatewayMockRecorder struct {
        mock *MockMetaDataGateway
}

// NewMockMetaDataGateway creates a new mock instance.
func NewMockMetaDataGateway(ctrl *gomock.Controller) *MockMetaDataGateway {
        mock := &MockMetaDataGateway{ctrl: ctrl}
        mock.recorder = &MockMetaDataGatewayMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetaDataGateway) EXPECT() *MockMetaDataGatewayMockRecorder {
        return m.recorder
}

// Get mocks base method.
func (m *MockMetaDataGateway) Get(ctx context.Context, id string) (*pkg.MetaData, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Get", ctx, id)
        ret0, _ := ret[0].(*pkg.MetaData)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMetaDataGatewayMockRecorder) Get(ctx, id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMetaDataGateway)(nil).Get), ctx, id)
}