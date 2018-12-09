package main

import (
	"bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
	"sort"
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
	var g map[int]int
	g = make(map[int]int)
	var out int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	sort.Strings(s)

	var guard int = 0
	var ftime int = 0
	
	for _, element := range s {
		var words = strings.Fields(element)
		if (words[2] == "Guard"){
			guard,_ = strconv.Atoi(words[3][1:])
		}else if (words[2] == "falls"){
			ftime,_ = strconv.Atoi(words[1][3:5])
		}else{
			var atime,_ = strconv.Atoi(words[1][3:5])
			for i := ftime; i < atime; i++{
				if (m[guard] == nil){
					m[guard] = make(map[int]int)
				}
				if (m[guard][i] > 0){
					m[guard][i]++
				}else{
					m[guard][i] = 1
				}
				if (g[guard] > 0){
					g[guard]++
				}else{
					g[guard] = 1
				}
			}
		}
	}

	var max int = 0
	var maxmin int = 0
	for gu,_ := range g {
		for k,v := range m[gu]{
			if (v > m[max][maxmin]){
				max = gu
				maxmin = k
			}
		}
	}

	out = max * maxmin
	fmt.Print(out)
}