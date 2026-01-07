package controller
import "fmt"

// check health 
func CheckHealth()error{
	return nil
	}

//check version
func CheckVersion()error{
// command : tasker --version
	// this command should print the current version of the application. 
	// --version flag is used or in short -v.
	fmt.Printf("Tasker\nversion: v0.0.1\n")	
return nil
}
