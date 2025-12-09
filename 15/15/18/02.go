package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := 100
	grid := make([][]rune, N)
	for i := range grid {
		grid[i] = make([]rune, N)
	}
	input, _ := os.Open("input.txt")
	defer input.Close()
	reader := bufio.NewReader(input)
	i := 0

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		s := string(line)
		for j, c := range s {
			grid[i][j] = c
		}
		i++
	}

	for x := 0; x < N; x++ {
		newGrid := make([][]rune, N)
		for i := range newGrid {
			newGrid[i] = make([]rune, N)
		}

		grid[0][0] = '#'
		grid[0][N-1] = '#'
		grid[N-1][0] = '#'
		grid[N-1][N-1] = '#'

		fmt.Println("Old:")
		for i := 0; i < N; i++ {
			fmt.Println(string(grid[i]))
		}

		reader := bufio.NewReader(os.Stdin)
		reader.ReadRune()

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				neighborsOn := 0
				//top
				if i-1 >= 0 && grid[i-1][j] == '#' {
					neighborsOn++
				}

				//top right
				if i-1 >= 0 && j+1 < N && grid[i-1][j+1] == '#' {
					neighborsOn++
				}

				//right
				if j+1 < N && grid[i][j+1] == '#' {
					neighborsOn++
				}

				//bottom right
				if i+1 < N && j+1 < N && grid[i+1][j+1] == '#' {
					neighborsOn++
				}

				//bottom
				if i+1 < N && grid[i+1][j] == '#' {
					neighborsOn++
				}

				//bottom left
				if i+1 < N && j-1 >= 0 && grid[i+1][j-1] == '#' {
					neighborsOn++
				}

				//left
				if j-1 >= 0 && grid[i][j-1] == '#' {
					neighborsOn++
				}

				//top left
				if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == '#' {
					neighborsOn++
				}

				if grid[i][j] == '#' && (neighborsOn == 2 || neighborsOn == 3) {
					newGrid[i][j] = '#'
				} else if grid[i][j] == '.' && neighborsOn == 3 {
					newGrid[i][j] = '#'
				} else {
					newGrid[i][j] = '.'
				}

			}
		}

		newGrid[0][0] = '#'
		newGrid[0][N-1] = '#'
		newGrid[N-1][0] = '#'
		newGrid[N-1][N-1] = '#'


		fmt.Println("New:")
		for i := 0; i < N; i++ {
			fmt.Println(string(newGrid[i]))
		}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				grid[i][j] = newGrid[i][j]
			}
		}

		reader.ReadRune()
	}

	res := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '#' {
				res++
			}
		}
	}

	fmt.Println(res)
}
