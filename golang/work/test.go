package work

import (
	"fmt"
)

//Given the following array
//[[a1,a2,a3],[b1,b2,b3,b4],[c1],[d1,d2],[e1]]

//Create a function that returns a transpose and flattened version of the array
//[a1,b1,c1,d1,e1,a2,b2,d2,a3,b3,b4]

func TransposeAndFlattenMatrix() {
	var (
		solutionArr = []string{}
		maxIndex    = 0
		maxLen      = 0
		counter     = 0
	)
	arr := [][]string{
		{"a1", "a2", "a3"},
		{"b1", "b2", "b3", "b4"},
		{"c1"},
		{"d1", "d2"},
		{"e1"},
	}

	// 0,0 1,0 2,0 3,0
	// Get max length in both axis
	for _, e := range arr {
		if len(e)-1 > maxIndex {
			maxIndex = len(e) - 1
		}
		for _, _ = range e {
			maxLen++
		}
	}

	for len(solutionArr) < maxLen {
		for i := 0; i < len(arr); i++ {
			if len(arr[i]) > counter {
				solutionArr = append(solutionArr, arr[i][counter])
			}
		}
		if counter >= maxIndex {
			counter = 0
		} else {
			counter++
		}
	}

	fmt.Println(solutionArr)
}
