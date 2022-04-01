package main

import (
	"fmt"
)

func main() {
	t := "Hello"
	fmt.Printf("вся строка -  %T \n", t)
	fmt.Printf("одиночный индекс строки - %T \n", t[1])
	fmt.Printf("срез строки - %T \n", t[1:2])

}
