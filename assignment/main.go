package main

import (
	
	"assignment/logics"
	"fmt"
	// "strings"
	// "strconv"
    //  "os"
	//  "bufio"
	//  "log"
	//  "strconv"

)

type Data struct {
   num int

   str string

}

func main(){

	fmt.Println("output of assignment 1 :")

	a,a1,b,b1,c,c1,d,d1 := "abcdef","def","computer","muter","stringmatchingmat","ingmat","videobox","videobox"

    o1 := logic.CheckSubString(a,a1)
	o2 := logic.CheckSubString(b,b1)
	o3 := logic.CheckSubString(c,c1)
	o4 := logic.CheckSubString(d,d1)

	fmt.Println(o1,o2,o3,o4)


	fmt.Println("output of assignment 2 :")




    super_digit := 12351511321231
	fmt.Println(logic.SuperDigit(super_digit))

	fmt.Println("output of assignment 3 :")

	e,f,g := "omahgoTuRuLob","OmAhgotUrulobEI" ,"aeKORONAoiBATCHu"

	fmt.Println(logic.CheckVowel(e))
	fmt.Println(logic.CheckVowel(f))
	fmt.Println(logic.CheckVowel(g))



}

