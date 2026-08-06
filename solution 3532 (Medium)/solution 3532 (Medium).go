package main

type Group struct {
	Start, End int
}

/*
Умови:
  - n - кількість нод
  - nums - різниця `nums[i]-nums[i-1]`, це дистанція між нодами `i` та `i-1`
  - maxDiff - максимальна дистанція для зв'язку між нодами

Функція повертає відповіді на `queries`, які запитують
чи є зв'язок між нодами `queries[i][0]` та `queries[i][1]`

	pathExistenceQueries(
		4,
		[]int{2, 5, 6, 8},
		2,
		[][]int{ {0, 1}, {0, 2}, {1, 3}, {2, 3} },
	) ->       [  false,  false,   true,   true ]

	// пари нод:    0-1  |  1-2  |  2-3
	// дистанції:  5-2=3 | 6-5=1 | 8-6=2
	// зв'язок:    false | true  | true

	//   (2)
	//   / \     (0)
	// (1) (3)
*/
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	nodes := make([]*Group, n)
	group := &Group{}
	nodes[0] = group

	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			group = &Group{i, i}
		}

		group.End = i
		nodes[i] = group
	}

	ans := make([]bool, len(queries))

	for i := range queries {
		n1, n2 := queries[i][0], queries[i][1]
		group = nodes[n1]

		if group.Start <= n2 && n2 <= group.End {
			ans[i] = true
		}
	}

	return ans
}

func main() {
	ans := pathExistenceQueries(4, []int{2, 5, 6, 8}, 2, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}})

	for i := range ans {
		print(ans[i], " ")
	}
}
