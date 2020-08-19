# Client API (version 1) <br/> EventLog Microservices Client SDK for Node.js

Node.js client API for EventLog microservice is a thin layer on the top of
communication protocols. It hides details related to specific protocol implementation
and provides high-level API to access the microservice for simple and productive development.

* [Installation](#install)
* [Getting started](#get_started)
* [SystemEventV1 class](#class1)
* [IEventLogClientV1 interface](#interface)
    - [getEvents()](#operation1)
    - [logEvent()](#operation2)
* [EventLogHttpClientV1 class](#client_http)
* [EventLogSenecaClientV1 class](#client_seneca)
* [EventLogDirectClientV1 class](#client_direct)
* [EventLogNullClientV1 class](#client_null)

## <a name="install"></a> Installation

To work with the client SDK add dependency into package.json file:

```javascript
{
    ...
    "dependencies": {
        ....
        "pip-clients-eventlog-node": "^1.0.0",
        ...
    }
}
```

Then download the dependency using **npm**:

```javascript
# Installing dependencies
npm install

# Updating dependencies
npm update
```

## <a name="get_started"></a> Getting started

This is a simple example on how to work with the microservice using REST client:

```javascript
// Get Client SDK for Version 1 
var sdk = new require('pip-clients-eventlog-node');

// Client configuration
var config = {
    connection: {
        protocol: 'http',
        host: 'localhost', 
        port: 8080
    }
};

// Create the client instance
var client = sdk.EventLogHttpClientV1(config);

// Open client connection to the microservice
client.open(null, function(err) {
    if (err) {
        console.error(err);
        return; 
    }
    
    console.log('Opened connection');
        
    // Log system event
    client.logEvent(
        null,
        {
            type: 'restart',
            source: 'server1',
            severity: 500,
            message: 'Server restarted'
        }, 
        function (err, event) {
            if (err) {
                console.error(err);
                return;
            }
            
            console.log('Logged system event is');
            console.log(event);
            
            var now = new Date();
    
            // Read server events
            client.getEvents(
                {
                    source: 'server1',
                    from_time: new Date(now.getTime() - 24 * 3600 * 1000),
                    to_time: now
                },
                {
                    total: true,
                    skip: 0, 
                    take: 100
                },
                function (err, page) {
                    if (err) {
                        console.error(err);
                        return;
                    }
                    
                    console.log('Events for server 1 were');
                    console.log(page.data);
                    
                    // Close connection
                    client.close(); 
                }
            );
        }
    );
});
```

### <a name="class1"></a> SystemEventV1 class

Represents a record of a system activity performed in the past

**Properties:**
- id: string - unique record id
- correlation_id: string - unique id of transaction that caused the event
- time: Date - date and time in UTC when the event took place (default: current time)
- source: string - server name where event took place (default: current host)
- type: string - event type: 'restart', 'upgrade', 'shutdown', 'transaction' etc.
- severity: number - severity level (impact on system operations) from 0: Low to 1000: High
- message: string - descriptive message
- details: Object - additional details that can help system administrators in troubleshooting

## <a name="interface"></a> IEventLogClientV1 interface

If you are using Typescript, you can use IEventLogClientV1 as a common interface across all client implementations. 
If you are using plain Javascript, you shall not worry about IEventLogClientV1 interface. You can just expect that
all methods defined in this interface are implemented by all client classes.

```javascript
interface IEventLogClientV1 {
    getEvents(correlationId, filter, paging, callback);
    logEvent(correlationId, event, callback);
}
```

### <a name="operation1"></a> getEvents(correlationId, filter, paging, callback)

Retrieves system events by specified criteria

**Arguments:** 
- correlationId: string - id that uniquely identifies transaction
- filter: object - filter parameters
  - search: string - (optional) search substring to find in source, type or message
  - type: string - (optional) type events
  - source: string - (optional) server where events occured
  - severity: number - (optional) severity of events
  - from_time: Date - (optional) start of the time range
  - to_time: Date - (optional) end of the time range
- paging: object - paging parameters
  - skip: int - (optional) start of page (default: 0)
  - take: int - (optional) page length (default: 100)
  - total: boolean - (optional) include total counter into paged result (default: false)
- callback: (err, page) - callback function
  - err: Error - occured error or null for success
  - page: DataPage<SystemEventV1> - retrieved SystemEventV1 objects in paged format

### <a name="operation2"></a> logEvent(correlationId, event, callback)

Log system event

**Activities:** 
- correlationId: string - id that uniquely identifies transaction
- event: SystemEventV1 - system evemt to be logged
- callback: (err, event) => void - callback function
  - err: Error - occured error or null for success
  - event: SystemEventV1 - logged system event
 
## <a name="client_http"></a> EventLogHttpClientV1 class

EventLogHttpClientV1 is a client that implements HTTP protocol

```javascript
class EventLogHttpClientV1 extends CommandableHttpClient implements IEventLogClientV1 {
    constructor(config?: any);
    setReferences(references);
    open(correlationId, callback);
    close(correlationId, callback);
    getEvents(correlationId, filter, paging, callback);
    logEvent(correlationId, event, callback);
}
```

**Constructor config properties:** 
- connection: object - HTTP transport configuration options
  - type: string - HTTP protocol - 'http' or 'https' (default is 'http')
  - host: string - IP address/hostname binding (default is '0.0.0.0')
  - port: number - HTTP port number

## <a name="client_seneca"></a> EventLogSenecaClientV1 class

EventLogSenecaClientV1 is a client that implements Seneca protocol

```javascript
class EventLogSenecaClientV1 extends CommandableSenecaClient implements IEventLogClientV1 {
    constructor(config?: any);
    setReferences(references);
    open(correlationId, callback);
    close(correlationId, callback);
    getEvents(correlationId, filter, paging, callback);
    logEvent(correlationId, event, callback);
}
```

**Constructor config properties:** 
- connection: object - (optional) Seneca transport configuration options. See http://senecajs.org/api/ for details.
  - type: string - Seneca transport type 
  - host: string - IP address/hostname binding (default is '0.0.0.0')
  - port: number - Seneca port number

## <a name="client_direct"></a> EventLogDirectClientV1 class

EventLogDirectClientV1 is a client that calls controller directly from the same container.
It can be used in monolythic deployments when multiple microservices run in the same process.

```javascript
class EventLogDirectClientV1 extends DirectClient implements IEventLogClientV1 {
    constructor();        
    setReferences(references);
    open(correlationId, callback);
    close(correlationId, callback);
    getEvents(correlationId, filter, paging, callback);
    logEvent(correlationId, event, callback);
}
```

## <a name="client_null"></a> EventLogNullClientV1 class

EventLogNullClientV1 is a dummy client that mimics the real client but doesn't call a microservice. 
It can be useful in testing scenarios to cut dependencies on external microservices.

```javascript
class EventLogNullClientV1 implements IEventLogClientV1 {
    constructor();
    getEvents(correlationId, filter, paging, callback);
    logEvent(correlationId, event, callback);
}
```
