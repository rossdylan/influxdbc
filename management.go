package influxdbc

import (
	"fmt"
	"net/http"
)

type CreateDBReq struct {
	Name              string
	ReplicationFactor int
}

func (db InfluxDB) CreateDatabase(repFactor int) {
	url := fmt.Sprintf("http://%s/db?u=%s&p=%s", db.host, db.username, db.password)
	reqStruct := CreateDBReq{db.database, repFactor}
	PostStruct(url, reqStruct)
}

func (db InfluxDB) DeleteDatabase(database string) error {
	url := fmt.Sprintf("http://%s/db/%s?u=%s&p=%s", db.host, db.database, db.username, db.password)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	result, _ := http.DefaultClient.Do(req)
	defer result.Body.Close()
	return nil
}
