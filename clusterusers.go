package influxdbc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (db InfluxDB) AddClusterAdmin(name, password string) error {
	url := fmt.Sprintf("http://%s/cluster_admins?u=%s&p=%s", db.host, db.username, db.password)
	reqMap := map[string]string{
		"username": name,
		"password": password,
	}
	_, err := PostStruct(url, reqMap)
	return err
}

func (db InfluxDB) UpdateClusterAdmin(name, password string) error {
	url := fmt.Sprintf("http://%s/cluster_admins/%s?u=%s&p=%s", db.host, name, db.username, db.password)
	reqMap := map[string]string{
		"password": password,
	}
	_, err := PostStruct(url, reqMap)
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

func (db InfluxDB) GetClusterAdmins() (map[string]string, error) {
	url := fmt.Sprintf("http://%s/cluster_admins?u=%s&p=%s", db.host, db.username, db.password)
	result, err := http.Get(url)
	defer result.Body.Close()
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	admins := make(map[string]string)
	fmt.Println(admins)
	json.Unmarshal(buf.Bytes(), &admins)
	return admins, nil
}
