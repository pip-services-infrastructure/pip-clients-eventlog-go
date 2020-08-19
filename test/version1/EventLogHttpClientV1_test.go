package test_version1

import (
	"testing"

	version1 "github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
	slogic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	spersist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	sservices1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/services/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestEventLogHttpClientV1(t *testing.T) {
	var persistence *spersist.EventLogMemoryPersistence
	var controller *slogic.EventLogController
	var service *sservices1.EventLogHttpServiceV1
	var client *version1.EventLogHttpClientV1
	var fixture *EventLogClientV1Fixture

	persistence = spersist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = slogic.NewEventLogController()
	controller.Configure(cconf.NewEmptyConfigParams())

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

	service = sservices1.NewEventLogHttpServiceV1()
	service.Configure(httpConfig)

	client = version1.NewEventLogHttpClientV1()
	client.Configure(httpConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-eventlog", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("pip-services-eventlog", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("pip-services-eventlog", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("pip-services-eventlog", "client", "direct", "default", "1.0"), client,
	)

	controller.SetReferences(references)
	service.SetReferences(references)
	client.SetReferences(references)
	fixture = NewEventLogClientV1Fixture(client)

	err := persistence.Open("")
	if err != nil {
		panic("TestEventLogHttpClientV1:Error open persistence!")
	}

	err = service.Open("")
	if err != nil {
		panic("TestEventLogHttpClientV1:Error open service!")
	}

	client.Open("")

	defer client.Close("")
	defer service.Close("")
	defer persistence.Close("")

	t.Run("TestEventLogHttpClientV1:CRUD Operations", fixture.TestCrudOperations)
}
