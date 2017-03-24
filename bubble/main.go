package main

var toBeSorted [10]int = [10]int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}

func bubbleSort(input [10]int) {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
}

func main() {
	for c := 0; c < 10000000; c++ {
		bubbleSort(toBeSorted)
	}
}
