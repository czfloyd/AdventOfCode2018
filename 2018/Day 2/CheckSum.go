package main

import (
	"bufio"
    "fmt"
	"os"
)

func main()  {
	file,err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	var s []string
	var m []map[rune]int
	var twos int = 0
	var threes int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	
	for i, element := range s {
		m = append(m, make(map[rune]int))
		for _,l := range element {
			if m[i][l] >= 0 {
				m[i][l]++
			}else{
				m[i][l] = 0
			}
		}
		var f2 = false
		var f3 = false
		for _,v := range m[i] {
			if v == 2 && !f2 {
				twos++
				f2 = true
			}else if v == 3 && !f3 {
				threes++
				f3 = true
			}
		}
	}
	fmt.Printf("%d %d\n", twos, threes)
	fmt.Print(twos * threes)
}