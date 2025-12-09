package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Transformation struct {
	from string
	to   string
}

// ReverseTransformations reverses the transformation rules (to -> from)
func ReverseTransformations(transformations []Transformation) []Transformation {
	reversed := make([]Transformation, len(transformations))
	for i, t := range transformations {
		reversed[i] = Transformation{from: t.to, to: t.from}
	}
	return reversed
}

// FindAllReplacements returns all possible strings by replacing one occurrence of 'from' with 'to'
func FindAllReplacements(input, from, to string) []string {
	if from == "" {
		return []string{input}
	}

	var results []string
	searchStart := 0

	for {
		idx := strings.Index(input[searchStart:], from)
		if idx == -1 {
			break
		}

		absIdx := searchStart + idx

		var b strings.Builder
		b.Grow(len(input) + len(to) - len(from))
		b.WriteString(input[:absIdx])
		b.WriteString(to)
		b.WriteString(input[absIdx+len(from):])

		results = append(results, b.String())
		searchStart = absIdx + len(from)
	}

	return results
}

// GreedyBackwardSearch works backwards from target to "e" with randomization
func GreedyBackwardSearch(target string, transformations []Transformation, maxAttempts int) int {
	reversed := ReverseTransformations(transformations)
	rand.Seed(time.Now().UnixNano())

	bestSteps := -1

	for attempt := 0; attempt < maxAttempts; attempt++ {
		current := target
		steps := 0

		// Keep reducing until we reach "e" or get stuck
		for current != "e" && steps < 1000 {
			// Find all possible reductions
			var allReductions []string
			var reductionIndices []int

			for i, trans := range reversed {
				reductions := FindAllReplacements(current, trans.from, trans.to)
				for _, red := range reductions {
					allReductions = append(allReductions, red)
					reductionIndices = append(reductionIndices, i)
				}
			}

			if len(allReductions) == 0 {
				break // Stuck, no more reductions possible
			}

			// Randomly pick a reduction (this helps explore different paths)
			idx := rand.Intn(len(allReductions))
			current = allReductions[idx]
			steps++

			// Early termination if we found "e"
			if current == "e" {
				if bestSteps == -1 || steps < bestSteps {
					bestSteps = steps
					fmt.Printf("Found solution in %d steps (attempt %d/%d)\n", steps, attempt+1, maxAttempts)
				}
				break
			}
		}
	}

	return bestSteps
}

// DeterministicBackwardSearch uses a greedy approach prioritizing shortest molecule
func DeterministicBackwardSearch(target string, transformations []Transformation) int {
	reversed := ReverseTransformations(transformations)
	current := target
	steps := 0

	for current != "e" && steps < 1000 {
		// Find all possible reductions
		var allReductions []string

		for _, trans := range reversed {
			reductions := FindAllReplacements(current, trans.from, trans.to)
			allReductions = append(allReductions, reductions...)
		}

		if len(allReductions) == 0 {
			return -1 // No solution
		}

		// Pick the shortest reduction (greedy heuristic)
		shortest := allReductions[0]
		for _, red := range allReductions[1:] {
			if len(red) < len(shortest) {
				shortest = red
			}
		}

		current = shortest
		steps++

		if steps%50 == 0 {
			fmt.Printf("Step %d, molecule length: %d\n", steps, len(current))
		}
	}

	if current == "e" {
		return steps
	}
	return -1
}

func main() {
	transformations := []Transformation{
		{"Al", "ThF"},
		{"Al", "ThRnFAr"},
		{"B", "BCa"},
		{"B", "TiB"},
		{"B", "TiRnFAr"},
		{"Ca", "CaCa"},
		{"Ca", "PB"},
		{"Ca", "PRnFAr"},
		{"Ca", "SiRnFYFAr"},
		{"Ca", "SiRnMgAr"},
		{"Ca", "SiTh"},
		{"F", "CaF"},
		{"F", "PMg"},
		{"F", "SiAl"},
		{"H", "CRnAlAr"},
		{"H", "CRnFYFYFAr"},
		{"H", "CRnFYMgAr"},
		{"H", "CRnMgYFAr"},
		{"H", "HCa"},
		{"H", "NRnFYFAr"},
		{"H", "NRnMgAr"},
		{"H", "NTh"},
		{"H", "OB"},
		{"H", "ORnFAr"},
		{"Mg", "BF"},
		{"Mg", "TiMg"},
		{"N", "CRnFAr"},
		{"N", "HSi"},
		{"O", "CRnFYFAr"},
		{"O", "CRnMgAr"},
		{"O", "HP"},
		{"O", "NRnFAr"},
		{"O", "OTi"},
		{"P", "CaP"},
		{"P", "PTi"},
		{"P", "SiRnFAr"},
		{"Si", "CaSi"},
		{"Th", "ThCa"},
		{"Ti", "BP"},
		{"Ti", "TiTi"},
		{"e", "HF"},
		{"e", "NAl"},
		{"e", "OMg"},
	}

	target := "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"

	fmt.Println("Backward search: target -> 'e'")
	fmt.Printf("Target length: %d characters\n\n", len(target))

	// Try deterministic approach first
	fmt.Println("=== Deterministic Greedy Search ===")
	steps := DeterministicBackwardSearch(target, transformations)
	if steps != -1 {
		fmt.Printf("\nMinimum transformations: %d\n\n", steps)
	} else {
		fmt.Println("Deterministic search failed, trying randomized...\n")

		// Fall back to randomized search
		fmt.Println("=== Randomized Search (10 attempts) ===")
		steps = GreedyBackwardSearch(target, transformations, 10)
		if steps != -1 {
			fmt.Printf("\nMinimum transformations found: %d\n", steps)
		} else {
			fmt.Println("No solution found in 10 attempts")
		}
	}
}
