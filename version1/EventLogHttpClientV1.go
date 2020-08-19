package version1

import (
	"os"
	"reflect"
	"time"

	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	clients "github.com/pip-services3-go/pip-services3-rpc-go/clients"
)

var (
	systemEventV1DataPageType = reflect.TypeOf(&SystemEventV1DataPage{})
	systemEventV1Type         = reflect.TypeOf(&SystemEventV1{})
)

type EventLogHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewEventLogHttpClientV1() *EventLogHttpClientV1 {
	c := EventLogHttpClientV1{}
	c.CommandableHttpClient = clients.NewCommandableHttpClient("v1/eventlog")
	return &c
}

func (c *EventLogHttpClientV1) GetEvents(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*SystemEventV1DataPage, error) {
	params := cdata.NewEmptyStringValueMap()
	c.AddFilterParams(params, filter)
	c.AddPagingParams(params, paging)

	val, err := c.CallCommand(systemEventV1DataPageType, "get_events", correlationId, params, nil)
	if err != nil {
		return nil, err
	}
	page, _ := val.(*SystemEventV1DataPage)
	return page, nil
}

func (c *EventLogHttpClientV1) LogEvent(correlationId string, event *SystemEventV1) error {
	event.Time = time.Now()
	if event.Source == "" {
		event.Source, _ = os.Hostname()
	}

	params := cdata.NewAnyValueMapFromTuples(
		"event", event,
	)

	_, err := c.CallCommand(nil, "log_event", correlationId, nil, params.Value())
	return err
}
