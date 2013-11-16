package influxdbc

import "net/http"
import "encoding/json"
import "fmt"
import "bytes"

type Series struct {
	name       string
	columns    []string
	dataPoints [][]string
}

func NewSeries(name string, cols ...string) *Series {
	s := new(Series)
	s.name = name
	s.columns = cols
	s.dataPoints = make([][]string, 0)
	return s
}

func (s *Series) AddPoint(point []string) {
	s.dataPoints = append(s.dataPoints, point)
}

type InfluxDB struct {
	host     string
	username string
	password string
	database string
}

func (db InfluxDB) PostURL() string {
	return fmt.Sprintf("http://%s/db/%s/series?u=%s&p=%s", db.host, db.database, db.username, db.password)
}

func (db InfluxDB) WriteSeries(s []Series) {
	marshalled, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(marshalled)
	url := db.PostURL()
	result, _ := http.Post(url, "application/json", buf)
	defer result.Body.Close()
}
