package influxdbc

import "testing"

func Test_createURL(t *testing.T) {
	testDB := InfluxDB{"localhost", "tester", "password", "testdb"}
	url := testDB.PostURL()
	expectedURL := "http://localhost/testdb/series?u=tester&p=password"
	if url != "http://localhost/db/testdb/series?u=tester&p=password" {
		t.Errorf("'%s' does not match expected url, '%s'", url, expectedURL)
	}
}

func Test_NewSeries(t *testing.T) {
	columns := []string{"hello", "this", "is", "a", "test"}
	series := NewSeries("Test", columns[0], columns[1], columns[2], columns[3], columns[4])
	if series.name != "Test" {
		t.Errorf("name: '%s' != 'Test'")
	}
	var failed bool
	failed = false
	for index, value := range series.columns {
		if columns[index] != value {
			failed = true
		}
	}
	if failed {
		t.Errorf("Columns did not match: %v", series.columns)
	}
}
