package main

import (
	"fmt"
	"github.com/testisnullus/data-structs-and-algorithms/algos"
)

func main() {
	var sortedArr = []int{1, 5, 12, 23, 46, 87, 101, 187, 291, 332, 341, 396, 401, 502, 505, 555, 591, 600}

	fmt.Println(algos.BinarySearch(sortedArr, 101))
	fmt.Println(algos.BS(sortedArr, 101))
}
