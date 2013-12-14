package influxdbc

import "testing"
import "fmt"

func Test_createURL(t *testing.T) {
	testDB := InfluxDB{"localhost", "testdb", "tester", "password"}
	url := testDB.SeriesURL()
	expectedURL := "http://localhost/testdb/series?u=tester&p=password"
	if url != "http://localhost/db/testdb/series?u=tester&p=password" {
		t.Errorf("'%s' does not match expected url, '%s'", url, expectedURL)
	}
}

func Test_NewSeries(t *testing.T) {
	columns := []string{"hello", "this", "is", "a", "test"}
	series := NewSeries("Test", columns[0], columns[1], columns[2], columns[3], columns[4])
	if series.Name != "Test" {
		t.Errorf("name: '%s' != 'Test'")
	}
	var failed bool
	failed = false
	for index, value := range series.Columns {
		if columns[index] != value {
			failed = true
		}
	}
	if failed {
		t.Errorf("Columns did not match: %v", series.Columns)
	}
}

func Test_SeriesWrite(t *testing.T) {
	database := InfluxDB{"localhost:8086", "testdb", "root", "root"}
	series := NewSeries("testseries", "col1", "col2")
	series.AddPoint("col1", "data", "col2 data")
	fmt.Println(series)
	database.WriteSeries([]Series{*series})
}

func Test_Query(t *testing.T) {
	database := InfluxDB{"localhost:8086", "testdb", "root", "root"}
	fmt.Println(database.Query("select * from testseries;", "s"))
}

func Test_GetClusterAdmins(t *testing.T) {
	testDB := InfluxDB{"localhost:8086", "testdb", "root", "root"}
	testDB.AddClusterAdmin("herp", "derp")
	admins, _ := testDB.GetClusterAdmins()
	fmt.Println(admins)

}
