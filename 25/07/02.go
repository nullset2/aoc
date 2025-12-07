package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]int, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}
		row := make([]int, 0)
		for _, c := range s {
			if c == '.' {
				row = append(row, 0)
			} else {
				row = append(row, -1)
			}
		}
		grid = append(grid, row)
	}

	for _, r := range grid {
		fmt.Println(r)
	}

	n := len(grid)
	m := len(grid[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	i, j := 0, 70
	currentRays := [][]int{[]int{i, j, 1}}
	for {
		fmt.Print("\033[H\033[2J")
		for _, v := range grid {
			for i, vv := range v {
				if i > 0 {
					fmt.Print("  ")
				}
				fmt.Printf("%2d", vv)
			}
			fmt.Println()
		}
		for x, r := range grid {
			for y, _ := range r {
				if grid[x][y] > 0 && !visited[x][y] {
					currentRays = append(currentRays, []int{x, y, grid[x][y]})
					visited[x][y] = true
				}
			}
		}

		for len(currentRays) > 0 {
			curr := currentRays[0]
			currentRays = currentRays[1:]
			if curr[0]+1 < n {
				if grid[curr[0]+1][curr[1]] == -1 &&
					curr[1]+1 < m &&
					curr[1]-1 >= 0 {
					currentRays = append(currentRays, []int{curr[0], curr[1] - 1, curr[2]})
					currentRays = append(currentRays, []int{curr[0], curr[1] + 1, curr[2]})
				} else {
					grid[curr[0]+1][curr[1]] += curr[2]
				}
			} else {
				lastRow := grid[n-1]
				sum := 0
				for _, x := range lastRow {
					sum += x
				}
				fmt.Println(sum)
				return
			}
		}
	}
}
