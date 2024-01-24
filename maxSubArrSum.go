package main

import "fmt"

func main() {
	ia := []int{-2, 3, 2, -1, 5, -2, 3}

	gm, cm := ia[0], ia[0]

	for i := 1; i < len(ia); i++ {
		cm = max(ia[i], cm+ia[i])
		if cm > gm {
			gm = cm
		}
	}

	fmt.Println(gm)
}
