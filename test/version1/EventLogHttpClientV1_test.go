package test_clients1

import (
	"testing"

	clients1 "github.com/NationalOilwellVarco/max-system/client-eventlog-go/version1"
	logic "github.com/NationalOilwellVarco/max-system/service-eventlog-go/logic"
	persist "github.com/NationalOilwellVarco/max-system/service-eventlog-go/persistence"
	services1 "github.com/NationalOilwellVarco/max-system/service-eventlog-go/services/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestEventLogHttpClientV1(t *testing.T) {
	var persistence *persist.EventLogMemoryPersistence
	var controller *logic.EventLogController
	var service *services1.EventLogHttpServiceV1
	var client *clients1.EventLogHttpClientV1
	var fixture *EventLogClientV1Fixture

	persistence = persist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = logic.NewEventLogController()
	controller.Configure(cconf.NewEmptyConfigParams())

	httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

	service = services1.NewEventLogHttpServiceV1()
	service.Configure(httpConfig)

	client = clients1.NewEventLogHttpClientV1()
	client.Configure(httpConfig)

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("nov-max-system-eventlog", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("nov-max-system-eventlog", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("nov-max-system-eventlog", "service", "http", "default", "1.0"), service,
		cref.NewDescriptor("nov-max-system-eventlog", "client", "direct", "default", "1.0"), client,
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
