package main

import( "fmt"
	"flag"
	"Task"
)


func main(){
	//define a flag
	versionFlag := flag.Bool("version",false,"show version")	
	versionFlagShort := flag.Bool("v",false,"show version (short)")
	addTaskName := flag.String("add-task","","Name of the task")
	addDuration := flag.String("duration","","duration of the task in hh:mm:ss")
	helpFlag := flag.Bool("help",false,"show help")
	helpFlagShort := flag.bool("h",false,"show help(short)")
	//parse the flags
	flag.Parse()

	if *helpFlagShort == true || *versionFlag==true{
		showHelp()
	}
	if *versionFlag==true || *versionFlagShort==true{
		checkVersion()
	}
	
	if flag.Lookup("add-task").Value.String()!=""{
		if flag.Lookup("duration").Value.String()!=""{
			Task.AddTask(*addTaskName,*addDuration)
		}else{
			Task.AddTask(*addTaskName,"00:10:00")
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


