package test_version1

import (
	"testing"

	version1 "github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
	slogic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	spersist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

func TestEventLogDirectClientV1(t *testing.T) {
	var persistence *spersist.EventLogMemoryPersistence
	var controller *slogic.EventLogController
	var client *version1.EventLogDirectClientV1
	var fixture *EventLogClientV1Fixture

	persistence = spersist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = slogic.NewEventLogController()
	controller.Configure(cconf.NewEmptyConfigParams())
	client = version1.NewEventLogDirectClientV1()

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
