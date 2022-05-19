package main

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

// }


// func main(){

// 	var numOfStrings int

// 	var array_input[numOfStrings]string


// 	file1,err := os.Open("input1.txt")

// 	if err != nil {
// 	 panic(err)
// 	 }
// 	defer file1.Close()
	


// 	scanner := bufio.NewScanner(file1)
// 	// optionally, resize scanner's capacity for lines over 64K, see next example
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())

// 		capacity, err := strconv.Atoi(scanner.Text());

// 		if(err == nil){

// 		   numOfStrings = capacity
// 		}


// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

//    fmt.Println(numOfStrings)


// }