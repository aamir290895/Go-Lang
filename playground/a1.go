// package main

// import (
 
// 	"os"
// 	"fmt"

// )

// var (
// 	L int32
// 	M float32
// 	N bool
// 	O string
// )

// func CreateFile(){

//  os.Create("empty.txt")


// }

// func main(){
// 	file, err := os.Open("empty.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
	
// 	fmt.Fscan(file, &L, &M, &N, &O)		
// 	fmt.Printf("\nL M N O: %v %v %v %v", L,M,N,O)
	
// 	fmt.Fscanln(file, &L, &M, &N, &O)
// 	fmt.Printf("\nL M N O: %v %v %v %v", L,M,N,O)
	
// 	fmt.Fscanf(file, "%s %d %t %f",&O, &L, &N, &M)
// 	fmt.Printf("\nO L N M: %v %v %v %v", O,L,N,M)

}