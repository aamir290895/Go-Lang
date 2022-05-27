package mi

import (
	"bufio"
	// // "strconv"
	 "os"
	"io/ioutil"
	
	"fmt"
	"strings"
	"math/rand"
	"time"
	"encoding/json"


)


type Shield struct {

	Mission_Name string
	Mission_Details string
	Status string
    Avengers []Avenger `json : "data"`
	

}



type Avenger struct {
   Name string       `json : "Name"`
   Real_Name string   `json : "Real_Name`
   Abilities string    `json : "Abilities"`
   Mission_Assigned int 
   Mission_Completed int `json : "Mission_Completed`

}


var Count int = 5
var shield [10]Shield

var avengers [10]Avenger
func CheckAvengers() bool{

	rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) == 1

}

func JsonObj(){
    jsonFile,_ := os.Open("data.json")
	file,_ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	var data Shield
 
	json.Unmarshal(file, &data)

	
    fmt.Println(data)
	


}


func Details(){
   avengers[0].Name = "Thor"
   avengers[0].Abilities = "Superhuman strength, speed, agility, durability and immunity to most diseases."
   avengers[0].Real_Name = "Chris Hemsworth"
   avengers[0].Mission_Completed = 0

   avengers[1].Name = "Hawkeye"
   avengers[1].Abilities = "Superhuman strength and durability Size and mass manipulation."
   avengers[1].Real_Name = "Clint Barton"
   avengers[1].Mission_Completed = 0

   avengers[2].Name = "Hulk"
   avengers[2].Abilities = "Superhuman strength, speed, agility, durability and immunity to most diseases."
   avengers[2].Real_Name = "Robert Bruce"
   avengers[2].Mission_Completed = 0

   avengers[3].Name = "Black Widow"
   avengers[3].Abilities = "Superhuman strength, speed, agility, durability and immunity to most diseases."
   avengers[3].Real_Name = "Natasha Romanova"
   avengers[3].Mission_Completed = 0


   avengers[4].Name = "Captain America"
   avengers[4].Abilities = "Superhuman strength, speed, agility, durability and immunity to most diseases."
   avengers[4].Real_Name = "Chris Evens"
   avengers[4].Mission_Completed = 0


}

func InputDetails(i int , mission_details string ,mission string){
	      
		shield[i].Mission_Details = mission_details
		shield[i].Mission_Name= mission
		shield[i].Status = "Assigned"
}

func AssignMission() {
	

	// var avenger Avengers

	
    JsonObj()
	fmt.Println("Enter Avenger's Name: \n")

	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	names := strings.Split(scanner1.Text(),",")
	// var k [2]int

	// fmt.Scanf("%s" ,&name)

	fmt.Println("Enter mission name: \n")

	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	mission := scanner2.Text()
	// fmt.Scanf("%s",&mission)


	fmt.Println("Enter mission details: \n")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	mission_details := scanner3.Text()
	
   if (len(names) == 2) {
	for i :=0; i<=Count-1; i++{
            
        if (avengers[i].Name ==names[0] ||avengers[i].Name == names[1]){
	     	    

                InputDetails(i,mission_details,mission)
				
		     
	    } 

		
	  }      
	}else if (len(names)==1){
        for j :=0; j<=Count-1; j++{
            
			if (avengers[j].Name ==names[0]){
					 
	
					InputDetails(j,mission_details,mission)
				 
			} 
	
			
		  }      

	}		
    
	fmt.Println("The Mission has assigned .Email Notification has send\n\n\n")
          
}


func CheckMission() string{

	var output string
	for i := 0; i<=Count-1; i++ {

	if shield[i].Mission_Name == "" && len(shield) ==Count{
		output = "No mission assigned to avengers\n\n\n"
	}else{


			if shield[i].Mission_Name != "" {
				fmt.Println(shield[i].Mission_Name + "\t|\t" + avengers[i].Name + "\t|\t" + shield[i].Status )


			}
		  
		}
	}

	
	return output 
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


func MissionDetails(){

	var mission string
    fmt.Println("Enter Mission Name:")

	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	mission = scanner1.Text()
    
	for i :=0;i<=Count;i++ {
		if ( mission == shield[i].Mission_Name){

			fmt.Println("Mission Details: " + shield[i].Mission_Details + "Mission Status: "+ shield[i].Status + "\n\n\n")
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

	if (mission == shield[i].Mission_Name){
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
		 fmt.Println(avengers[i].Name,"|\t","On the mission","|\t" ,shield[i].Mission_Name +"\n\n\n")  
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
            if (shield[i].Mission_Name == ""){
				shield[i].Mission_Name = mission
			}else {
				fmt.Println(avengers[i].Name +" is already on the mission\n\n\n")
			}
		}
	}


}

