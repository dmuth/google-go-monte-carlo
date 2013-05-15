
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

	monte := monte.New(config.Size, config.Num_points, 
		config.Num_goroutines)
	pi := monte.Main(config)

	fmt.Println("Pi is:", pi)

} // End of main()



