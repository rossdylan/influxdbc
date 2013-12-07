package influxdbc

import "net/http"
import "encoding/json"
import "fmt"
import "bytes"

type InfluxDB struct {
	host     string
	username string
	password string
}

func (db InfluxDB) SeriesURL(database string) string {
	return fmt.Sprintf("http://%s/db/%s/series?u=%s&p=%s", db.host, database, db.username, db.password)
}

func (db InfluxDB) QueryURL(database, query, timePrecision string) string {
	return fmt.Sprintf("http://%s/db/%s/series?u=%s&p=%s&q=query&time_precision=%s", db.host, database, db.username, db.password, query, timePrecision)
}

func (db InfluxDB) WriteSeries(database string, s []Series) {
	url := db.SeriesURL(database)
	PostStruct(url, s)
}

func (db InfluxDB) Query(database, query, tp string) []Series {
	url := db.QueryURL(database, query, tp)
	result, _ := http.Get(url)
	defer result.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	series := make([]Series, 0)
	err := json.Unmarshal(buf.Bytes(), &series)
	if err != nil {
		panic(err)
	}
	return series
}

func PostStruct(url string, reqStruct interface{}) {
	marshalled, err := json.Marshal(reqStruct)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(marshalled)
	result, _ := http.Post(url, "application/json", buf)
	defer result.Body.Close()
}
