package logic

import (
	"fmt"
	"strings"
)

func CheckSubString (inputString,s string)string{


	
    var output string
	// var i string
	// var check string


	// fmt.Scanf("%s",&i)
	// fmt.Scanf("%s",&check)


	// val, err := strconv.Atoi(i)


	if inputString == s {
		output = "YES"
	}else if strings.Contains(inputString , s){

		// split_input := strings.Split(inputString , "")
		// split_subString  := strings.Split(s,"")

		// if (&split_input[len(split_input)-1] == &split_subString[len(split_subString)-1]){

		// 	fmt.Println("YES")
		// }else{
		// 	fmt.Println("NO")
		// }
        // a1 := len(inputString)-1
		// b1 := len(s)-1

		// if (inputString[a1-b1:] == s ){

		// 	fmt.Println("YES")
		// }else{
		// 	fmt.Println("NO")
		// }


		output = "YES"
	    
		
		
	}else{
		output = "NO"
	}
    
	return output
}

func CheckInteger (inputInt int) {


	

	if (1<= inputInt) && (inputInt <=10){

		fmt.Println("YES")
	} else{
		fmt.Println("YES")
	}
}


