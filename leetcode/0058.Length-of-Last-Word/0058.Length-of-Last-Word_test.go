package leetcode

import (
	"fmt"
	"testing"
)

type question58 struct {
	pare58
	ans58
}

type pare58 struct {
	one string
}

type ans58 struct {
	one int
}

func Test_Problem58(t *testing.T) {

	qs := []question58{

		{
			pare58{"Hello World"},
			ans58{5},
		},

		{
			pare58{"a "},
			ans58{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 58------------------------\n")
	num := 0
	for _, q := range qs {
		ans0, p := q.ans58.one, q.pare58.one
		q1 := lengthOfLastWord(p)
		fmt.Printf("【input】:%v  【output】:%v  【needing】:%v  【scorce】:", p, q1, ans0)
		if q1 != ans0 {
			fmt.Println("fail")
			num++
		} else {
			fmt.Println("pass")
		}
	}
	if num > 0 {
		fmt.Printf("have %d errors answer", num)
	} else {
		fmt.Printf("all pass")
	}
	fmt.Printf("\n\n")

}
