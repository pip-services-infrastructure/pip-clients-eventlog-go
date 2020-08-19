let os = require('os');

import { ConfigParams } from 'pip-services3-commons-node';
import { IReferences } from 'pip-services3-commons-node';
import { FilterParams } from 'pip-services3-commons-node';
import { PagingParams } from 'pip-services3-commons-node';
import { DataPage } from 'pip-services3-commons-node';
import { CommandableHttpClient } from 'pip-services3-rpc-node';

import { SystemEventV1 } from './SystemEventV1';
import { IEventLogClientV1 } from './IEventLogClientV1';

export class EventLogHttpClientV1 extends CommandableHttpClient implements IEventLogClientV1 {

    constructor(config?: any) {
        super('v1/eventlog');

        if (config != null)
            this.configure(ConfigParams.fromValue(config));
    }
        
    public getEvents(correlationId: string, filter: FilterParams, paging: PagingParams,
        callback: (err: any, page: DataPage<SystemEventV1>) => void) {
        this.callCommand(
            'get_events',
            correlationId,
            {
                filter: filter,
                paging: paging
            },
            callback
        );
    }

    public logEvent(correlationId: string, event: SystemEventV1,
        callback?: (err: any, event: SystemEventV1) => void) {

        event.time = event.time || new Date();
        event.source = event.source || os.hostname(); 

        this.callCommand(
            'log_event',
            correlationId,
            {
                event: event
            }, 
            callback
        );
    }

}
