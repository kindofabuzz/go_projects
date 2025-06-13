package packageone

import "fmt"

var PackageVar = 3

func PrintMe(i1, i2 int) {
	fmt.Println(i1, i2, PackageVar)
}
