
package main

//
// Load Google Go packages
//
import "fmt"
import "runtime"


//
// Load packages local to this project
//
import "./src/args"
import "./src/monte"


/**
* Our main entry point.
*/
func main() {

	//
	// Crank up the number of processors used
	//
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	config := args.ParseArgs()

	/* TODO:
	- modify messages for MD5 to be [max, num random numbers]
	- param for chunk size
	- modify messages for random to [max, num random numbers]
		- caching of point matches
	- change function to calculate Pi to do it manually when we're done
	*/

	monte := monte.New(config.Size, config.Num_points, config.Num_goroutines)
	pi := monte.Main(config)

	fmt.Println("Pi is:", pi)

} // End of main()



