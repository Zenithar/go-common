package mock

import "github.com/stretchr/testify/mock"

type EventBus struct {
	mock.Mock
}

func (_m *EventBus) Subscribe(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *EventBus) SubscribeAsync(topic string, fn interface{}, transactional bool) error {
	ret := _m.Called(topic, fn, transactional)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, bool) error); ok {
		r0 = rf(topic, fn, transactional)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *EventBus) SubscribeOnce(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *EventBus) SubscribeOnceAsync(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *EventBus) HasCallback(topic string) bool {
	ret := _m.Called(topic)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(topic)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
func (_m *EventBus) Unsubscribe(topic string, handler interface{}) error {
	ret := _m.Called(topic, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *EventBus) Publish(topic string, args ...interface{}) {
	_m.Called(topic, args)
}
func (_m *EventBus) WaitAsync() {
	_m.Called()
}
