package problems

import "fmt"

func Solve1() {
	sum := 0
	for i := 0; i < 1000; i += 3 {
		sum += i;
	}

	for i := 0; i < 1000; i += 5 {
		sum += i;
	}

	for i := 0; i < 1000; i += 15 {
		sum -= i
	}

	fmt.Println(sum)
}
