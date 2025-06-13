package main

import (
	"myapp/packageone"
)

var myVar = 1

func main() {
	var blockVar = 2
	packageone.PrintMe(myVar, blockVar)
}
