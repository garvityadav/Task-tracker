package controller
import "fmt"
// create a task

func AddTask(taskName ,time string)error{
	//this is used to create a new task
	fmt.Printf(taskName+" & "+time+"\n")

	return nil
}


