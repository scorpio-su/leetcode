package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	num := 0
	for _, q := range qs {
		ans0, p := q.ans1.one, q.para1
		num1 := twoSum2(p.nums, p.target)
		tes := true
		fmt.Printf("【input】:%v  【output】:%v  【needing】:%v  【scorce】:", p, num1, ans0)
		if len(num1) != len(ans0) {
			tes = false
		} else {
			for i := 0; i < len(ans0); i++ {
				if ans0[i] != num1[i] {
					tes = false
					break
				}
			}
		}
		if !tes {
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
