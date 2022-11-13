package main

func CulcIntersection(l1, l2 []int) []int {
	s := make(map[int]struct{}, len(l1))

	for _, data := range l1 {
		s[data] = struct{}{}
	}

	r := make([]int, 0, len(l1))

	for _, data := range l2 {
		if _, ok := s[data]; !ok {
			continue
		}

		r = append(r, data)
	}
	return r
}

func CulcUnion(l1, l2 []int) []int {
	s := make(map[int]struct{}, len(l1))

	for _, data := range l1 {
		s[data] = struct{}{}
	}

	for _, data := range l2 {
		s[data] = struct{}{}
	}

	r := make([]int, 0, len(s))

	for key, _ := range s {
		r = append(r, key)
	}

	return r
}
