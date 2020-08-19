let _ = require('lodash');
let async = require('async');
let assert = require('chai').assert;

import { SystemEventV1 } from '../../src/version1/SystemEventV1';
import { EventLogTypeV1 } from '../../src/version1/EventLogTypeV1';
import { EventLogSeverityV1 } from '../../src/version1/EventLogSeverityV1';
import { IEventLogClientV1 } from '../../src/version1/IEventLogClientV1';

let EVENT1: SystemEventV1 = new SystemEventV1(
    null,
    'test',
    EventLogTypeV1.Restart,
    EventLogSeverityV1.Important,
    'test restart #1'
);
let EVENT2: SystemEventV1 = new SystemEventV1(
    null,
    'test',
    EventLogTypeV1.Failure,
    EventLogSeverityV1.Critical,
    'test error'
);

export class EventLogClientFixtureV1 {
    private _client: IEventLogClientV1;
    
    constructor(client: IEventLogClientV1) {
        this._client = client;
    }
        
    testCrudOperations(done) {
        let event1;
        let event2;

        async.series([
        // Create one event
            (callback) => {
                this._client.logEvent(
                    null,
                    EVENT1,
                    (err, event) => {
                        assert.isNull(err);
                        
                        assert.isObject(event);
                        assert.isNotNull(event.time);
                        assert.isNotNull(event.source);
                        assert.equal(event.type, EVENT1.type);
                        assert.equal(event.message, EVENT1.message);

                        event1 = event;

                        callback();
                    }
                );
            },
        // Create another event
            (callback) => {
                this._client.logEvent(
                    null,
                    EVENT2,
                    (err, event) => {
                        assert.isNull(err);
                        
                        assert.isObject(event);
                        assert.isNotNull(event.time);
                        assert.isNotNull(event.source);
                        assert.equal(event.type, EVENT2.type);
                        assert.equal(event.message, EVENT2.message);

                        event2 = event;

                        callback();
                    }
                );
            },
        // Get all system events
            (callback) => {
                this._client.getEvents(
                    null,
                    null,
                    null,
                    (err, page) => {
                        assert.isNull(err);
                        
                        assert.isObject(page);
                        assert.lengthOf(page.data, 2);

                        callback();
                    }
                );
            }
        ], done);
    }
}
