package mi

import (
	"bufio"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/json"
)


type Shield struct {

	Mission_Name []string
	Mission_Details []string
	Status string
	

}



type Avenger struct {
   Name string       `json:"name"`
   Real_Name string   `json:"real_name"`
   Abilities string    `json:"abilities"`
   Mission_Assigned int 
   Mission_Completed int `json:"mission_completed"`

}


var shield [10]Shield

var avengers []Avenger

func JsonObj() int{

	file,_ := ioutil.ReadFile("mission/data.json")
    
    var data []Avenger
	err := json.Unmarshal(file, &data)

	if (err != nil){
       fmt.Println("Error in parse json")

	}
	avengers = append(avengers, data...)
		
	return len(data)

}

var Count int = JsonObj()

// func InputDetails(i int , mission_details []string ,mission []string){
	   
// 	    a := append(shield[i].Mission_Name, mission...)
	
// 		b := append(shield[i].Mission_Details, mission_details...)

// 		fmt.Println(a ,b)
// 		shield[i].Status = "Assigned"
// }

func AssignMission() {
	
	
	var alert string
    
	fmt.Println("Enter Avenger's Name: \n")

	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	names := strings.Split(scanner1.Text(),",")


	
	
	fmt.Println("Enter mission name: \n")

	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	mission := scanner2.Text()


	fmt.Println("Enter mission details: \n")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	mission_details := scanner3.Text()
    
	var slice_mission []string 
	var slice_mission_details []string

	slice_mission = append(slice_mission ,mission)
	slice_mission_details = append(slice_mission_details,mission_details)
	
   if (len(names) == 2 ) {
	for i :=0; i<=Count-1; i++ {
		
            
        if (avengers[i].Name ==names[0] ||avengers[i].Name == names[1] && len(shield[i].Mission_Name)<=2 ){
	     	    
			shield[i].Mission_Name = append(shield[i].Mission_Name, slice_mission...)
	
			shield[i].Mission_Details = append(shield[i].Mission_Details, slice_mission_details...)
	
			shield[i].Status = "Assigned"
            
				
		     
	    } 
		
	  }  
	  fmt.Println("The Mission has assigned .Email Notification has send\n\n\n")
    
	}else if (len(names)==1){
        for j :=0; j<=Count-1; j++{
            
			if (avengers[j].Name ==names[0] && len(shield[j].Mission_Name) <=2){
					 
	
				shield[j].Mission_Name = append(shield[j].Mission_Name, slice_mission...)
	
				shield[j].Mission_Details = append(shield[j].Mission_Details, slice_mission_details...)
		
				shield[j].Status = "Assigned"
				
				 
			}
			
		  }  
		  fmt.Println("The Mission has assigned .Email Notification has send\n\n\n")
    

	}else {

	}		
    
     fmt.Println(alert)     
}


func CheckMission(){
	fmt.Println("Mission name \t| Avengers \t | Status")

    
    
	var j int = 0
	for i := 0; i<Count; i++ {
    
    	if len(shield[i].Mission_Name) != 0 {

			for _,n := range shield[i].Mission_Name {
				fmt.Println(n + "\t|\t" + avengers[i].Name + "\t|\t" + shield[i].Status )
			}
		}else if (len(shield[i].Mission_Name) == 0){
             j++
          
		}
		  	
    	}

	if ( j == Count){
		fmt.Println("No avenger's assign to mission")
	}
	
}

func FindIndex(a []string, x string) int {
   var out int
    for i, n := range a {
        if x == n {
            out = i
        }
    }
    return out
}

func CheckPresence(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}



func MissionDetails(){

	var mission string
    fmt.Println("Enter Mission Name:")

	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	mission = scanner1.Text()
    
	for i :=0;i<=Count;i++ {
		if (CheckPresence(shield[i].Mission_Name ,mission)){

			for _,n := range shield[i].Mission_Details{
				fmt.Println("Mission Details: " + n + "\t"+" Mission Status: "+ shield[i].Status + "\n\n\n")


			}

		}

	}
}



func CheckAvengersDetail() {

	var av_name string
    fmt.Println("Enter Avenger Name:")

	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	av_name = scanner1.Text()
    
	for i :=0;i<= Count;i++ {
		if ( av_name == avengers[i].Name){

			fmt.Println("Avenger Name:" + avengers[i].Name+"\n" + "Person Name:" + avengers[i].Real_Name +"\n"+ "Abilities: " + avengers[i].Abilities+   "\n"+ " Mission Status: "+ shield[i].Status + "\n\n\n")
		}

	}

}


func UpdateMissionStatus(){

  var mission string	
  fmt.Println("Enter Mission's name :")
  scanner1 := bufio.NewScanner(os.Stdin)
  scanner1.Scan()
  mission = scanner1.Text()

  for i :=0; i<=Count; i++ {

	if (CheckPresence(shield[i].Mission_Name , mission)){
		fmt.Println("Enter new status:")
		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		status := scanner2.Text()

		shield[i].Status = status

		avengers[i].Mission_Completed ++
	}
  }
  
}

func ListOfAvengers(){

	for i:=0; i<=Count; i++ {
       if (shield[i].Status == ""){
		fmt.Println("Mission" +  "|\t" + avengers[i].Name  + "|\t", "Available")


	   }else{
		 fmt.Println(avengers[i].Name,"|\t","On the mission","|\t" ,shield[i].Mission_Name )  
	   }

	}


}


func AssignAvengersToMission(){

	fmt.Println("Enter Avenger name :")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	avenger := scanner1.Text()

	fmt.Println("Enter Mission name :")

	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	mission := scanner2.Text()

	for i := 0; i<=Count-1; i++{

		if (avengers[i].Name == avenger){
            if (len(shield[i].Mission_Name) == 0){
				shield[i].Mission_Name = append (shield[i].Mission_Name,mission)
			}else {
				fmt.Println(avengers[i].Name +" is already on the mission\n\n\n")
			}
		}
	}


}

