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
	if *addFlag {
		fmt.Printf("you added : %s \n",*addFlag)
	}
	fmt.Printf("Hello,%s!\n",*nameFlag)

}
