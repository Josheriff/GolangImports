package main

import "jose/scopelesson/scope"

func main() {
	scope.Addmodel("trolololo")
	// we can see scope.Addmodel because the function starts in capital
	// the MVC ...
	// we only has uppercase in the controller that can see everything
	// we put lowercase in the model, only accesible by the controller

}

// mlab connection:
// user:gopher
// password: gogogo
// mongodb://gopher:gogogo@ds119395.mlab.com:19395/tryinggolang
