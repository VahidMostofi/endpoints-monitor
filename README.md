# Endpoints Monitoring System
Monitors different endpoints using nginx and stores information about request response time in influxdb. The stored information would be available using an API and a dashboard.

<p align="center">
  <img src="https://raw.githubusercontent.com/VahidMostofi/endpoints-monitor/dev-v0/Monitoring-Ednpoints.png" />
</p>

[How to use](##how-to-use)
<br/><br/>
[nginx-gateway](##nginx-gateway)
<br/><br/>
[telegraf-agent](##telegraf-agent)
<br/><br/>
[example](##example)

## How to use?
You definetly don't need all of these stuff as you may already have an nginx running as your API gateway. Just use the the telegraf agent wiht appropriate environemtn variables as a side-car in the same pod of nginx (assuming you are using kubernetes). It will also be scalable too. Then add these 2 lines to your nginx config file and enjoy your metrics on you influxdb (or anyother telegraf output, but the default config on this Docker image is for influxdb).

```
log_format influxdb 'request_info,provider=nginx duration=$request_time,status=$status,uri="$request_uri",method="$request_method",ust=$upstream_response_time,usc=$upstream_connect_time $time_iso8601';

access_log  syslog:server=${TELEGRAF_SYSLOG_SERVER},nohostname influxdb;
```

replace ```TELEGRAF_SYSLOG_SERVER``` with appropriate url.

refer to ```deployment/k8s```.


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

The metrics reported as syslog format by nginx-gateway. Telegraf has it's own [inputs.syslog](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/syslog) but it only supports syslog [RFC 5424](https://tools.ietf.org/html/rfc5424). Nginx syslog format is [RFC 3164](https://tools.ietf.org/html/rfc3164#section-4.1.1). There might be a way to convert these two easily but I ended up creating [syslog Parser](###syslog-parser) to convert syslog to [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/). The source code is available at ```telegraf-agent/get_logs.go```. It listens to port ```SYSLOG_SERVER_PORT``` environment variable and expects UDP packages in the format specified in [###sample-syslog-report]. It simply remove the first part of syslog message and changes the format fo the timestamp.

Telegraf ```inputs.execd``` runs this code and uses the stdout to send metrics to influxdb. It also reports ```inputs.nginx``` to influxdb gathered from ```NGINX_DEFAULT_STATS_URL``` environment variables.

The only metrics which are reported are the ones coming from [nginx-gateway](##nginx-gateway) and [/nginx_status](https://www.nginx.com/blog/monitoring-nginx/). If you need to add more metrics edit the ```telegraf.conf``` file.

### InfluxDB environment variables

This repo is using influxdb_v2 and uses these environment variables to authenticate and connect to influxdb.

- ```INFLUXDB_URL```
- ```INFLUXDB_BUCKET```
- ```INFLUXDB_TOKEN```
- ```INFLUXDB_ORG```

For more info about the [influxdb](https://docs.influxdata.com/influxdb/v2.0/).

### syslog Parser

## example
