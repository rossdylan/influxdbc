package influxdbc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (db *InfluxDB) GetDatabaseUsers() (map[string]string, error) {
	url := fmt.Sprintf("http://%s/db/%s/users?u=%s&p=%s", db.host, db.database, db.username, db.password)
	result, err := http.Get(url)
	defer result.Body.Close()
	if err != nil {
		return nil, err
	}
	users := make(map[string]string)
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	json.Unmarshal(buf.Bytes(), &users)
	return users, nil
}

func (db *InfluxDB) AddDatabaseUser(name, password string) error {
	url := fmt.Sprintf("http://%s/db/%s/users?u=%s&p=%s", db.host, db.database, db.username, db.password)
	userInfo := map[string]string{
		"name":     name,
		"password": password,
	}
	_, err := PostStruct(url, userInfo)
	if err != nil {
		return err
	}
	return nil
}

func (db *InfluxDB) DeleteDatabaseUser(name string) error {
	url := fmt.Sprintf("http://%s/db/%s/users/%s?u=%s&p=%s", db.host, db.database, name, db.username, db.password)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	result, _ := http.DefaultClient.Do(req)
	defer result.Body.Close()
	return nil
}

func (db *InfluxDB) UpdateDatabaseUserPassword(name, password string) error {
	url := fmt.Sprintf("http://%s/db/%s/users/%s?u=%s&p=%s", db.host, db.database, name, db.username, db.password)
	reqMap := map[string]string{
		"password": password,
	}
	_, err := PostStruct(url, reqMap)
	return err
}

func (db *InfluxDB) UpdateDatabaseUserPrivileges(name string, admin bool) error {
	url := fmt.Sprintf("http://%s/db/%s/users/%s?u=%s&p=%s", db.host, db.database, name, db.username, db.password)
	reqMap := map[string]bool{
		"admin": admin,
	}
	_, err := PostStruct(url, reqMap)
	return err
}
