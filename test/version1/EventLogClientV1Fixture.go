package test_version1

import (
	"testing"

	version1 "github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
	"github.com/stretchr/testify/assert"
)

type EventLogClientV1Fixture struct {
	Event1 *version1.SystemEventV1
	Event2 *version1.SystemEventV1
	client version1.IEventLogClientV1
}

func NewEventLogClientV1Fixture(client version1.IEventLogClientV1) *EventLogClientV1Fixture {
	c := EventLogClientV1Fixture{}
	c.Event1 = &version1.SystemEventV1{
		Id:       "1",
		Source:   "test",
		Type:     version1.Restart,
		Severity: version1.Important,
		Message:  "test restart #1",
	}
	c.Event2 = &version1.SystemEventV1{
		Id:       "2",
		Source:   "test",
		Type:     version1.Failure,
		Severity: version1.Critical,
		Message:  "test error",
	}
	c.client = client
	return &c
}

func (c *EventLogClientV1Fixture) TestCrudOperations(t *testing.T) {
	// Create one event
	err := c.client.LogEvent("", c.Event1)
	assert.Nil(t, err)

	// Create another event
	err = c.client.LogEvent("", c.Event2)
	assert.Nil(t, err)

	// Get all system events
	page, err1 := c.client.GetEvents("", nil, nil)
	assert.Nil(t, err1)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	event1 := page.Data[0]
	assert.NotNil(t, event1)
	assert.Equal(t, c.Event1.Id, event1.Id)
	assert.Equal(t, c.Event1.Source, event1.Source)
	assert.Equal(t, c.Event1.Type, event1.Type)
	assert.Equal(t, c.Event1.Message, event1.Message)

	event2 := page.Data[1]
	assert.NotNil(t, event2)
	assert.Equal(t, c.Event2.Id, event2.Id)
	assert.Equal(t, c.Event2.Source, event2.Source)
	assert.Equal(t, c.Event2.Type, event2.Type)
	assert.Equal(t, c.Event2.Message, event2.Message)
}
