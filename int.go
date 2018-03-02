package wtf

// IntsUnion return a union b
func IntsUnion(a, b []int) []int {
	want := []int{}
	want = append(want, a...)
	want = append(want, b...)
	return want
}

// IntsDiff return a diff from b
func IntsDiff(a, b []int) []int {
	want := []int{}
	for _, ae := range a {
		have := false
		for _, be := range b {
			if be == ae {
				have = true
				break
			}
		}
		if !have {
			want = append(want, ae)
		}
	}
	return want
}

// IntsIntersect return a intersect b
func IntsIntersect(a, b []int) []int {
	want := []int{}
	for _, ae := range a {
		have := false
		for _, be := range b {
			if be == ae {
				have = true
				break
			}
		}
		if have {
			want = append(want, ae)
		}
	}
	return want
}
