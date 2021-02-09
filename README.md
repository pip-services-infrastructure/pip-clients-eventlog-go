# <img src="https://seekvectorlogo.com/wp-content/uploads/2018/05/national-oilwell-varco-nov-vector-logo.png" alt="NOV Logo" width="200"> <br/> Client Library for Event Log Microservice

This client library allows to use the [nov-max-system-eventlog](https://github.com/NationalOilwellVarco/max-system/service-eventlog-go) microservice to store and process system events.

Supported functionality:
* Commandable HTTP client
* Direct client to connect to the microservice in-process
* Memory client to be used as a mock in automated tests

<a name="links"></a> Quick links:

* Communication Protocols:
  - [HTTP Version 1](swagger/eventlog_v1.yaml)
=* [API Reference](https://NationalOilwellVarco.github.io/max-system/pages/client-eventlog-do/globals.html)
* [Change Log](CHANGELOG.md)

##  Contract

```golang

type SystemEventV1 struct {
	Id            string               `json:"id" bson:"_id"`
	Time          time.Time            `json:"time" bson:"time"`
	CorrelationId string               `json:"correlation_id" bson:"correlation_id"`
	Source        string               `json:"source" bson:"source"`
	Type          string               `json:"type" bson:"type"`
	Severity      int                `json:"severity" bson:"severity"`
	Message       string               `json:"message" bson:"message"`
	Details       cdata.StringValueMap `json:"details" bson:"details"`
}

// EventLogTypeV1
const Restart = "restart"
const Failure = "failure"
const Warning = "warning"
const Transaction = "transaction"
const Other = "other"

// EventLogSeverityV1
const Critical = 0
const Important = 500
const Informational = 1000

interface IEventLogV1 {
    getEvents(correlationId: string, filter: FilterParams, paging: PagingParams, 
        callback: (err: any, page: DataPage<SystemEventV1>) => void): void;
    
    logEvent(correlationId: string, event: SystemEventV1, 
        callback?: (err: any, event: SystemEventV1) => void): void;
}

```

## Get

Get the client library source from GitHub:
```bash
git clone git@github.com:NationalOilwellVarco/max-system.git
```

## Use

Inside your code get the reference to the client SDK
```golang
import (
	eventlog1 "github.com/NationalOilwellVarco/max-system/client-eventlog-go/version1"
)

var client *eventlog1.EventLogHttpClientV1
```

Define client configuration parameters that match configuration of the microservice external API
```golang
// Client configuration
httpConfig := cconf.NewConfigParamsFromTuples(
    "connection.protocol", "http",
    "connection.port", "3000",
    "connection.host", "localhost",
)

client = eventlog1.NewEventLogHttpClientV1()
client.Configure(httpConfig)
```

Instantiate the client and open connection to the microservice
```golang

// Connect to the microservice
err := client.Open("")
 if (err) {
        panic("Connection to the microservice failed");
    }
defer client.Close("")
// Work with the microservice

```

Now the client is ready to perform operations
```golang
// Log system event
event1 := &eventlog1.SystemEventV1{
    Type: "restart",
    Source: "server1",
    Message: "Restarted server",
}

err := client.LogEvent("", event1);
```

```golang
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

```    

## Develop

For development you shall install the following prerequisites:
* Golang 1.13+
* Visual Studio Code or another IDE of your choice
* Docker

In order to retrieve dependencies from the private repository
you shall set the following environment variable:
```bash
GOPRIVATE=github.com/NationalOilwellVarco/*
```

Install dependencies:
```bash
make install
```

Compile the code:
```bash
make build
```

Run automated tests:
```bash
make test
```

<!--
Generate GRPC protobuf stubs:
```bash
./protogen.ps1
```
-->

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized build and test as:
```bash
./build.ps1
./test.ps1
./clear.ps1
```

## Contacts

This client SDK was created and currently maintained by *Sergey Seroukhov* and *Michael Wright*.
