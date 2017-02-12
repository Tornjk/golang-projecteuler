package problems	

import "fmt"

func Solve2() {
	fib := []int{1, 1}
	for {
		i := len(fib)-1
		n := fib[i-1] + fib[i]
		fib = append(fib, n)

		if n > 4E6 {
			break
		}
	}

	sum := 0
	for i := 2; i < len(fib); i+=3{
		sum += fib[i]
	}

	fmt.Println(sum)
}
