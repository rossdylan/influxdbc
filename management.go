package influxdbc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateDBReq struct {
	Name              string
	ReplicationFactor int
}

type ClusterAdmin struct {
	Username string
	Password string
}

type ClusterAdminUpdate struct {
	Password string
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

func (db InfluxDB) AddClusterAdmin(name, password string) error {
	url := fmt.Sprintf("http://%s/cluster_admins?u=%s&p=%s", db.host, db.username, db.password)
	reqStruct := ClusterAdmin{name, password}
	_, err := PostStruct(url, reqStruct)
	return err
}

func (db InfluxDB) UpdateClusterAdmin(name, password string) error {
	url := fmt.Sprintf("http://%s/cluster_admins/%s?u=%s&p=%s", db.host, name, db.username, db.password)
	reqStruct := ClusterAdminUpdate{password}
	_, err := PostStruct(url, reqStruct)
	return err
}

func (db InfluxDB) DeleteClusterAdmin(name string) error {
	url := fmt.Sprintf("http://%s/cluster_admins/%s?u=%s&p=%s", db.host, name, db.username, db.password)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	result, _ := http.DefaultClient.Do(req)
	defer result.Body.Close()
	return nil
}

func (db InfluxDB) GetClusterAdmins() ([]ClusterAdmin, error) {
	url := fmt.Sprintf("http://%s/cluster_admins?u=%s&p=%s", db.host, db.username, db.password)
	result, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	admins := make([]ClusterAdmin, 0)
	fmt.Println(admins)
	json.Unmarshal(buf.Bytes(), &admins)
	return admins, nil
}
