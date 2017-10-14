package scope

import "fmt"

// This is the model so only can be seen from the controller
// all the funcname are lowercase
// this function and/ or vars are visible in "package scope" files
// our main entry point of entry/file is "package main"
// where we are importing the package scope with "import" not with "package"
// so "package scope" will see it
// "import scope" only will see function and vars that:
// first capital letter
func addmodel(dataModel string) {
	fmt.Println("Hello", dataModel)
}
