// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../x/interchaintxs/types/expected_keepers.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/capability/types"
	types2 "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	types3 "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	exported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
	types4 "github.com/neutron-org/neutron/x/contractmanager/types"
	types5 "github.com/neutron-org/neutron/x/feerefunder/types"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockContractManagerKeeper is a mock of ContractManagerKeeper interface.
type MockContractManagerKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockContractManagerKeeperMockRecorder
}

// MockContractManagerKeeperMockRecorder is the mock recorder for MockContractManagerKeeper.
type MockContractManagerKeeperMockRecorder struct {
	mock *MockContractManagerKeeper
}

// NewMockContractManagerKeeper creates a new mock instance.
func NewMockContractManagerKeeper(ctrl *gomock.Controller) *MockContractManagerKeeper {
	mock := &MockContractManagerKeeper{ctrl: ctrl}
	mock.recorder = &MockContractManagerKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractManagerKeeper) EXPECT() *MockContractManagerKeeperMockRecorder {
	return m.recorder
}

// AddContractFailure mocks base method.
func (m *MockContractManagerKeeper) AddContractFailure(ctx types.Context, channelId, address string, ackID uint64, ackType string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddContractFailure", ctx, channelId, address, ackID, ackType)
}

// AddContractFailure indicates an expected call of AddContractFailure.
func (mr *MockContractManagerKeeperMockRecorder) AddContractFailure(ctx, channelId, address, ackID, ackType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContractFailure", reflect.TypeOf((*MockContractManagerKeeper)(nil).AddContractFailure), ctx, channelId, address, ackID, ackType)
}

// HasContractInfo mocks base method.
func (m *MockContractManagerKeeper) HasContractInfo(ctx types.Context, contractAddress types.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasContractInfo", ctx, contractAddress)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasContractInfo indicates an expected call of HasContractInfo.
func (mr *MockContractManagerKeeperMockRecorder) HasContractInfo(ctx, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasContractInfo", reflect.TypeOf((*MockContractManagerKeeper)(nil).HasContractInfo), ctx, contractAddress)
}

// SudoError mocks base method.
func (m *MockContractManagerKeeper) SudoError(ctx types.Context, senderAddress types.AccAddress, request types3.Packet, details string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoError", ctx, senderAddress, request, details)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoError indicates an expected call of SudoError.
func (mr *MockContractManagerKeeperMockRecorder) SudoError(ctx, senderAddress, request, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoError", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoError), ctx, senderAddress, request, details)
}

// SudoOnChanOpenAck mocks base method.
func (m *MockContractManagerKeeper) SudoOnChanOpenAck(ctx types.Context, contractAddress types.AccAddress, details types4.OpenAckDetails) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoOnChanOpenAck", ctx, contractAddress, details)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoOnChanOpenAck indicates an expected call of SudoOnChanOpenAck.
func (mr *MockContractManagerKeeperMockRecorder) SudoOnChanOpenAck(ctx, contractAddress, details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoOnChanOpenAck", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoOnChanOpenAck), ctx, contractAddress, details)
}

// SudoResponse mocks base method.
func (m *MockContractManagerKeeper) SudoResponse(ctx types.Context, senderAddress types.AccAddress, request types3.Packet, msg []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoResponse", ctx, senderAddress, request, msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoResponse indicates an expected call of SudoResponse.
func (mr *MockContractManagerKeeperMockRecorder) SudoResponse(ctx, senderAddress, request, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoResponse", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoResponse), ctx, senderAddress, request, msg)
}

// SudoTimeout mocks base method.
func (m *MockContractManagerKeeper) SudoTimeout(ctx types.Context, senderAddress types.AccAddress, request types3.Packet) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SudoTimeout", ctx, senderAddress, request)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SudoTimeout indicates an expected call of SudoTimeout.
func (mr *MockContractManagerKeeperMockRecorder) SudoTimeout(ctx, senderAddress, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SudoTimeout", reflect.TypeOf((*MockContractManagerKeeper)(nil).SudoTimeout), ctx, senderAddress, request)
}

// MockICAControllerKeeper is a mock of ICAControllerKeeper interface.
type MockICAControllerKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockICAControllerKeeperMockRecorder
}

// MockICAControllerKeeperMockRecorder is the mock recorder for MockICAControllerKeeper.
type MockICAControllerKeeperMockRecorder struct {
	mock *MockICAControllerKeeper
}

// NewMockICAControllerKeeper creates a new mock instance.
func NewMockICAControllerKeeper(ctrl *gomock.Controller) *MockICAControllerKeeper {
	mock := &MockICAControllerKeeper{ctrl: ctrl}
	mock.recorder = &MockICAControllerKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICAControllerKeeper) EXPECT() *MockICAControllerKeeperMockRecorder {
	return m.recorder
}

// GetActiveChannelID mocks base method.
func (m *MockICAControllerKeeper) GetActiveChannelID(ctx types.Context, connectionID, portID string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveChannelID", ctx, connectionID, portID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetActiveChannelID indicates an expected call of GetActiveChannelID.
func (mr *MockICAControllerKeeperMockRecorder) GetActiveChannelID(ctx, connectionID, portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveChannelID", reflect.TypeOf((*MockICAControllerKeeper)(nil).GetActiveChannelID), ctx, connectionID, portID)
}

// GetInterchainAccountAddress mocks base method.
func (m *MockICAControllerKeeper) GetInterchainAccountAddress(ctx types.Context, connectionID, portID string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterchainAccountAddress", ctx, connectionID, portID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetInterchainAccountAddress indicates an expected call of GetInterchainAccountAddress.
func (mr *MockICAControllerKeeperMockRecorder) GetInterchainAccountAddress(ctx, connectionID, portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterchainAccountAddress", reflect.TypeOf((*MockICAControllerKeeper)(nil).GetInterchainAccountAddress), ctx, connectionID, portID)
}

// RegisterInterchainAccount mocks base method.
func (m *MockICAControllerKeeper) RegisterInterchainAccount(ctx types.Context, connectionID, owner string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInterchainAccount", ctx, connectionID, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterInterchainAccount indicates an expected call of RegisterInterchainAccount.
func (mr *MockICAControllerKeeperMockRecorder) RegisterInterchainAccount(ctx, connectionID, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterchainAccount", reflect.TypeOf((*MockICAControllerKeeper)(nil).RegisterInterchainAccount), ctx, connectionID, owner)
}

// SendTx mocks base method.
func (m *MockICAControllerKeeper) SendTx(ctx types.Context, chanCap *types1.Capability, connectionID, portID string, icaPacketData types2.InterchainAccountPacketData, timeoutTimestamp uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTx", ctx, chanCap, connectionID, portID, icaPacketData, timeoutTimestamp)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTx indicates an expected call of SendTx.
func (mr *MockICAControllerKeeperMockRecorder) SendTx(ctx, chanCap, connectionID, portID, icaPacketData, timeoutTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTx", reflect.TypeOf((*MockICAControllerKeeper)(nil).SendTx), ctx, chanCap, connectionID, portID, icaPacketData, timeoutTimestamp)
}

// MockFeeRefunderKeeper is a mock of FeeRefunderKeeper interface.
type MockFeeRefunderKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFeeRefunderKeeperMockRecorder
}

// MockFeeRefunderKeeperMockRecorder is the mock recorder for MockFeeRefunderKeeper.
type MockFeeRefunderKeeperMockRecorder struct {
	mock *MockFeeRefunderKeeper
}

// NewMockFeeRefunderKeeper creates a new mock instance.
func NewMockFeeRefunderKeeper(ctrl *gomock.Controller) *MockFeeRefunderKeeper {
	mock := &MockFeeRefunderKeeper{ctrl: ctrl}
	mock.recorder = &MockFeeRefunderKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeeRefunderKeeper) EXPECT() *MockFeeRefunderKeeperMockRecorder {
	return m.recorder
}

// DistributeAcknowledgementFee mocks base method.
func (m *MockFeeRefunderKeeper) DistributeAcknowledgementFee(ctx types.Context, receiver types.AccAddress, packetID types5.PacketID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DistributeAcknowledgementFee", ctx, receiver, packetID)
}

// DistributeAcknowledgementFee indicates an expected call of DistributeAcknowledgementFee.
func (mr *MockFeeRefunderKeeperMockRecorder) DistributeAcknowledgementFee(ctx, receiver, packetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeAcknowledgementFee", reflect.TypeOf((*MockFeeRefunderKeeper)(nil).DistributeAcknowledgementFee), ctx, receiver, packetID)
}

// DistributeTimeoutFee mocks base method.
func (m *MockFeeRefunderKeeper) DistributeTimeoutFee(ctx types.Context, receiver types.AccAddress, packetID types5.PacketID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DistributeTimeoutFee", ctx, receiver, packetID)
}

// DistributeTimeoutFee indicates an expected call of DistributeTimeoutFee.
func (mr *MockFeeRefunderKeeperMockRecorder) DistributeTimeoutFee(ctx, receiver, packetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTimeoutFee", reflect.TypeOf((*MockFeeRefunderKeeper)(nil).DistributeTimeoutFee), ctx, receiver, packetID)
}

// LockFees mocks base method.
func (m *MockFeeRefunderKeeper) LockFees(ctx types.Context, payer types.AccAddress, packetID types5.PacketID, fee types5.Fee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockFees", ctx, payer, packetID, fee)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockFees indicates an expected call of LockFees.
func (mr *MockFeeRefunderKeeperMockRecorder) LockFees(ctx, payer, packetID, fee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockFees", reflect.TypeOf((*MockFeeRefunderKeeper)(nil).LockFees), ctx, payer, packetID, fee)
}

// MockScopedKeeper is a mock of ScopedKeeper interface.
type MockScopedKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockScopedKeeperMockRecorder
}

// MockScopedKeeperMockRecorder is the mock recorder for MockScopedKeeper.
type MockScopedKeeperMockRecorder struct {
	mock *MockScopedKeeper
}

// NewMockScopedKeeper creates a new mock instance.
func NewMockScopedKeeper(ctrl *gomock.Controller) *MockScopedKeeper {
	mock := &MockScopedKeeper{ctrl: ctrl}
	mock.recorder = &MockScopedKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScopedKeeper) EXPECT() *MockScopedKeeperMockRecorder {
	return m.recorder
}

// ClaimCapability mocks base method.
func (m *MockScopedKeeper) ClaimCapability(ctx types.Context, cap *types1.Capability, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimCapability", ctx, cap, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimCapability indicates an expected call of ClaimCapability.
func (mr *MockScopedKeeperMockRecorder) ClaimCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimCapability", reflect.TypeOf((*MockScopedKeeper)(nil).ClaimCapability), ctx, cap, name)
}

// GetCapability mocks base method.
func (m *MockScopedKeeper) GetCapability(ctx types.Context, name string) (*types1.Capability, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapability", ctx, name)
	ret0, _ := ret[0].(*types1.Capability)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCapability indicates an expected call of GetCapability.
func (mr *MockScopedKeeperMockRecorder) GetCapability(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapability", reflect.TypeOf((*MockScopedKeeper)(nil).GetCapability), ctx, name)
}

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types.Context, srcPort, srcChan string) (types3.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types3.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// GetConnection mocks base method.
func (m *MockChannelKeeper) GetConnection(ctx types.Context, connectionID string) (exported.ConnectionI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", ctx, connectionID)
	ret0, _ := ret[0].(exported.ConnectionI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockChannelKeeperMockRecorder) GetConnection(ctx, connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockChannelKeeper)(nil).GetConnection), ctx, connectionID)
}

// GetNextSequenceSend mocks base method.
func (m *MockChannelKeeper) GetNextSequenceSend(ctx types.Context, portID, channelID string) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSequenceSend", ctx, portID, channelID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextSequenceSend indicates an expected call of GetNextSequenceSend.
func (mr *MockChannelKeeperMockRecorder) GetNextSequenceSend(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSequenceSend", reflect.TypeOf((*MockChannelKeeper)(nil).GetNextSequenceSend), ctx, portID, channelID)
}
