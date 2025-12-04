package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	res := 0
	grid := make([][]rune, 0)
	copygrid := make([][]rune, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			break
		}

		grid = append(grid, []rune(s))
		copygrid = append(copygrid, []rune(s))
	}

	n := len(grid)
	m := len(grid[0])
	found := false
	oldRes := 0

	for !found {
		fmt.Print("\033[H\033[2J")
		for i := range grid {
			grid[i] = copygrid[i]
		}
		oldRes = res
		for i := range grid {
			fmt.Println(string(grid[i]))
		}
		for i := range grid {
			fmt.Println(string(grid[i]))
			for j := range grid[i] {
				if grid[i][j] == '@' {
					neighbors := 0

					if j+1 < m && grid[i][j+1] == '@' {
						neighbors++
					}

					if i+1 < n && j+1 < m && grid[i+1][j+1] == '@' {
						neighbors++
					}

					if i+1 < n && grid[i+1][j] == '@' {
						neighbors++
					}

					if i+1 < n && j-1 >= 0 && grid[i+1][j-1] == '@' {
						neighbors++
					}

					if j-1 >= 0 && grid[i][j-1] == '@' {
						neighbors++
					}

					if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == '@' {
						neighbors++
					}

					if i-1 >= 0 && grid[i-1][j] == '@' {
						neighbors++
					}

					if i-1 >= 0 && j+1 < m && grid[i-1][j+1] == '@' {
						neighbors++
					}

					if neighbors < 4 {
						copygrid[i][j] = '.'
						res++
						found = true
						break
					}
				}
			}
			if found {
				found = false
				break
			}
		}

		if res == oldRes {
			break
		}
	}
	fmt.Println(res)
}
