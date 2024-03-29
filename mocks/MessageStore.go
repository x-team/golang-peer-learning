// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	conversation "golang/golang-study/internal/conversation"

	mock "github.com/stretchr/testify/mock"
)

// MessageStore is an autogenerated mock type for the MessageStore type
type MessageStore struct {
	mock.Mock
}

// DeleteMessage provides a mock function with given fields: ctx, uuid
func (_m *MessageStore) DeleteMessage(ctx context.Context, uuid string) (bool, error) {
	ret := _m.Called(ctx, uuid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMessage")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: _a0, _a1
func (_m *MessageStore) GetMessage(_a0 context.Context, _a1 string) (conversation.Message, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetMessage")
	}

	var r0 conversation.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (conversation.Message, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) conversation.Message); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(conversation.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostMessage provides a mock function with given fields: _a0, _a1
func (_m *MessageStore) PostMessage(_a0 context.Context, _a1 conversation.Message) (conversation.Message, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PostMessage")
	}

	var r0 conversation.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, conversation.Message) (conversation.Message, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, conversation.Message) conversation.Message); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(conversation.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, conversation.Message) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMessage provides a mock function with given fields: ctx, id, message
func (_m *MessageStore) UpdateMessage(ctx context.Context, id string, message conversation.Message) (conversation.Message, error) {
	ret := _m.Called(ctx, id, message)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMessage")
	}

	var r0 conversation.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, conversation.Message) (conversation.Message, error)); ok {
		return rf(ctx, id, message)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, conversation.Message) conversation.Message); ok {
		r0 = rf(ctx, id, message)
	} else {
		r0 = ret.Get(0).(conversation.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, conversation.Message) error); ok {
		r1 = rf(ctx, id, message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMessageStore creates a new instance of MessageStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageStore {
	mock := &MessageStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
