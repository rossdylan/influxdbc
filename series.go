package influxdbc

type Series struct {
	Name       string
	Columns    []string
	DataPoints [][]string
}

func NewSeries(name string, cols ...string) *Series {
	s := new(Series)
	s.Name = name
	s.Columns = cols
	s.DataPoints = make([][]string, 0)
	return s
}

func (s *Series) AddPoint(point ...string) {
	s.DataPoints = append(s.DataPoints, point)
}
