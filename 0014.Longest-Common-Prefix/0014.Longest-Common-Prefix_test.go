package leetcode

import (
	"fmt"
	"testing"
)

type question14 struct {
	para14
	ans14
}

type para14 struct {
	lis []string
}

type ans14 struct {
	lis string
}

func Test_Problem14(t *testing.T) {

	qs := []question14{

		{
			para14{[]string{"flower", "flow", "flight"}},
			ans14{"fl"},
		},

		{
			para14{[]string{"dog", "racecar", "car"}},
			ans14{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 9------------------------\n")
	for _, q := range qs {
		ans0, p := q.ans14, q.para14
		fmt.Printf("【input】:%v  【output】:%v  【needing】:%v\n", p.lis, longestCommonPrefix(p.lis), ans0.lis)
	}
	fmt.Printf("\n\n")
}
