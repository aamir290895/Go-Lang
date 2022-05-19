package logic

import (
	
	"strings"
)

func CheckVowel(s string) string{
    var output string
    if (strings.ContainsAny(s,"aeiou") && strings.ContainsAny(s,"AEIOU")){
	    output = "lovely string"


	}else if (strings.ContainsAny(s,"aeiou") || strings.ContainsAny(s,"AEIOU")){

		output = "ugly string"
	}

   return output

}