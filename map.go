package main

import (
	"fmt"
)

func makeArray(num int, t string) []string {
	switch num {
	case 0:
		return ([]string{t[1:2], t[3:4], t[4:5]})
	case 1:
		return []string{t[0:1], t[2:3], t[3:4], t[4:5], t[5:6]}
	case 2:
		return []string{t[1:2], t[4:5], t[5:6]}
	case 3:
		return []string{t[0:1], t[1:2], t[4:5], t[6:7], t[7:8]}
	case 4:
		return []string{t[0:1], t[1:2], t[2:3], t[3:4], t[5:6], t[6:7], t[7:8], t[8:]}
	case 5:
		return []string{t[1:2], t[2:3], t[4:5], t[7:8], t[8:]}
	case 6:
		return []string{t[3:4], t[4:5], t[7:8]}
	case 7:
		return []string{t[3:4], t[4:5], t[5:6], t[6:7], t[8:]}
	case 8:
		return []string{t[4:5], t[5:6], t[7:8]}
	}
	return nil
}

func makeMap(i string) map[string][]string {
	m := make(map[string][]string)
	for index, num := range i {
		m[string(num)] = makeArray(index, i)
	}
	return m
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func get_Time(m map[string][]string, c string) int {
	prev_button := "nil"
	time := 0
	for _, num := range c {
		if Contains(m[prev_button], string(num)) {
			time += 1
		} else if prev_button == string(num) {
			time += 0
		} else if prev_button == "nil" {
			time += 0
		} else {
			time += 2
		}
		prev_button = string(num)
	}
	return time
}
func main() {
	var i, code string

	code = "912132132155672245566165" //getNums()
	i = "639485712"                   //getNums()
	res := get_Time(makeMap(i), code)
	fmt.Print(res)
}
