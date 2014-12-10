InfluxDB Client for Golang
==========================

Golang bindings for the InfluxDB http API

``` go
import "github.com/rossdylan/influxdbc"
func main() {
    database := influxdbc.InfluxDB{"localhost:8083", "testdb", "username", "password"}
    series := influxdbc.NewSeries{"Col1", "Col2"}
    series.AddPoint("Col1 data", "Col2 data")
    err := database.WriteSeries("DB")
}
```
