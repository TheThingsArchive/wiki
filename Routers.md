## NOTE

This document is neither final nor authoritative. 
It's currently only for the purpose of facilitating discussion. 

## Overview

The router is responsible for forwarding packets from gateways to the appropriate
handler (this has changed ?). The router needs to be configurable by application developers in order 
add/update/remove applications ~~and configure handlers~~. 

## Application

From the point of view of the router, an application consists of the following 
information: 

* AppID   -- Unique application id
* APIKey  -- Key for accessing API - hashed token (?) 
* Name    -- Human-readable name - display / lookup purposes
* Email   -- Contact address - resetting APIKey, etc. 
* Handler -- Address of handler for application
* Other things - timestamps, etc. 

### Ranges

An application can have any number of ranges. A range is specified as the first
3 bytes of an 4 byte integer, so each range may contain 256 node addresses. An 
application's nodes will be assigned unique addresses from the application's 
assigned ranges. 

## Storage

In a relational system, we'd need two tables: one having a row for each 
application and one mapping ranges to application ids. The database would 
need to be shared between all routers. 

## API

The router should be configuratable via an API. 

**Endpoints**

* app/create
* app/appID
* app/appID/range/request
* app/appID/range/create
* app/appID/range/delete/rangeID
* app/appID/range/list

* app/appID/handler
* app/appID/handler/create
* app/appID/handler/update
* app/appID/handler/delete

The odd one out here is the `range/request` which is used if an application 
wants to requests a specific range. This is useful if an application is using
more than one router. 

### app/create

Arguments:

```
{
	"name": "My app name",
    "email": "me@app.com",
    ???
}
```

Return:

```
{
    "app_id": 0x01234567,
    "api_key": "my secret api key"
}
```

