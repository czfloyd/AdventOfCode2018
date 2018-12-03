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
	var out []rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	
	for _, e1 := range s {
		for _, e2 := range s {
			if offByOne(e1, e2) {
				fmt.Printf("%s %s\n", e1, e2)
				for i, r := range e1 {
					if e1[i] == e2[i] {
						out = append(out, r)
					}
				}
			}
		}
	}
}

func offByOne(w1 string, w2 string) bool {
	if w1 == w2{
		return false
	}
	var m int = 1
	
	for i,_ := range w1 {
		if w1[i] != w2[i] {
			m--;
			if m < 0 {
				return false
			}
		}
	}

	return true
}