package main

import( "fmt"
	"flag"
	"github.com/garvityadav/Task-tracker/controller"
)


func main(){
	//define a flag
	
	//version Flag
	var version bool
	flag.Bool(&version,"version",false,"show version")	
	flag.Bool(&version,"v",false,"show version (short)")

	//Task flag
	var addTaskName string
	flag.String(&addTaskName,"add-task","","Name of the task")
	var addDuration string
	flag.String(&addDuration,"duration","","duration of the task in hh:mm:ss")
	var helpFlag bool
	flag.Bool(*helpFlag,"help",false,"show help")
	flag.bool(*helpFlag,"h",false,"show help(short)")
	
	//parse the flags
	flag.Parse()

	if  helpFlag== true{
		controller.ShowHelp()
		return
	}
	if version==true {
		controller.CheckVersion()
		return 
	}
	
	if *addTaskName!=""{
		if *addDuration!=""{
			tasks.AddTask(*addTaskName,*addDuration)
		}else{
			tasks.AddTask(*addTaskName,"00:10:00")
		}
		return
	}

}


