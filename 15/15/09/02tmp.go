package main

import (
	"bufio"
	"fmt"
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

	// Find longest path visiting each city exactly once
	maxDist, path := findLongestPath(graph, cities)

	// Print results
	fmt.Println("Longest path visiting each city once:")
	for i := 0; i < len(path)-1; i++ {
		fmt.Printf("%s -> ", path[i])
	}
	fmt.Printf("%s\n", path[len(path)-1])
	fmt.Printf("\nTotal distance: %d\n", maxDist)
}

func findLongestPath(graph map[string][]Entry, cities []string) (int, []string) {
	maxDist := 0
	var bestPath []string

	// Try starting from each city
	for i := 0; i < len(cities); i++ {
		visited := make(map[string]bool)
		path := []string{cities[i]}
		visited[cities[i]] = true

		dist := dfs(graph, cities[i], visited, path, 0, &maxDist, &bestPath)
		if dist > maxDist {
			maxDist = dist
		}
	}

	return maxDist, bestPath
}

func dfs(graph map[string][]Entry, current string, visited map[string]bool, path []string, dist int, maxDist *int, bestPath *[]string) int {
	// If all cities visited, check if this is the best path
	if len(visited) == len(graph) {
		if dist > *maxDist {
			*maxDist = dist
			*bestPath = make([]string, len(path))
			copy(*bestPath, path)
		}
		return dist
	}

	result := 0

	// Try visiting each unvisited neighbor
	for _, entry := range graph[current] {
		if !visited[entry.To] {
			visited[entry.To] = true
			path = append(path, entry.To)

			newDist := dfs(graph, entry.To, visited, path, dist+entry.Dist, maxDist, bestPath)
			if newDist > result {
				result = newDist
			}

			// Backtrack
			path = path[:len(path)-1]
			delete(visited, entry.To)
		}
	}

	return result
}
