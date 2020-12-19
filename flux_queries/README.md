```
from(bucket: "general")
  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)
  |> filter(fn: (r) => r["_measurement"] == "request_info" and r["_field"] =~ /uri|duration/)
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> filter(fn: (r) => r.uri == "/auth/login")
  |> aggregateWindow(every: 30s, fn: count, createEmpty: true, column: "duration")
```

```
from(bucket: "general")
  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)
  |> filter(fn: (r) => r["_measurement"] == "request_info" and r["_field"] =~ /uri|duration|method/)
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> filter(fn: (r) => r.uri =~ /book*/ and r.method == "GET")
  |> aggregateWindow(every: 30s, fn: count, createEmpty: true, column: "duration")
```

```
from(bucket: "general")
  |> range(start: v.timeRangeStart, stop: v.timeRangeStop)
  |> filter(fn: (r) => r["_measurement"] == "request_info" and r["_field"] =~ /uri|duration|method/)
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> filter(fn: (r) => r.uri =~ /book*/ and r.method == "GET")
  |> aggregateWindow(every: 30s, fn: count, createEmpty: true, column: "duration")
```