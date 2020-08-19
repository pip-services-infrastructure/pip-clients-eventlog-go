package version1

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type IEventLogClientV1 interface {
	GetEvents(correlationId string, filter *cdata.FilterParams,
		paging *cdata.PagingParams) (page *SystemEventV1DataPage, err error)

	LogEvent(correlationId string, event *SystemEventV1) error
}
