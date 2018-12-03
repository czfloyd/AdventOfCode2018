package main

import (
	"bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
)

func main()  {
	file,err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	var s []string
	var m map[int]map[int]int
	m = make(map[int]map[int]int)
	var out int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	
	for _, element := range s {
		var n [5]int
		var words = strings.Fields(element)
		num1,_ := strconv.Atoi(words[0][1:])
		n[0] = num1
		var pos = strings.Split(words[2], ",")
		num2,_ := strconv.Atoi(pos[0])
		n[1] = num2
		num3,_ := strconv.Atoi(pos[1][:len(pos[1]) - 1])
		n[2] = num3
		var size = strings.Split(words[3], "x")
		num4,_ := strconv.Atoi(size[0])
		n[3] = num4
		num5,_ := strconv.Atoi(size[1])
		n[4] = num5
		for i := n[1]; i < n[1] + n[3]; i++ {
			if (m[i] == nil){
				m[i] = make(map[int]int)
			}
			for j := n[2]; j < n[2] + n[4]; j++ {
				if m[i][j] >= 0 {
					m[i][j]++
				}else{
					m[i][j] = 1
				}
			}
		}
	}

	for _, element := range s {
		var unique = true

		var n [5]int
		var words = strings.Fields(element)
		num1,_ := strconv.Atoi(words[0][1:])
		n[0] = num1
		var pos = strings.Split(words[2], ",")
		num2,_ := strconv.Atoi(pos[0])
		n[1] = num2
		num3,_ := strconv.Atoi(pos[1][:len(pos[1]) - 1])
		n[2] = num3
		var size = strings.Split(words[3], "x")
		num4,_ := strconv.Atoi(size[0])
		n[3] = num4
		num5,_ := strconv.Atoi(size[1])
		n[4] = num5
		for i := n[1]; i < n[1] + n[3] && unique; i++ {
			for j := n[2]; j < n[2] + n[4]; j++ {
				if m[i][j] > 1 {
					unique = false
					break
				}
			}
		}
		if unique {
			fmt.Printf("%d\n",n[0])
		}
	}

	fmt.Print(out)
}