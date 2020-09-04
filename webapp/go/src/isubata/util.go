package main

func uniq(a []int64) ([]int64, error) {
	ret := make([]int64, 0)
	m := make(map[int64]struct{}, 0)
	for _, el := range a {
		if _, ok := m[el]; !ok {
			m[el] = struct{}{}
			ret = append(ret, el)
		}
	}
	return ret, nil
}

func merge(a, b []int64) ([]int64, error) {
	merged := a[:]
	merged = append(merged, b...)

	ret, err := uniq(merged)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
