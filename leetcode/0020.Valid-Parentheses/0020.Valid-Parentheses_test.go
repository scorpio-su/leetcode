package leetcode

import (
	"fmt"
	"testing"
)

type question20 struct {
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func Test_Problem20(t *testing.T) {

	qs := []question20{

		{
			para20{"()[]{}"},
			ans20{true},
		},
		{
			para20{"(]"},
			ans20{false},
		},
		{
			para20{"({[]})"},
			ans20{true},
		},
		{
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			para20{"(())]]"},
			ans20{false},
		},
		{
			para20{""},
			ans20{true},
		},
		{
			para20{"["},
			ans20{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")
	num := 0
	for _, q := range qs {
		ans0, p := q.ans20.one, q.para20
		q1 := isValid(p.one)
		fmt.Printf("【input】:%v  【output】:%v  【needing】:%v\n", p, q1, ans0)
		if q1 != ans0 {
			fmt.Println("fail")
			num++
		} else {
			fmt.Println("pass")
		}
		if num > 0 {
			fmt.Printf("have %d errors answer", num)
		} else {
			fmt.Printf("all pass")
		}
		fmt.Printf("\n\n")
	}
}
