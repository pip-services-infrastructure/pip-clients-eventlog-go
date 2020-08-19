package build

import (
	"github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type EventLogClientFactory struct {
	cbuild.Factory
}

func NewEventLogClientFactory() *EventLogClientFactory {
	c := EventLogClientFactory{}
	c.Factory = *cbuild.NewFactory()

	nullClientDescriptor := cref.NewDescriptor("pip-services-eventlog", "client", "null", "*", "1.0")
	directClientDescriptor := cref.NewDescriptor("pip-services-eventlog", "client", "direct", "*", "1.0")
	httpClientDescriptor := cref.NewDescriptor("pip-services-eventlog", "client", "http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, version1.NewEventLogNullClientV1)
	c.RegisterType(directClientDescriptor, version1.NewEventLogDirectClientV1)
	c.RegisterType(httpClientDescriptor, version1.NewEventLogHttpClientV1)
	return &c
}
