// Code generated by MockGen. DO NOT EDIT.
// Source: bank_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "paymentplatform/pkg/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBankClient is a mock of BankClient interface.
type MockBankClient struct {
	ctrl     *gomock.Controller
	recorder *MockBankClientMockRecorder
}

// MockBankClientMockRecorder is the mock recorder for MockBankClient.
type MockBankClientMockRecorder struct {
	mock *MockBankClient
}

// NewMockBankClient creates a new mock instance.
func NewMockBankClient(ctrl *gomock.Controller) *MockBankClient {
	mock := &MockBankClient{ctrl: ctrl}
	mock.recorder = &MockBankClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankClient) EXPECT() *MockBankClientMockRecorder {
	return m.recorder
}

// ProcessPayment mocks base method.
func (m *MockBankClient) ProcessPayment(ctx context.Context, paymentDetails entity.PaymentDetails) (entity.PaymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPayment", ctx, paymentDetails)
	ret0, _ := ret[0].(entity.PaymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessPayment indicates an expected call of ProcessPayment.
func (mr *MockBankClientMockRecorder) ProcessPayment(ctx, paymentDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPayment", reflect.TypeOf((*MockBankClient)(nil).ProcessPayment), ctx, paymentDetails)
}

// ProcessRefund mocks base method.
func (m *MockBankClient) ProcessRefund(ctx context.Context, refundDetails entity.RefundDetails) (entity.RefundResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessRefund", ctx, refundDetails)
	ret0, _ := ret[0].(entity.RefundResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessRefund indicates an expected call of ProcessRefund.
func (mr *MockBankClientMockRecorder) ProcessRefund(ctx, refundDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessRefund", reflect.TypeOf((*MockBankClient)(nil).ProcessRefund), ctx, refundDetails)
}
