package main

func getIntersectionOfSets(l1, l2 []int) []int {
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

func getUnionOfSets(l1, l2 []int) []int {
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

func getDifferenceOfSets(l1, l2 []int) []int {
	s := make(map[int]struct{}, len(l1))

	for _, data := range l2 {
		s[data] = struct{}{}
	}

	r := make([]int, 0, len(l2))

	for _, data := range l1 {
		if _, ok := s[data]; ok {
			continue
		}
		r = append(r, data)
	}
	return r
}

func getSymmetricDifferenceOfSets(l1, l2 []int) []int {
	u := getUnionOfSets(l1, l2)
	s := getIntersectionOfSets(l1, l2)

	smap := make(map[int]struct{}, len(s))

	for _, data := range s {
		smap[data] = struct{}{}
	}

	r := make([]int, 0, len(u))

	for _, data := range u {
		if _, ok := smap[data]; ok {
			continue
		}
		r = append(r, data)
	}

	return r
}
