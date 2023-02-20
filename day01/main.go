package main

func Part1() int {
	max := 0
	sum := 0
	for _, calories := range Data {
		if calories == 0 {
			if sum > max {
				max = sum
			}
			sum = 0
		}

		sum += calories
	}

	return max
}

type Stack struct {
	items [3]int
}

func (s *Stack) Push(data int) {
	s.items[2], s.items[1], s.items[0] = s.items[1], s.items[0], data

}
func (s *Stack) Total() int {
	return s.items[0] + s.items[1] + s.items[2]
}

func Part2() int {
	var max Stack

	sum := 0
	for _, calories := range Data {
		if calories == 0 {
			if sum > max.items[0] {
				max.Push(sum)
			}
			sum = 0
		}

		sum += calories
	}

	return max.Total()
}
