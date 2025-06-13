package packageone

import "fmt"

var PackageVar = 3

func instructions() {
	// variables
	// declare a package level variable for the main
	// package named myVar

	// declare a block level variable for the main function
	// called blockVar

	// declare a package level variable in the packageone
	// package named PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function, print out the values of myVar,
	// blockVar, and PackageVar on one line using the PrintMe
	// function in packageone.

}

func PrintMe(i1, i2 int) {
	fmt.Println(i1, i2, PackageVar)
}
