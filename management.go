package influxdbc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateDBReq struct {
	name              string
	replicationFactor int
}

type ClusterAdmin struct {
	name     string
	password string
}

type ClusterAdminUpdate struct {
	password string
}

func (db InfluxDB) CreateDatabase(database string, repFactor int) {
	url := fmt.Sprintf("http://%s/db?u=%s&p=%s", db.host, db.username, db.password)
	reqStruct := CreateDBReq{database, repFactor}
	PostStruct(url, reqStruct)
}

func (db InfluxDB) DeleteDatabase(database string) {
	url := fmt.Sprintf("http://%s/db/%s?u=%s&p=%s", db.host, database, db.username, db.password)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	result, _ := http.DefaultClient.Do(req)
	defer result.Body.Close()
}

func (db InfluxDB) AddClusterAdmin(name, password string) {
	url := fmt.Sprintf("http://%s/cluster_admins?u=%s&p=%s", db.host, db.username, db.password)
	reqStruct := ClusterAdmin{name, password}
	PostStruct(url, reqStruct)
}

func (db InfluxDB) UpdateClusterAdmin(name, password string) {
	url := fmt.Sprintf("http://%s/cluster_admins/%s?u=%s&p=%s", db.host, name, db.username, db.password)
	reqStruct := ClusterAdminUpdate{password}
	PostStruct(url, reqStruct)
}

func (db InfluxDB) DeleteClusterAdmin(name string) {
	url := fmt.Sprintf("http://%s/cluster_admins/%s?u=%s&p=%s", db.host, name, db.username, db.password)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	result, _ := http.DefaultClient.Do(req)
	defer result.Body.Close()
}

func (db InfluxDB) GetClusterAdmins() []ClusterAdmin {
	url := fmt.Sprintf("http://%s/cluster_admins?u=%s&p=%s", db.host, db.username, db.password)
	result, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer result.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	admins := make([]ClusterAdmin, 0)
	json.Unmarshal(buf.Bytes(), &admins)
	return admins
}
