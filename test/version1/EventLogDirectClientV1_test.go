package test_clients1

import (
	"testing"

	clients1 "github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
	logic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestEventLogDirectClientV1(t *testing.T) {
	var persistence *persist.EventLogMemoryPersistence
	var controller *logic.EventLogController
	var client *clients1.EventLogDirectClientV1
	var fixture *EventLogClientV1Fixture

	persistence = persist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = logic.NewEventLogController()
	controller.Configure(cconf.NewEmptyConfigParams())
	client = clients1.NewEventLogDirectClientV1()

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-eventlog", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("pip-services-eventlog", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("pip-services-eventlog", "client", "direct", "default", "1.0"), client,
	)

	controller.SetReferences(references)
	client.SetReferences(references)
	fixture = NewEventLogClientV1Fixture(client)

	persistence.Open("")
	defer persistence.Close("")

	t.Run("TestEventLogDirectClientV1:CRUD Operations", fixture.TestCrudOperations)
}
