// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/umee-network/peggo/orchestrator/ethereum/provider (interfaces: EVMProviderWithRet)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEVMProviderWithRet is a mock of EVMProviderWithRet interface.
type MockEVMProviderWithRet struct {
	ctrl     *gomock.Controller
	recorder *MockEVMProviderWithRetMockRecorder
}

// MockEVMProviderWithRetMockRecorder is the mock recorder for MockEVMProviderWithRet.
type MockEVMProviderWithRetMockRecorder struct {
	mock *MockEVMProviderWithRet
}

// NewMockEVMProviderWithRet creates a new mock instance.
func NewMockEVMProviderWithRet(ctrl *gomock.Controller) *MockEVMProviderWithRet {
	mock := &MockEVMProviderWithRet{ctrl: ctrl}
	mock.recorder = &MockEVMProviderWithRetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEVMProviderWithRet) EXPECT() *MockEVMProviderWithRetMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockEVMProviderWithRet) CallContract(arg0 context.Context, arg1 ethereum.CallMsg, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockEVMProviderWithRetMockRecorder) CallContract(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockEVMProviderWithRet)(nil).CallContract), arg0, arg1, arg2)
}

// CodeAt mocks base method.
func (m *MockEVMProviderWithRet) CodeAt(arg0 context.Context, arg1 common.Address, arg2 *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockEVMProviderWithRetMockRecorder) CodeAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockEVMProviderWithRet)(nil).CodeAt), arg0, arg1, arg2)
}

// EstimateGas mocks base method.
func (m *MockEVMProviderWithRet) EstimateGas(arg0 context.Context, arg1 ethereum.CallMsg) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas.
func (mr *MockEVMProviderWithRetMockRecorder) EstimateGas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockEVMProviderWithRet)(nil).EstimateGas), arg0, arg1)
}

// FilterLogs mocks base method.
func (m *MockEVMProviderWithRet) FilterLogs(arg0 context.Context, arg1 ethereum.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterLogs", arg0, arg1)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterLogs indicates an expected call of FilterLogs.
func (mr *MockEVMProviderWithRetMockRecorder) FilterLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterLogs", reflect.TypeOf((*MockEVMProviderWithRet)(nil).FilterLogs), arg0, arg1)
}

// HeaderByNumber mocks base method.
func (m *MockEVMProviderWithRet) HeaderByNumber(arg0 context.Context, arg1 *big.Int) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockEVMProviderWithRetMockRecorder) HeaderByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockEVMProviderWithRet)(nil).HeaderByNumber), arg0, arg1)
}

// PendingCodeAt mocks base method.
func (m *MockEVMProviderWithRet) PendingCodeAt(arg0 context.Context, arg1 common.Address) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingCodeAt", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingCodeAt indicates an expected call of PendingCodeAt.
func (mr *MockEVMProviderWithRetMockRecorder) PendingCodeAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingCodeAt", reflect.TypeOf((*MockEVMProviderWithRet)(nil).PendingCodeAt), arg0, arg1)
}

// PendingNonceAt mocks base method.
func (m *MockEVMProviderWithRet) PendingNonceAt(arg0 context.Context, arg1 common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PendingNonceAt", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PendingNonceAt indicates an expected call of PendingNonceAt.
func (mr *MockEVMProviderWithRetMockRecorder) PendingNonceAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingNonceAt", reflect.TypeOf((*MockEVMProviderWithRet)(nil).PendingNonceAt), arg0, arg1)
}

// SendTransaction mocks base method.
func (m *MockEVMProviderWithRet) SendTransaction(arg0 context.Context, arg1 *types.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockEVMProviderWithRetMockRecorder) SendTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockEVMProviderWithRet)(nil).SendTransaction), arg0, arg1)
}

// SendTransactionWithRet mocks base method.
func (m *MockEVMProviderWithRet) SendTransactionWithRet(arg0 context.Context, arg1 *types.Transaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransactionWithRet", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransactionWithRet indicates an expected call of SendTransactionWithRet.
func (mr *MockEVMProviderWithRetMockRecorder) SendTransactionWithRet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransactionWithRet", reflect.TypeOf((*MockEVMProviderWithRet)(nil).SendTransactionWithRet), arg0, arg1)
}

// SubscribeFilterLogs mocks base method.
func (m *MockEVMProviderWithRet) SubscribeFilterLogs(arg0 context.Context, arg1 ethereum.FilterQuery, arg2 chan<- types.Log) (ethereum.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeFilterLogs", arg0, arg1, arg2)
	ret0, _ := ret[0].(ethereum.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeFilterLogs indicates an expected call of SubscribeFilterLogs.
func (mr *MockEVMProviderWithRetMockRecorder) SubscribeFilterLogs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeFilterLogs", reflect.TypeOf((*MockEVMProviderWithRet)(nil).SubscribeFilterLogs), arg0, arg1, arg2)
}

// SuggestGasPrice mocks base method.
func (m *MockEVMProviderWithRet) SuggestGasPrice(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasPrice", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasPrice indicates an expected call of SuggestGasPrice.
func (mr *MockEVMProviderWithRetMockRecorder) SuggestGasPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasPrice", reflect.TypeOf((*MockEVMProviderWithRet)(nil).SuggestGasPrice), arg0)
}

// SuggestGasTipCap mocks base method.
func (m *MockEVMProviderWithRet) SuggestGasTipCap(arg0 context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuggestGasTipCap", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SuggestGasTipCap indicates an expected call of SuggestGasTipCap.
func (mr *MockEVMProviderWithRetMockRecorder) SuggestGasTipCap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuggestGasTipCap", reflect.TypeOf((*MockEVMProviderWithRet)(nil).SuggestGasTipCap), arg0)
}

// TransactionByHash mocks base method.
func (m *MockEVMProviderWithRet) TransactionByHash(arg0 context.Context, arg1 common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", arg0, arg1)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockEVMProviderWithRetMockRecorder) TransactionByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockEVMProviderWithRet)(nil).TransactionByHash), arg0, arg1)
}

// TransactionReceipt mocks base method.
func (m *MockEVMProviderWithRet) TransactionReceipt(arg0 context.Context, arg1 common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockEVMProviderWithRetMockRecorder) TransactionReceipt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockEVMProviderWithRet)(nil).TransactionReceipt), arg0, arg1)
}