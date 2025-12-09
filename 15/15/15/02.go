package main

import (
	"fmt"
	"strings"
)

func generateCombinations(n, k int) [][]int {
	var res [][]int
	curr := make([]int, k)
	generate(n, 0, curr, &res)
	return res
}

func generate(remaining, i int, curr []int, res *[][]int) {
	if i == len(curr)-1 {
		curr[i] = remaining
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	for j := 0; j <= remaining; j++ {
		curr[i] = j
		generate(remaining-j, i+1, curr, res)
	}
}

type Properties struct {
	Capacity   int64
	Durability int64
	Flavor     int64
	Texture    int64
	Calories   int64
}

func main() {
	var max int64
	input := `Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1`
	lines := strings.Split(input, "\n")
	properties := make(map[string]Properties)

	for _, l := range lines {
		var n string
		var c, d, f, t, cal int64
		fmt.Sscanf(l, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &n, &c, &d, &f, &t, &cal)
		n = strings.TrimSuffix(n, ":")
		properties[n] = Properties{Capacity: c, Durability: d, Flavor: f, Texture: t, Calories: cal}
	}
	combinations := generateCombinations(100, 4)

	for _, c := range combinations {
		fmt.Println(c)
		amount := make(map[string]int)
		amount["Frosting"] = c[0]
		amount["Candy"] = c[1]
		amount["Butterscotch"] = c[2]
		amount["Sugar"] = c[3]

		var totalCapacity, totalDurability, totalFlavor, totalTexture, totalCalories int64

		for ing, amt := range amount {
			totalCapacity += properties[ing].Capacity * int64(amt)
			totalDurability += properties[ing].Durability * int64(amt)
			totalFlavor += properties[ing].Flavor * int64(amt)
			totalTexture += properties[ing].Texture * int64(amt)
			totalCalories += properties[ing].Calories * int64(amt)
		}

		if totalCapacity <= 0 {
			totalCapacity = 0
		}

		if totalDurability <= 0 {
			totalDurability = 0
		}

		if totalFlavor <= 0 {
			totalFlavor = 0
		}

		if totalTexture <= 0 {
			totalTexture = 0
		}

		if totalCalories <= 0 {
			totalCalories = 0
		}

		score := totalCapacity * totalDurability * totalFlavor * totalTexture

		if score > max && totalCalories == 500 {
			max = score
		}
	}

	fmt.Println(max)
}
