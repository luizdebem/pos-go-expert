package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = TestEvent{
		Name:    "test",
		Payload: nil,
	}

	suite.event2 = TestEvent{
		Name:    "test2",
		Payload: nil,
	}

	suite.handler = TestEventHandler{
		ID: 1,
	}

	suite.handler2 = TestEventHandler{
		ID: 2,
	}

	suite.handler3 = TestEventHandler{
		ID: 3,
	}

	suite.eventDispatcher = NewEventDispatcher()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	_ = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	_ = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler2)
	_ = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler3)

	suite.eventDispatcher.Clear()

	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	_ = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)

	suite.True(suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	suite.False(suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	_ = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	_ = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	_ = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	_ = suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
