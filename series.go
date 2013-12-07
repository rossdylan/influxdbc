package influxdbc

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
