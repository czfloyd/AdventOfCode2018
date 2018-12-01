package main

import (
	"bufio"
    "fmt"
	"os"
	"strconv"
)

func main()  {
	file,err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	var s []string
	var count int = 0
	var m map[int]bool
	m = make(map[int]bool)
	var d bool = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	
	for d {
		for _, element := range s {
			if string(element[0]) == "+" {
				num,_ := strconv.Atoi(element[1:])
				count += num
			} else {
				num,_ := strconv.Atoi(element[1:])
				count -= num
			}
			if m[count] {
				fmt.Print(count)
				d = false
				break
			}
			m[count] = true
		}
	}
}