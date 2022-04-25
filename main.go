package main

import "fmt"
import "strconv"

func main() {

	var cities = [4][4]int{
		{0, 2, 5, 7}, // a b c d
		{2, 0, 8, 3}, // b
		{5, 8, 0, 1}, // c
		{7, 3, 1, 0}} // d

	gg := permutation([]int{1, 2, 3})

	total := 0

	shortestDistance := -1
	shortestPaths := []string{}

	for _, elem := range gg {

		fmt.Println("route:", routesToStr(elem))

		lastCity := 0

		for _, city := range elem {

			total += cities[lastCity][city]

			lastCity = city
		}

		total += cities[lastCity][0]

		fmt.Println("total distance:", total)

		if shortestDistance == -1 || shortestDistance > total {

			shortestDistance = total
			shortestPaths = append(shortestPaths, routesToStr(elem))
		}

		total = 0

	}

	fmt.Println()
	fmt.Println()
	fmt.Println("shortestDistance:", shortestDistance)
	fmt.Println("shortestPaths:", shortestPaths)
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + 1 + start
	}
	return xs
}

func routesToStr(arr []int) string {

	result := "(0,"

	for _, o := range arr {
		result += strconv.Itoa(o)
		result += ","
	}

	result += "0)"

	return result

}
