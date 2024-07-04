package main

import "fmt"

func protectedChickens(roof int, chickens []int) int {
	maxprotect := 0
	// Pop value decrese Loop
	tmpchicken := chickens
	for _, min := range chickens {
		max := min + roof
		// Count  within the range [min, max]
		count := 0
		for _, n := range tmpchicken {
			if n >= min && n < max {
				count++
			} else if n >= max {
				break
			} else if count == roof {
				break
			}
		}
		if count > maxprotect {
			maxprotect = count
		}
		// POP Value
		tmpchicken = chickens[1:]

	}

	return maxprotect
}

func main() {
	chickens := []int{2, 5, 10, 12, 15}
	roof := 5
	result := protectedChickens(roof, chickens)
	fmt.Printf("roof %d chickens %v  ==> Output %d\n", roof, chickens, result)

}
