package version1

import (
	"os"
	"time"

	sdata1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	logic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	clients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

type EventLogDirectClientV1 struct {
	clients.DirectClient
	controller logic.IEventLogController
}

func NewEventLogDirectClientV1() *EventLogDirectClientV1 {
	c := EventLogDirectClientV1{}
	c.DirectClient = *clients.NewDirectClient()
	c.DependencyResolver.Put("controller", cref.NewDescriptor("pip-services-eventlog", "controller", "*", "*", "1.0"))
	return &c
}

func (c *EventLogDirectClientV1) SetReferences(references cref.IReferences) {
	c.DirectClient.SetReferences(references)

	controller, ok := c.Controller.(logic.IEventLogController)
	if !ok {
		panic("EventLogDirectClientV1: Cant't resolv dependency 'controller' to IEventLogClientV1")
	}
	c.controller = controller
}

func toServerObject(value *SystemEventV1) *sdata1.SystemEventV1 {
	if value == nil {
		return nil
	}

	return &sdata1.SystemEventV1{
		Id:            value.Id,
		Time:          value.Time,
		CorrelationId: value.CorrelationId,
		Source:        value.Source,
		Type:          value.Type,
		Severity:      value.Severity,
		Message:       value.Message,
		Details:       value.Details,
	}
}

func fromServerObject(value *sdata1.SystemEventV1) *SystemEventV1 {
	if value == nil {
		return nil
	}

	return &SystemEventV1{
		Id:            value.Id,
		Time:          value.Time,
		CorrelationId: value.CorrelationId,
		Source:        value.Source,
		Type:          value.Type,
		Severity:      value.Severity,
		Message:       value.Message,
		Details:       value.Details,
	}
}

func fromServerPage(value *sdata1.SystemEventV1DataPage) *SystemEventV1DataPage {
	if value == nil {
		return nil
	}

	data := make([]*SystemEventV1, len(value.Data))
	for i, v := range value.Data {
		data[i] = fromServerObject(v)
	}

	return NewSystemEventV1DataPage(value.Total, data)
}

func (c *EventLogDirectClientV1) GetEvents(correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*SystemEventV1DataPage, error) {
	timing := c.Instrument(correlationId, "eventlog.get_events")
	res, err := c.controller.GetEvents(correlationId, filter, paging)
	timing.EndTiming()
	return fromServerPage(res), err
}

func (c *EventLogDirectClientV1) LogEvent(correlationId string, event *SystemEventV1) error {
	event.Time = time.Now()
	if event.Source == "" {
		event.Source, _ = os.Hostname()
	}

	timing := c.Instrument(correlationId, "eventlog.log_event")
	p1 := toServerObject(event)
	err := c.controller.LogEvent(correlationId, p1)
	timing.EndTiming()
	return err
}
