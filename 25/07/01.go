package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	res := 0
	grid := make([][]rune, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		grid = append(grid, []rune(s))
	}

	for _, r := range grid {
		fmt.Println(string(r))
	}

	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	i, j := 0, 70
	currentRays := [][]int{[]int{i, j}}
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println(res)
		for _, r := range grid {
			fmt.Println(string(r))
		}
		for x, r := range grid {
			for y, _ := range r {
				if grid[x][y] == '|' && !visited[x][y] {
					currentRays = append(currentRays, []int{x, y})
					visited[x][y] = true
				}
			}
		}

		for len(currentRays) > 0 {
			curr := currentRays[0]
			currentRays = currentRays[1:]
			if curr[0]+1 < n {
				if grid[curr[0]+1][curr[1]] == '^' &&
					curr[1]+1 < m &&
					curr[1]-1 >= 0 {
					res++
					currentRays = append(currentRays, []int{curr[0], curr[1] - 1})
					currentRays = append(currentRays, []int{curr[0], curr[1] + 1})
				} else {
					grid[curr[0]+1][curr[1]] = '|'
				}
			} else {
				return
			}
		}
	}
}
