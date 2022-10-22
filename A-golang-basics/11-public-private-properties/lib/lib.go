package lib

type Stats struct {
	data    []int
	len     int
	average float64
}

func MakeStats(data []int) Stats {
	temp := Stats{
		data: data,
	}
	temp.updateLen()
	temp.updateAvg()
	return temp
}

func (s Stats) GetLen() int {
	return s.len
}

// any methods or props with name started with lowercase is private, and cannot be accessed from outside this package/folder/file
func (s *Stats) updateLen() {
	s.len = len(s.data)
}

func (s Stats) GetAvg() float64 {
	return s.average
}

func (s *Stats) updateAvg() {
	temp := 0
	for _, e := range s.data {
		temp += e
	}
	temp /= s.len

	s.average = float64(temp)
}

func (s *Stats) Append(data int) {
	s.data = append(s.data, data)
	s.updateLen()
	s.updateAvg()
}
