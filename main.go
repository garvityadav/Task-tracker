package main

import(
	"flag"
	"github.com/garvityadav/Task-tracker/controller"
)


func main(){
	//define a flag
	
	//version Flag
	var version bool
	flag.BoolVar(&version,"version",false,"show version")	
	flag.BoolVar(&version,"v",false,"show version (short)")

	//Task flag
	var addTaskName string
	flag.StringVar(&addTaskName,"add-task","","Name of the task")
	var addDuration string
	flag.StringVar(&addDuration,"duration","","duration of the task in hh:mm:ss")
	var helpFlag bool
	flag.BoolVar(&helpFlag,"help",false,"show help")
	flag.BoolVar(&helpFlag,"h",false,"show help(short)")
	
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
	
	if addTaskName!=""{
		if addDuration!=""{
			controller.AddTask(addTaskName,addDuration)
		}else{
			controller.AddTask(addTaskName,"00:10:00")
		}
		return
	}

}


