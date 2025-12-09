package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	To   string
	Dist int
}

func main() {
	graph := make(map[string][]Entry)
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)

	// Parse input file
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		parts := strings.Split(s, " = ")
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		points := strings.Split(parts[0], " to ")
		from, to := points[0], points[1]

		// Add both directions for undirected graph
		graph[from] = append(graph[from], Entry{To: to, Dist: dist})
		graph[to] = append(graph[to], Entry{To: from, Dist: dist})
	}

	// Get all cities
	cities := make([]string, 0, len(graph))
	for city := range graph {
		cities = append(cities, city)
	}

	// Find shortest path visiting each city exactly once
	minDist, path := findShortestPath(graph, cities)

	// Print results
	fmt.Println("Shortest path visiting each city once:")
	for i := 0; i < len(path)-1; i++ {
		fmt.Printf("%s -> ", path[i])
	}
	fmt.Printf("%s\n", path[len(path)-1])
	fmt.Printf("\nTotal distance: %d\n", minDist)
}

func findShortestPath(graph map[string][]Entry, cities []string) (int, []string) {
	minDist := math.MaxInt
	var bestPath []string

	// Try starting from each city
	for i := 0; i < len(cities); i++ {
		visited := make(map[string]bool)
		path := []string{cities[i]}
		visited[cities[i]] = true

		dist := dfs(graph, cities[i], visited, path, 0, &minDist, &bestPath)
		if dist < minDist {
			minDist = dist
		}
	}

	return minDist, bestPath
}

func dfs(graph map[string][]Entry, current string, visited map[string]bool, path []string, dist int, minDist *int, bestPath *[]string) int {
	// If all cities visited, check if this is the best path
	if len(visited) == len(graph) {
		if dist < *minDist {
			*minDist = dist
			*bestPath = make([]string, len(path))
			copy(*bestPath, path)
		}
		return dist
	}

	// Prune if current distance already exceeds minimum
	if dist >= *minDist {
		return math.MaxInt
	}

	result := math.MaxInt

	// Try visiting each unvisited neighbor
	for _, entry := range graph[current] {
		if !visited[entry.To] {
			visited[entry.To] = true
			path = append(path, entry.To)

			newDist := dfs(graph, entry.To, visited, path, dist+entry.Dist, minDist, bestPath)
			if newDist < result {
				result = newDist
			}

			// Backtrack
			path = path[:len(path)-1]
			delete(visited, entry.To)
		}
	}

	return result
}
