package golang

// VisitAllPlacesMinimumTime This implementation does not work for array where len(array) > 1000.
// Extremily underperforming. Should not be used.
func VisitAllPlacesMinimumTime(A []int) int {
	var (
		visitTracker = make(map[int]bool)
		res          = 0
	)

	for _, e := range A {
		visitTracker[e] = false
	}

	for i, _ := range A {
		for j := i + 1; j <= len(A)-1; j++ {
			tempArr := A[i:j]
			for _, e := range tempArr {
				visitTracker[e] = true
			}
			if checkVisitedAll(visitTracker) && (len(tempArr) < res || res == 0) {
				res = len(visitTracker)
			}
			resetMap(visitTracker)
		}
	}
	return res
}

func resetMap(tracker map[int]bool) {
	for k, _ := range tracker {
		tracker[k] = false
	}
}

func checkVisitedAll(tracker map[int]bool) bool {
	all := true
	for _, ok := range tracker {
		if !ok {
			all = false
		}
	}
	return all
}

func VisitAllPlacesMinimumTimeV2(A []int) int {
	needed := make(map[int]int)
	for _, e := range A {
		needed[e] = 1
	}

	missing := len(needed)

	curI := 0
	resultI := 0
	resultJ := 0

	for curJ, num := range A {
		curJ++
		if needed[num] > 0 {
			missing--
		}
		needed[num]--

		if missing == 0 {
			for curI < curJ && needed[A[curI]] < 0 {
				needed[A[curI]]++
				curI++
			}

			if resultJ == 0 || curJ-curI <= resultJ-resultI {
				resultI, resultJ = curI, curJ
			}
		}
	}
	return resultJ - resultI

}
