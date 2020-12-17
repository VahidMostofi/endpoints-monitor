# Endpoints Monitor
Monitors different endpoints using nginx and stores information about request/response in influxdb. The stored information would be available using an API and a dashboard.

## nginx-gateway

Acts as gateway for endpoints. Pass the requests to the backend as revers proxy server. The backend is specified by ```PROXY_PASS_URL``` environment variable. Also, sends information about each request to the telegraf agent using syslog at ```TELEGRAF_SYSLOG_SERVER```. The provided information are as follow using [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/).

-  ```duration=$request_time```
-  ```status=$status```
-  ```uri="$request_uri"```
-  ```method="$request_method"```
-  ```ust=$upstream_response_time```
-  ```usc=$upstream_connect_time```

Also, time is reported using ```$time_iso8601``` at the end of the line. This is not a valid format for [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/). But it will be changed at telegraf agent before sending to influxdb.

Specify the port which nginx listen to for incoming traffic using environment variable ```PORT```.


