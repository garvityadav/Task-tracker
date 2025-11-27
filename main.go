package main

import( "fmt"
	"flag")


func main(){
	fmt.Println("this is task tracker")
	//define a flag
	nameFlag:= flag.String("name","World","Your name")

	add := flag.String("add-task","new task","Add your task name")
	description := flag.String("description","description","add description")
	setTime := flag.String("set-time","1****","set time in cron expression")
	//parse the flags
	flag.Parse()
	//use the value
	fmt.Printf("Hello,%s!\n",*name)

}
