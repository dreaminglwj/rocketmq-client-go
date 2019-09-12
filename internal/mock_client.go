/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.dreaminglwj.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package internal is a generated GoMock package.
package internal

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/dreaminglwj/rocketmq-client-go/internal/remote"
	"github.com/dreaminglwj/rocketmq-client-go/primitive"
)

// MockInnerProducer is a mock of InnerProducer interface
type MockInnerProducer struct {
	ctrl     *gomock.Controller
	recorder *MockInnerProducerMockRecorder
}

// MockInnerProducerMockRecorder is the mock recorder for MockInnerProducer
type MockInnerProducerMockRecorder struct {
	mock *MockInnerProducer
}

// NewMockInnerProducer creates a new mock instance
func NewMockInnerProducer(ctrl *gomock.Controller) *MockInnerProducer {
	mock := &MockInnerProducer{ctrl: ctrl}
	mock.recorder = &MockInnerProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInnerProducer) EXPECT() *MockInnerProducerMockRecorder {
	return m.recorder
}

// PublishTopicList mocks base method
func (m *MockInnerProducer) PublishTopicList() []string {
	ret := m.ctrl.Call(m, "PublishTopicList")
	ret0, _ := ret[0].([]string)
	return ret0
}

// PublishTopicList indicates an expected call of PublishTopicList
func (mr *MockInnerProducerMockRecorder) PublishTopicList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTopicList", reflect.TypeOf((*MockInnerProducer)(nil).PublishTopicList))
}

// UpdateTopicPublishInfo mocks base method
func (m *MockInnerProducer) UpdateTopicPublishInfo(topic string, info *TopicPublishInfo) {
	m.ctrl.Call(m, "UpdateTopicPublishInfo", topic, info)
}

// UpdateTopicPublishInfo indicates an expected call of UpdateTopicPublishInfo
func (mr *MockInnerProducerMockRecorder) UpdateTopicPublishInfo(topic, info interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicPublishInfo", reflect.TypeOf((*MockInnerProducer)(nil).UpdateTopicPublishInfo), topic, info)
}

// IsPublishTopicNeedUpdate mocks base method
func (m *MockInnerProducer) IsPublishTopicNeedUpdate(topic string) bool {
	ret := m.ctrl.Call(m, "IsPublishTopicNeedUpdate", topic)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPublishTopicNeedUpdate indicates an expected call of IsPublishTopicNeedUpdate
func (mr *MockInnerProducerMockRecorder) IsPublishTopicNeedUpdate(topic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPublishTopicNeedUpdate", reflect.TypeOf((*MockInnerProducer)(nil).IsPublishTopicNeedUpdate), topic)
}

// IsUnitMode mocks base method
func (m *MockInnerProducer) IsUnitMode() bool {
	ret := m.ctrl.Call(m, "IsUnitMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnitMode indicates an expected call of IsUnitMode
func (mr *MockInnerProducerMockRecorder) IsUnitMode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnitMode", reflect.TypeOf((*MockInnerProducer)(nil).IsUnitMode))
}

// MockInnerConsumer is a mock of InnerConsumer interface
type MockInnerConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockInnerConsumerMockRecorder
}

// MockInnerConsumerMockRecorder is the mock recorder for MockInnerConsumer
type MockInnerConsumerMockRecorder struct {
	mock *MockInnerConsumer
}

// NewMockInnerConsumer creates a new mock instance
func NewMockInnerConsumer(ctrl *gomock.Controller) *MockInnerConsumer {
	mock := &MockInnerConsumer{ctrl: ctrl}
	mock.recorder = &MockInnerConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInnerConsumer) EXPECT() *MockInnerConsumerMockRecorder {
	return m.recorder
}

// PersistConsumerOffset mocks base method
func (m *MockInnerConsumer) PersistConsumerOffset() error {
	ret := m.ctrl.Call(m, "PersistConsumerOffset")
	ret0, _ := ret[0].(error)
	return ret0
}

// PersistConsumerOffset indicates an expected call of PersistConsumerOffset
func (mr *MockInnerConsumerMockRecorder) PersistConsumerOffset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistConsumerOffset", reflect.TypeOf((*MockInnerConsumer)(nil).PersistConsumerOffset))
}

// UpdateTopicSubscribeInfo mocks base method
func (m *MockInnerConsumer) UpdateTopicSubscribeInfo(topic string, mqs []*primitive.MessageQueue) {
	m.ctrl.Call(m, "UpdateTopicSubscribeInfo", topic, mqs)
}

// UpdateTopicSubscribeInfo indicates an expected call of UpdateTopicSubscribeInfo
func (mr *MockInnerConsumerMockRecorder) UpdateTopicSubscribeInfo(topic, mqs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicSubscribeInfo", reflect.TypeOf((*MockInnerConsumer)(nil).UpdateTopicSubscribeInfo), topic, mqs)
}

// IsSubscribeTopicNeedUpdate mocks base method
func (m *MockInnerConsumer) IsSubscribeTopicNeedUpdate(topic string) bool {
	ret := m.ctrl.Call(m, "IsSubscribeTopicNeedUpdate", topic)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSubscribeTopicNeedUpdate indicates an expected call of IsSubscribeTopicNeedUpdate
func (mr *MockInnerConsumerMockRecorder) IsSubscribeTopicNeedUpdate(topic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubscribeTopicNeedUpdate", reflect.TypeOf((*MockInnerConsumer)(nil).IsSubscribeTopicNeedUpdate), topic)
}

// SubscriptionDataList mocks base method
func (m *MockInnerConsumer) SubscriptionDataList() []*SubscriptionData {
	ret := m.ctrl.Call(m, "SubscriptionDataList")
	ret0, _ := ret[0].([]*SubscriptionData)
	return ret0
}

// SubscriptionDataList indicates an expected call of SubscriptionDataList
func (mr *MockInnerConsumerMockRecorder) SubscriptionDataList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionDataList", reflect.TypeOf((*MockInnerConsumer)(nil).SubscriptionDataList))
}

// Rebalance mocks base method
func (m *MockInnerConsumer) Rebalance() {
	m.ctrl.Call(m, "Rebalance")
}

// Rebalance indicates an expected call of Rebalance
func (mr *MockInnerConsumerMockRecorder) Rebalance() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebalance", reflect.TypeOf((*MockInnerConsumer)(nil).Rebalance))
}

// IsUnitMode mocks base method
func (m *MockInnerConsumer) IsUnitMode() bool {
	ret := m.ctrl.Call(m, "IsUnitMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUnitMode indicates an expected call of IsUnitMode
func (mr *MockInnerConsumerMockRecorder) IsUnitMode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUnitMode", reflect.TypeOf((*MockInnerConsumer)(nil).IsUnitMode))
}

// MockRMQClient is a mock of RMQClient interface
type MockRMQClient struct {
	ctrl     *gomock.Controller
	recorder *MockRMQClientMockRecorder
}

// MockRMQClientMockRecorder is the mock recorder for MockRMQClient
type MockRMQClientMockRecorder struct {
	mock *MockRMQClient
}

// NewMockRMQClient creates a new mock instance
func NewMockRMQClient(ctrl *gomock.Controller) *MockRMQClient {
	mock := &MockRMQClient{ctrl: ctrl}
	mock.recorder = &MockRMQClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRMQClient) EXPECT() *MockRMQClientMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockRMQClient) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockRMQClientMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockRMQClient)(nil).Start))
}

// Shutdown mocks base method
func (m *MockRMQClient) Shutdown() {
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockRMQClientMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockRMQClient)(nil).Shutdown))
}

// ClientID mocks base method
func (m *MockRMQClient) ClientID() string {
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID
func (mr *MockRMQClientMockRecorder) ClientID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockRMQClient)(nil).ClientID))
}

// RegisterProducer mocks base method
func (m *MockRMQClient) RegisterProducer(group string, producer InnerProducer) {
	m.ctrl.Call(m, "RegisterProducer", group, producer)
}

// RegisterProducer indicates an expected call of RegisterProducer
func (mr *MockRMQClientMockRecorder) RegisterProducer(group, producer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterProducer", reflect.TypeOf((*MockRMQClient)(nil).RegisterProducer), group, producer)
}

// InvokeSync mocks base method
func (m *MockRMQClient) InvokeSync(ctx context.Context, addr string, request *remote.RemotingCommand, timeoutMillis time.Duration) (*remote.RemotingCommand, error) {
	ret := m.ctrl.Call(m, "InvokeSync", ctx, addr, request, timeoutMillis)
	ret0, _ := ret[0].(*remote.RemotingCommand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeSync indicates an expected call of InvokeSync
func (mr *MockRMQClientMockRecorder) InvokeSync(ctx, addr, request, timeoutMillis interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeSync", reflect.TypeOf((*MockRMQClient)(nil).InvokeSync), ctx, addr, request, timeoutMillis)
}

// InvokeAsync mocks base method
func (m *MockRMQClient) InvokeAsync(ctx context.Context, addr string, request *remote.RemotingCommand, timeoutMillis time.Duration, f func(*remote.RemotingCommand, error)) error {
	ret := m.ctrl.Call(m, "InvokeAsync", ctx, addr, request, timeoutMillis, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeAsync indicates an expected call of InvokeAsync
func (mr *MockRMQClientMockRecorder) InvokeAsync(ctx, addr, request, timeoutMillis, f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeAsync", reflect.TypeOf((*MockRMQClient)(nil).InvokeAsync), ctx, addr, request, timeoutMillis, f)
}

// InvokeOneWay mocks base method
func (m *MockRMQClient) InvokeOneWay(ctx context.Context, addr string, request *remote.RemotingCommand, timeoutMillis time.Duration) error {
	ret := m.ctrl.Call(m, "InvokeOneWay", ctx, addr, request, timeoutMillis)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeOneWay indicates an expected call of InvokeOneWay
func (mr *MockRMQClientMockRecorder) InvokeOneWay(ctx, addr, request, timeoutMillis interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeOneWay", reflect.TypeOf((*MockRMQClient)(nil).InvokeOneWay), ctx, addr, request, timeoutMillis)
}

// CheckClientInBroker mocks base method
func (m *MockRMQClient) CheckClientInBroker() {
	m.ctrl.Call(m, "CheckClientInBroker")
}

// CheckClientInBroker indicates an expected call of CheckClientInBroker
func (mr *MockRMQClientMockRecorder) CheckClientInBroker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckClientInBroker", reflect.TypeOf((*MockRMQClient)(nil).CheckClientInBroker))
}

// SendHeartbeatToAllBrokerWithLock mocks base method
func (m *MockRMQClient) SendHeartbeatToAllBrokerWithLock() {
	m.ctrl.Call(m, "SendHeartbeatToAllBrokerWithLock")
}

// SendHeartbeatToAllBrokerWithLock indicates an expected call of SendHeartbeatToAllBrokerWithLock
func (mr *MockRMQClientMockRecorder) SendHeartbeatToAllBrokerWithLock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeartbeatToAllBrokerWithLock", reflect.TypeOf((*MockRMQClient)(nil).SendHeartbeatToAllBrokerWithLock))
}

// UpdateTopicRouteInfo mocks base method
func (m *MockRMQClient) UpdateTopicRouteInfo() {
	m.ctrl.Call(m, "UpdateTopicRouteInfo")
}

// UpdateTopicRouteInfo indicates an expected call of UpdateTopicRouteInfo
func (mr *MockRMQClientMockRecorder) UpdateTopicRouteInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTopicRouteInfo", reflect.TypeOf((*MockRMQClient)(nil).UpdateTopicRouteInfo))
}

// ProcessSendResponse mocks base method
func (m *MockRMQClient) ProcessSendResponse(brokerName string, cmd *remote.RemotingCommand, resp *primitive.SendResult, msgs ...*primitive.Message) error {
	varargs := []interface{}{brokerName, cmd, resp}
	for _, a := range msgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProcessSendResponse", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessSendResponse indicates an expected call of ProcessSendResponse
func (mr *MockRMQClientMockRecorder) ProcessSendResponse(brokerName, cmd, resp interface{}, msgs ...interface{}) *gomock.Call {
	varargs := append([]interface{}{brokerName, cmd, resp}, msgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessSendResponse", reflect.TypeOf((*MockRMQClient)(nil).ProcessSendResponse), varargs...)
}

// RegisterConsumer mocks base method
func (m *MockRMQClient) RegisterConsumer(group string, consumer InnerConsumer) error {
	ret := m.ctrl.Call(m, "RegisterConsumer", group, consumer)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterConsumer indicates an expected call of RegisterConsumer
func (mr *MockRMQClientMockRecorder) RegisterConsumer(group, consumer interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConsumer", reflect.TypeOf((*MockRMQClient)(nil).RegisterConsumer), group, consumer)
}

// UnregisterConsumer mocks base method
func (m *MockRMQClient) UnregisterConsumer(group string) {
	m.ctrl.Call(m, "UnregisterConsumer", group)
}

// UnregisterConsumer indicates an expected call of UnregisterConsumer
func (mr *MockRMQClientMockRecorder) UnregisterConsumer(group interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterConsumer", reflect.TypeOf((*MockRMQClient)(nil).UnregisterConsumer), group)
}

// PullMessage mocks base method
func (m *MockRMQClient) PullMessage(ctx context.Context, brokerAddrs string, request *PullMessageRequest) (*primitive.PullResult, error) {
	ret := m.ctrl.Call(m, "PullMessage", ctx, brokerAddrs, request)
	ret0, _ := ret[0].(*primitive.PullResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullMessage indicates an expected call of PullMessage
func (mr *MockRMQClientMockRecorder) PullMessage(ctx, brokerAddrs, request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullMessage", reflect.TypeOf((*MockRMQClient)(nil).PullMessage), ctx, brokerAddrs, request)
}

// PullMessageAsync mocks base method
func (m *MockRMQClient) PullMessageAsync(ctx context.Context, brokerAddrs string, request *PullMessageRequest, f func(*primitive.PullResult)) error {
	ret := m.ctrl.Call(m, "PullMessageAsync", ctx, brokerAddrs, request, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullMessageAsync indicates an expected call of PullMessageAsync
func (mr *MockRMQClientMockRecorder) PullMessageAsync(ctx, brokerAddrs, request, f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullMessageAsync", reflect.TypeOf((*MockRMQClient)(nil).PullMessageAsync), ctx, brokerAddrs, request, f)
}

// RebalanceImmediately mocks base method
func (m *MockRMQClient) RebalanceImmediately() {
	m.ctrl.Call(m, "RebalanceImmediately")
}

// RebalanceImmediately indicates an expected call of RebalanceImmediately
func (mr *MockRMQClientMockRecorder) RebalanceImmediately() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebalanceImmediately", reflect.TypeOf((*MockRMQClient)(nil).RebalanceImmediately))
}

// UpdatePublishInfo mocks base method
func (m *MockRMQClient) UpdatePublishInfo(topic string, data *TopicRouteData) {
	m.ctrl.Call(m, "UpdatePublishInfo", topic, data)
}

// UpdatePublishInfo indicates an expected call of UpdatePublishInfo
func (mr *MockRMQClientMockRecorder) UpdatePublishInfo(topic, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePublishInfo", reflect.TypeOf((*MockRMQClient)(nil).UpdatePublishInfo), topic, data)
}
