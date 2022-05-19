package logic


func SuperDigit(n int) int{

   var sum int

   if n==0 {
     sum =0
   }else if n%9 == 0{
     sum = 9
   }else{
	 sum = n%9
   }
   

 
   return sum;


}

