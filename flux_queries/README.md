```
from(bucket: "general")
  |> range(start: -10m)
  |> filter(fn: (r) => r["_measurement"] == "request_info" and r["_field"] =~ /uri|duration/)
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> filter(fn: (r) => r.uri == "/auth/login")
```