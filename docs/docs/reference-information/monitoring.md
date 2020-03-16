---
id: monitoring
title: Monitoring
---

## Health checking

When you are running an inway or an outway you firstly want to make sure it is running and secondly you want to know if the in/outway is ready to receive requests.
Both the inway and the outway provide endpoints which allow you to check just that. 

### Liveness

The liveness endpoint is reachable on `health/live` and will return a http status `200 OK` when the in/outway has started.

### Readiness 

The readiness endpoint is reachable on `health/ready` and will return http status `200 OK` when the in/outway is ready to receive requests otherwise a http status `503 service unavailable` will be returned.  
The inway is ready when the service configuration has been loaded succefuly. 
The outway is ready when the service list has been retrieved from the directory.


### Configuration

By default the endpoints are reachable on port `8081`. If you want to change this port you can do so by setting the `monitoring-address` flag when starting the in/outway.

