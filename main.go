package main

import( "fmt"
	"flag")


func main(){
	fmt.Println("this is task tracker")
	//define a flag
	nameFlag:= flag.String("name","World","Your name")

	addFlag := flag.String("add-task","new task","Add your task name")
	 descriptionFlag := flag.String("description","description","add description")
	setTimeFlag := flag.String("set-time","1****","set time in cron expression")
	listFlag := flag.String("list","list all jobs")
	listRunningFlag := flag.String("running","list all running jobs")


	//parse the flags
	flag.Parse()
	
	//use the value
	
	if flag.Lookup("add-task").Value.String()!="new task"  {
		fmt.Printf("you added : %s \n",*addFlag)

	}
	if flag.Lookup("name").Value.String()!="World"{
	fmt.Printf("Hello,%s!\n",*nameFlag)
	}
	
	if flag.Lookup("description").Value.String()!="description"{
		fmt.Printf("Discription : ,%s!\n",*descriptionFlag)
	}


}

// create 
func addTask()

