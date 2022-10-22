package main

import "fmt"

type Stats struct {
	Data    []int
	Len     int
	Average float64
}

//      V ==> pointer method use a pointer as a receiver argument
func (s *Stats) setLen() {
	s.Len = len(s.Data)
}

func (s *Stats) setAVG() {
	temp := 0
	for _, e := range s.Data {
		temp += e
	}
	temp /= s.Len

	s.Average = float64(temp)
}

func (s *Stats) updateProp() {
	s.setLen()
	s.setAVG()
}

func (s *Stats) append(data int) {
	s.Data = append(s.Data, data)
	s.updateProp()
}

func main() {
	data1 := Stats{
		Data: []int{5, 1},
	}
	// with pointer method, we can change a property of an object, without asignment operator
	data1.updateProp()
	fmt.Println(data1)

	data1.append(6)
	fmt.Println(data1)
}
