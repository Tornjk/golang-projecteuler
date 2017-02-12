package problems

import ( "fmt"; "math" )

func Solve3() {
	prim := []int{}
	limit := 600851475143;
	reduced := limit;
	max := int(math.Sqrt(float64(limit)))
	for i := 2; i < max; i++ {
		b := true
		for j := 0; j < len(prim); j++ {
			if(i % prim[j] == 0) {
				b = false
				break;
			}
		}

		if(b) {
			prim = append(prim, i)
			for reduced % i == 0 && reduced != 1 {
				reduced /= i
			}

			if(reduced == 1) {
				fmt.Println(i)
				return;
			}
		}
	}
}
