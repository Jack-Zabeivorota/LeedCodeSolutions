package main

type Node struct {
	Links, SelfLinks []int
}

/*
Повертає коректно працюючі методи.
  - n - кількість методів пронумерованих від 0 до n-1;
  - k - номер метода, що виконався з помилкою;
  - invocs - виклики методів, де invocs[i] це [Ai, Bi], тобто метод A викликає метод B.

Умови фільтрації:
  - якщо метод виконався з помилкою, то весь ланцю викликаних ним методів
    (напряму або побічно) також ставиться під питання, формуючи підозрілу групу;
  - якщо коректно працюючий метод викликав один із методів підозрілої групи,
    то група перестає бути підозрілою.

Приклад 1:

	remainingMethods(4, 2, [][]int{{1, 0}, {2, 0}, {0, 3}}) -> [0, 1, 2, 3]

	//     3
	//     ↑
	// 1 → 0 ← 2

	Підозріла група: 2 -> 0 -> 3
	Але коректно працюючий метод `1` викликає метод `0`, що входить в групу,
	тому всі методи працюють коректно.

Приклад 2:

	remainingMethods(4, 2, [][]int{{2, 0}, {0, 3}}) -> [1]

	//     3
	//     ↑
	// 1   0 ← 2

	Жоден коректно працюючий метод не викликає метод підозрілої групи,
	тому група видаляється.
*/
func remainingMethods(n, k int, invocs [][]int) []int {
	// Prepare nodes

	nodes := make([]Node, n)

	for i := range invocs {
		n1, n2 := invocs[i][0], invocs[i][1]

		nodes[n1].Links = append(nodes[n1].Links, n2)
		nodes[n2].SelfLinks = append(nodes[n2].SelfLinks, n1)
	}

	// Find group

	group := map[int]bool{k: true}
	queue := []int{k}

	for i := 0; i < len(queue); i++ {
		n1 := queue[i]

		for _, n2 := range nodes[n1].Links {
			if !group[n2] {
				group[n2] = true
				queue = append(queue, n2)
			}
		}
	}

	// Find out links to group

	isFilter := true

A:
	for n1 := range group {
		for _, n2 := range nodes[n1].SelfLinks {
			if !group[n2] {
				isFilter = false
				break A
			}
		}
	}

	// Filter nodes

	if !isFilter {
		ans := make([]int, n)

		for node := 0; node < n; node++ {
			ans[node] = node
		}
		return ans
	}
	ans := make([]int, 0, n-len(group))

	for node := 0; node < n; node++ {
		if !group[node] {
			ans = append(ans, node)
		}
	}
	return ans
}

func main() {
	ans := remainingMethods(3, 2, [][]int{{1, 0}, {2, 0}})

	for _, node := range ans {
		print(node, " ")
	}
	println()
}
