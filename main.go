
package main

//
// Load Google Go packages
//
import "fmt"
import "strconv"
//import "time"

//
// Load packages local to this project
//
import "./src/args"
import "./src/monte"


/**
* Our main entry point.
*/
func main() {

	params := args.ParseArgs()
	fmt.Printf("Params: %s\n", params)

	size, _ := strconv.Atoi(params["size"])
	num_points, _ := strconv.Atoi(params["num_points"])
	num_cores, _ := strconv.Atoi(params["num_cores"])

	/* TODO:
		- caching of point matches
	*/

	monte := monte.New(size, num_points, num_cores)
	pi := monte.Main()

	fmt.Println("Pi is:", pi)

} // End of main()



