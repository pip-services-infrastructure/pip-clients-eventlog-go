let assert = require('chai').assert;
let async = require('async');

import { Descriptor } from 'pip-services3-commons-node';
import { ConfigParams } from 'pip-services3-commons-node';
import { References } from 'pip-services3-commons-node';
import { ConsoleLogger } from 'pip-services3-components-node';

import { EventLogMemoryPersistence } from 'pip-services-eventlog-node';
import { EventLogController } from 'pip-services-eventlog-node';
import { IEventLogClientV1 } from '../../src/version1/IEventLogClientV1';
import { EventLogDirectClientV1 } from '../../src/version1/EventLogDirectClientV1';
import { EventLogClientFixtureV1 } from './EventLogClientFixtureV1';

suite('EventLogDirectClientV1', ()=> {
    let client: EventLogDirectClientV1;
    let fixture: EventLogClientFixtureV1;

    suiteSetup((done) => {
        let logger = new ConsoleLogger();
        let persistence = new EventLogMemoryPersistence();
        let controller = new EventLogController();

        let references: References = References.fromTuples(
            new Descriptor('pip-services', 'logger', 'console', 'default', '1.0'), logger,
            new Descriptor('pip-services-eventlog', 'persistence', 'memory', 'default', '1.0'), persistence,
            new Descriptor('pip-services-eventlog', 'controller', 'default', 'default', '1.0'), controller,
        );
        controller.setReferences(references);
        client = new EventLogDirectClientV1();

        client.setReferences(references);

        fixture = new EventLogClientFixtureV1(client);

        client.open(null, done);
    });
    
    suiteTeardown((done) => {
        client.close(null, done);
    });

    test('CRUD Operations', (done) => {
        fixture.testCrudOperations(done);
    });

});
