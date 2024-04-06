package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n < cap(slice) { // reallocate if needed
		newSlice := make([]byte, (n+1)*2) // (n + 1) in case n == 0
		copy(newSlice, slice)
		slice = newSlice
	}

	for i := m; i < n; i++ {
		slice = append(slice, data[i-m])
	}

	// or we can do this
	// slice = append(slice, data...)

	return slice
}
