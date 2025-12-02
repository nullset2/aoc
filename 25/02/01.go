package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sillies := int64(0)
	input := "16100064-16192119,2117697596-2117933551,1-21,9999936269-10000072423,1770-2452,389429-427594,46633-66991,877764826-877930156,880869-991984,18943-26512,7216-9427,825-1162,581490-647864,2736-3909,39327886-39455605,430759-454012,1178-1741,219779-244138,77641-97923,1975994465-1976192503,3486612-3602532,277-378,418-690,74704280-74781349,3915-5717,665312-740273,69386294-69487574,2176846-2268755,26-45,372340114-372408052,7996502103-7996658803,7762107-7787125,48-64,4432420-4462711,130854-178173,87-115,244511-360206,69-86"
	//test_input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	ranges := strings.Split(input, ",")

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)

			for j := 1; j <= (len(s)/2)+1; j++ {
				slice := s[0:j]
				repetitions := strings.Repeat(slice, 2)
				converted, err := strconv.Atoi(repetitions)
				if err != nil {
					panic(err)
				}
				if converted == i {
					fmt.Println("Silly found:", i, converted)
					sillies += int64(converted)
				}
			}
		}

	}

	fmt.Println(sillies)

}
