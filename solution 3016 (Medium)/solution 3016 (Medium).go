package main

import "sort"

/*
Повертає мінімальну кількість натискань на клавіши для набору слова `word`.

Всього є 8 клавіш, кожній можно назначити ряд літер в будь якому порядку.
Для набору першої літери потрібно натиснути клавішу один раз, для другої
два рази і так далі (як на старих телефонах):
	//  1  |  2  |  3
	//     | abc | def
	// ----------------
	//  4  |  5  |  6
	// ghi | jkl | mno
	// ----------------
	//  7  |  8  |  9
	// pqrs| tuv | wxyz

Приклад:
	minimumPushes("aabccddeffgghiiiijjjj") -> 23

	//  1  |  2  |  3
	//     | ie  | jh
	// ----------------
	//  4  |  5  |  6
	//  a  |  c  |  d
	// ----------------
	//  7  |  8  |  9
	//  f  |  g  |  b
*/
func minimumPushes(word string) int {
	table := make([]int, 26)

	for _, c := range word {
		table[int(c-'a')]++
	}

	sort.Ints(table)
	pushes := 0
	key := 0
	mult := 1

	for i := len(table) - 1; i >= 0; i-- {
		pushes += table[i] * mult
		key++

		if key == 8 {
			key = 0
			mult++
		}
	}
	return pushes
}

func main() {
	println(minimumPushes("aabccddeffgghiiiijjjj"))
}
