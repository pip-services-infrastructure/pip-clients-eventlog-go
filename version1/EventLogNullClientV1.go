package version1

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type EventLogNullClientV1 struct {
}

func NewEventLogNullClientV1() *EventLogNullClientV1 {
	return &EventLogNullClientV1{}
}

func (c *EventLogNullClientV1) GetEvents(correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*SystemEventV1DataPage, error) {
	return NewEmptySystemEventV1DataPage(), nil
}

func (c *EventLogNullClientV1) LogEvent(correlationId string, event *SystemEventV1) error {
	return nil
}
