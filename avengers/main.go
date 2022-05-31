package main

import
(
"fmt"
"avengers/mission"
"bufio"
"os"
"strconv"


)

func CaptainFury(){
  fmt.Println("========SHIELD=============") 
  fmt.Println("Welcome to captain Fury:\n 1. Check the mission\n 2. Assigned missions to avengers\n 3. Check Mission's detail\n 4. Check Avengers detail's\n 5. Update Mission's status\n 6. List of avenger's\n 7. Assign avenger to mission\n")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()

  input, _ := strconv.ParseInt(scanner.Text() ,10 ,64)
  
  switch input{
  case 1 :
     mi.CheckMission()
  case 2 :
     mi.AssignMission()
  case 3 :
     mi.MissionDetails()
  case 4 :
     mi.CheckAvengersDetail()

  case 5 : 
     mi.UpdateMissionStatus()
  case 6 :
     mi.ListOfAvengers()
  case 7 :
     mi.AssignAvengersToMission()
  default : 
     fmt.Println("Invalid input:")  
  }

  CaptainFury()
}

func main(){

   CaptainFury()
 

}



