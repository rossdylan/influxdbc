InfluxDB Client for Golang
==========================

Golang bindings for the InfluxDB http API

``` go
import "github.com/rossdylan/influxdbc"
func main() {
    database := influxdbc.NewInfluxDB("localhost:8086", "testdb", "username", "password")
    series := influxdbc.NewSeries("Name", "Col1", "Col2")
    series.AddPoint("Col1 data", "Col2 data")
    err := database.WriteSeries([]influxdbc.Series{*series})
}
```
