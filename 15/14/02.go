package main

import (
	"bufio"
	"fmt"
	"os"
)

var TOTAL_TIME int = 2503

func totalDistance(s, d, r int) int {
	t := 0
	curr := 0

	for {
		t += d
		curr += s * d
		if t > TOTAL_TIME {
			curr -= s * (t - TOTAL_TIME)
			break
		}
		t += r
		if t >= TOTAL_TIME {
			break
		}
	}

	return curr
}

func main() {
	speeds := make(map[string]int)
	durations := make(map[string]int)
	rest := make(map[string]int)

	input, _ := os.Open("input.txt")
	defer input.Close()

	reader := bufio.NewReader(input)

	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}

		s := string(line)
		var name string
		var speed, duration, restTime int
		fmt.Sscanf(s, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &duration, &restTime)
		speeds[name] = speed
		durations[name] = duration
		rest[name] = restTime
	}

	max := 0

	for name, _ := range speeds {
		curr := totalDistance(speeds[name], durations[name], rest[name])
		if curr > max {
			max = curr
		}
	}
	fmt.Println(max)
}
