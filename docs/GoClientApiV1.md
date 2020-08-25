# Client API (version 1) <br/> EventLog Microservices Client SDK for Golang

Golang client API for EventLog microservice is a thin layer on the top of
communication protocols. It hides details related to specific protocol implementation
and provides high-level API to access the microservice for simple and productive development.

* [Installation](#install)
* [Getting started](#get_started)
* [SystemEventV1 class](#class1)
* [IEventLogClientV1 interface](#interface)
    - [getEvents()](#operation1)
    - [logEvent()](#operation2)
* [EventLogHttpClientV1 class](#client_http)
* [EventLogDirectClientV1 class](#client_direct)
* [EventLogNullClientV1 class](#client_null)

## <a name="install"></a> Installation

To work with the client SDK add dependency into go.mod file:

``` golang
...
require (

    github.com/pip-services-infrastructure/pip-services-eventlog-go v1.0.0
    ....
)
```

# Update source code updates from github

``` bash
go get -u github.com/pip-services-infrastructure/pip-services-eventlog-go@latest
```

## <a name="get_started"></a> Getting started

This is a simple example on how to work with the microservice using REST client:

Inside your code get the reference to the client SDK

``` golang
import (
	clients1 "github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
)
// Get Client SDK for Version 1 
var client *clients1.EventLogHttpClientV1

// Client configuration
httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

client = clients1.NewEventLogHttpClientV1()
client.Configure(httpConfig)

// Connect to the microservice
err := client.Open("")
 if (err) {
        panic("Connection to the microservice failed");
    }
defer client.Close("")

// Open client connection to the microservice
client.open(null, function(err) {
    if (err) {
        console.error(err);
        return; 
    }

// Log system event
event1:=&clients1.SystemEventV1{
        Type: "restart",
        source: "server1",
        Message: "Restarted server",
    }

err := client.LogEvent(
    "",
    event1,
);

if (err != nil) {
    print(err);
}

var now = time.Now();

// Get the list system events
page, err1 := client.getEvents(
    "",
    cdata.NewFilterParamsFromTuples(
        "from_time": new Date(now.getTime() - 24 * 3600 * 1000),
        "to_time": now,
        "source": "server1"
    ), cdata.NewEmptyPagingParams(),
);

 if (err1 != nil) {
    print(err1);
}
                    
print("Events for server 1 were");
print(page.Data);

```

### <a name="class1"></a> SystemEventV1 class

Represents a record of a system activity performed in the past

**Properties:**

* id: string - unique record id
* correlation_id: string - unique id of transaction that caused the event
* time: Date - date and time in UTC when the event took place (default: current time)
* source: string - server name where event took place (default: current host)
* type: string - event type: 'restart', 'upgrade', 'shutdown', 'transaction' etc.
* severity: number - severity level (impact on system operations) from 0: Low to 1000: High
* message: string - descriptive message
* details: Object - additional details that can help system administrators in troubleshooting

## <a name="interface"></a> IEventLogClientV1 interface

If you are using Typescript, you can use IEventLogClientV1 as a common interface across all client implementations. 
If you are using plain Golang, you shall not worry about IEventLogClientV1 interface. You can just expect that
all methods defined in this interface are implemented by all client classes.

``` golang
type IEventLogClientV1 interface {
	GetEvents(correlationId string, filter *cdata.FilterParams,
		paging *cdata.PagingParams) (page *SystemEventV1DataPage, err error)

	LogEvent(correlationId string, event *SystemEventV1) error
}
```

### <a name="operation1"></a> GetEvents(correlationId string, filter *cdata. FilterParams, paging *cdata. PagingParams) (page *SystemEventV1DataPage, err error)

Retrieves system events by specified criteria

**Arguments:** 

* correlationId: string - id that uniquely identifies transaction
* filter: object - filter parameters
  + search: string - (optional) search substring to find in source, type or message
  + type: string - (optional) type events
  + source: string - (optional) server where events occured
  + severity: number - (optional) severity of events
  + from_time: Date - (optional) start of the time range
  + to_time: Date - (optional) end of the time range
* paging: object - paging parameters
  + skip: int - (optional) start of page (default: 0)
  + take: int - (optional) page length (default: 100)
  + total: boolean - (optional) include total counter into paged result (default: false)
* return: (page, err) - returns param
  + err: Error - occured error or null for success
  + page: DataPage<SystemEventV1> - retrieved SystemEventV1 objects in paged format

### <a name="operation2"></a> LogEvent(correlationId string, event *SystemEventV1) error

Log system event

**Activities:** 

* correlationId: string - id that uniquely identifies transaction
* event: SystemEventV1 - system evemt to be logged
* return: 
  + err: Error - occured error or null for success

## <a name="client_http"></a> EventLogHttpClientV1 class

EventLogHttpClientV1 is a client that implements HTTP protocol

``` golang
type EventLogHttpClientV1 struct {
	*clients.CommandableHttpClient
}
func NewEventLogHttpClientV1() *EventLogHttpClientV1 
func (c *EventLogHttpClientV1) GetEvents(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*SystemEventV1DataPage, error) 
func (c *EventLogHttpClientV1) LogEvent(correlationId string, event *SystemEventV1) error 
```

**Constructor config properties:** 

* connection: object - HTTP transport configuration options
  + type: string - HTTP protocol - 'http' or 'https' (default is 'http')
  + host: string - IP address/hostname binding (default is '0.0.0.0')
  + port: number - HTTP port number

## <a name="client_direct"></a> EventLogDirectClientV1 class

EventLogDirectClientV1 is a client that calls controller directly from the same container.
It can be used in monolythic deployments when multiple microservices run in the same process.

``` golang
type EventLogDirectClientV1 struct {
	clients.DirectClient
	controller logic.IEventLogController
}
func NewEventLogDirectClientV1() *EventLogDirectClientV1 
func (c *EventLogDirectClientV1) SetReferences(references cref.IReferences) 
func toServerObject(value *SystemEventV1) *sdata1.SystemEventV1 
func fromServerObject(value *sdata1.SystemEventV1) *SystemEventV1 
func fromServerPage(value *sdata1.SystemEventV1DataPage) *SystemEventV1DataPage 
func (c *EventLogDirectClientV1) GetEvents(correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*SystemEventV1DataPage, error) 
func (c *EventLogDirectClientV1) LogEvent(correlationId string, event *SystemEventV1) error 
```

## <a name="client_null"></a> EventLogNullClientV1 class

EventLogNullClientV1 is a dummy client that mimics the real client but doesn't call a microservice. 
It can be useful in testing scenarios to cut dependencies on external microservices.

``` golang
type EventLogNullClientV1 struct {
}

func NewEventLogNullClientV1() *EventLogNullClientV1 {
	return &EventLogNullClientV1{}
}

func (c *EventLogNullClientV1) GetEvents(correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*SystemEventV1DataPage, error) {
	return NewEmptySystemEventV1DataPage(), nil
}

func (c *EventLogNullClientV1) LogEvent(correlationId string, event *SystemEventV1) error {
	return nil
}
```
