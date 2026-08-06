package main

/*
Повертає кількість повнозв'язкових підграфів в графі, який побудований
на основі ребер `edges` та має `n` нод (від 0 до n-1).

Повнозв'язковий підграф це граф, кожна нода якого пов'язана із
іншими його нодами.

	countCompleteComponents(9, [][]int{
		{0, 1}, {0, 2}, {1, 2},
		{3, 4},
		{6, 7}, {6, 8},
	}) -> 3

	//    (0)      (3)            (6)
	//    / \       |    (5)      / \
	//   /   \      |            /   \
	// (1)---(2)   (4)         (7)   (8)
	//
	//     +        +     +        Х
*/
func countCompleteComponents(n int, edges [][]int) int {
	// Detect nodes and it links

	nodes := make([][]int, n)

	for i := range edges {
		n1, n2 := edges[i][0], edges[i][1]

		nodes[n1] = append(nodes[n1], n2)
		nodes[n2] = append(nodes[n2], n1)
	}

	// Mark subgraphs and check rule

	marked := make([]bool, n)
	count := 0

	for i := range nodes {
		if marked[i] {
			continue
		}
		marked[i] = true

		queue := []int{}

		for _, node := range nodes[i] {
			queue = append(queue, node)
			marked[node] = true
		}

		isFull := true

		for j := 0; j < len(queue); j++ {
			n1 := queue[j]

			if len(nodes[n1]) != len(nodes[i]) {
				isFull = false
			}

			for _, n2 := range nodes[n1] {
				if !marked[n2] {
					queue = append(queue, n2)
					marked[n2] = true
					isFull = false
				}
			}
		}

		if isFull {
			count++
		}
	}

	return count
}

func main() {
	println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
}
