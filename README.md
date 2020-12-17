# Endpoints Monitoring System
Monitors different endpoints using nginx and stores information about request response time in influxdb. The stored information would be available using an API and a dashboard.

## nginx-gateway

Acts as gateway for endpoints. Pass the requests to the backend as revers proxy server. The backend is specified by ```PROXY_PASS_URL``` environment variable. Also, sends information about each request to the telegraf agent using syslog at ```TELEGRAF_SYSLOG_SERVER```. The provided information are as follow using [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/).

-  ```duration=$request_time```
-  ```status=$status```
-  ```uri="$request_uri"```
-  ```method="$request_method"```
-  ```ust=$upstream_response_time```
-  ```usc=$upstream_connect_time```

#### timestamp
Also, time is reported using ```$time_iso8601``` at the end of the line. This is not a valid format for [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/). But it will be changed at telegraf-agent [syslog Parser](###syslog-parser) before sending to influxdb.

#### nginx port
Specify the port which nginx listen to for incoming traffic using environment variable ```PORT```.

#### sample syslog report
This is what [syslog Parser](###syslog-parser) at telegraf-agent gets:

```
<190>Dec 16 04:21:13 nginx: request_info,provider=nginx duration=0.004,status=404,uri=/auth/login,method=GET,ust=0.003,usc=0.003 2020-12-16T23:27:37+00:00
```

and this is what it returns: (a valid [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/) example)

```
request_info,provider=nginx duration=0.004,status=404,uri=/auth/login,method=GET,ust=0.003,usc=0.003 1608161257000000000
```



## telegraf-agent

The only metrics which are reported are the ones coming from [nginx-gateway](##nginx-gateway) and [/nginx_status](https://www.nginx.com/blog/monitoring-nginx/). If you need to add more metrics edit the ```telegraf.conf``` file.


### syslog Parser
