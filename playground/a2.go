package main


import (
	"fmt"
	"strings"
)

func main() {
	var s string
	s = "hello"
	fmt.Println(strings.Split(s, ""))
	var input string
	fmt.Scanf("%s", &input)
    


	if strings.Contains(input , "puter"){

		fmt.Println("Yes")
	}
	
	fmt.Println(input)
}
