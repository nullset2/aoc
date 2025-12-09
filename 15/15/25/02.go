package main

import "fmt"

func main() {
	numbers := make(map[string]int, 0)
	diagonalSize := 1
	coords := make([]string, 0)

	for {
		i := 1
		j := diagonalSize - i
		for i < diagonalSize && j >= 0 {
			coords = append(coords, fmt.Sprintf("%v,%v", j, i))
			if j == 3010 && i == 3019 {
				break
			}
			i++
			j--
		}

		if j == 3010 && i == 3019 {
			break
		}

		diagonalSize++
	}

	numbers["1,1"] = 20151125

	for i := 1; i < len(coords); i++ {
		next := (numbers[coords[i-1]] * 252533) % 33554393
		numbers[coords[i]] = next
	}

	fmt.Println(numbers["3010,3019"])

}
