// Echo prints program name and command-line arguments
package main

import (
	"os"

	"main/echo"
)

func main() {
	// for i, arg := range os.Args {
	// 	fmt.Printf("%d: %s\n", i, arg)
	// }

	echo.EchoInefficient(os.Args)
}
