package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	para7
	ans7
}

// para 是参数
// one 代表第一个参数
type para7 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans7 struct {
	one int
}

func Test_Problem7(t *testing.T) {

	qs := []question7{

		{
			para7{321},
			ans7{123},
		},

		{
			para7{-123},
			ans7{-321},
		},

		{
			para7{120},
			ans7{21},
		},

		{
			para7{1534236469},
			ans7{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 7------------------------\n")
	num := 0
	for _, q := range qs {
		ans0, p := q.ans7.one, q.para7.one
		q1 := reverse7(p)
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
