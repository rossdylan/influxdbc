package influxdbc

import "net/http"
import "encoding/json"
import "fmt"
import "bytes"
import "errors"

type InfluxDB struct {
	host     string
	database string
	username string
	password string
}

func NewInfluxDB(host string, database string, username string, password string) *InfluxDB {
	return &InfluxDB{host: host, database: database, username: username, password: password}
}

func (db *InfluxDB) SeriesURL() string {
	return fmt.Sprintf("http://%s/db/%s/series?u=%s&p=%s", db.host, db.database, db.username, db.password)
}

func (db *InfluxDB) QueryURL(query, timePrecision string) string {
	return fmt.Sprintf("http://%s/db/%s/series?u=%s&p=%s&q=query&time_precision=%s", db.host, db.database, db.username, db.password, query, timePrecision)
}

func (db *InfluxDB) WriteSeries(s []Series) error {
	url := db.SeriesURL()
	_, err := PostStruct(url, s)
	return err
}

func (db *InfluxDB) Query(query, tp string) ([]Series, error) {
	url := db.QueryURL(query, tp)
	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	if result.StatusCode != 200 {
		return nil, errors.New(buf.String())
	}
	series := make([]Series, 0)
	err = json.Unmarshal(buf.Bytes(), &series)
	if err != nil {
		return nil, err
	}
	return series, nil
}

func PostStruct(url string, reqStruct interface{}) (string, error) {
	marshalled, err := json.Marshal(reqStruct)
	marshalled = bytes.ToLower(marshalled)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(marshalled)
	result, err := http.Post(url, "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	result_buf := new(bytes.Buffer)
	result_buf.ReadFrom(result.Body)
	if result.StatusCode != 200 {
		return "", errors.New(result_buf.String())
	}
	return result_buf.String(), nil

}
