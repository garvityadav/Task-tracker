package main

import( "fmt"
	"flag")


func main(){
	fmt.Println("this is task tracker")
	//define a flag
	nameFlag:= flag.String("name","World","Your name")
	add := flag.String("add"
	//parse the flags
	flag.Parse()
	//use the value
	fmt.Printf("Hello,%s!\n",*name)

}
