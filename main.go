package main

import( "fmt"
	"flag")


func main(){
	//define a flag
	versionFlag := flag.Bool("version",false,"show version")	
	versionFlagShort := flag.Bool("v",false,"show version (short)")
	addTaskName := flag.String("add-task","","Name of the task")
	addDuration := flag.String("duration","","duration of the task in hh:mm:ss")
	//parse the flags
	flag.Parse()
	if *versionFlag==true || *versionFlagShort==true{
		checkVersion()
	}
	
	if flag.Lookup("add-task").Value.String()!=""{
		if flag.Lookup("duration").Value.String()!=""{
			addTask(*addTaskName,*addDuration)
		}else{
			addTask(*addTaskName,"00:10:00")
		}
	}
}

// check health 
func checkHealth(){
	}

//check version
func checkVersion(){
// command : tasker --version
	// this command should print the current version of the application. 
	// --version flag is used or in short -v.
	fmt.Printf("Tasker\nversion: v0.0.1\n")	
}
// create a task
func addTask(taskName ,time string){
	fmt.Printf(taskName+" & "+time+"\n")
}
