package main

import( "fmt"
	"flag")


func main(){
	fmt.Println("this is task tracker")
	//define a flag
	nameFlag:= flag.String("name","World","Your name")

	addFlag := flag.String("add-task","new task","Add your task name")
	/* descriptionFlag := flag.String("description","description","add description")
	setTimeFlag := flag.String("set-time","1****","set time in cron expression")
*/
	//parse the flags
	flag.Parse()
	//use the value
	
	if flag.Lookup("add-task").Value.String()!="new task"  {
		fmt.Printf("you added : %s \n",*addFlag)

	}else if flag.Lookup("name").Value.String()!="World"{
	fmt.Printf("Hello,%s!\n",*nameFlag)
	}else{
	fmt.Printf("Hello,%s!\n",*nameFlag)
	}
}
