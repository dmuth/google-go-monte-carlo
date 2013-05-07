
/**
* This package is for parsing command line arguments.
*/
package args;

import "fmt"
import "strconv"
import "flag"
import "os"

/**
* Print our syntax diagram and exit.
*/
func PrintSyntax() {
	fmt.Printf(
		"Syntax: %s --size n --num-points n --num-cores n\n", 
		os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}


/**
* Loop through our arguments and parse them.
* @return {map} A map of our array values
*/
func ParseArgs() (map[string]string) {

	var params = make(map[string]string)
	params["size"] = strconv.FormatInt(10, 10)
	params["num_points"] = strconv.FormatInt(100, 10)
	params["num_cores"] = strconv.FormatInt(2, 10)

	for i:=1; i<len(os.Args); i++ {

		var arg = os.Args[i]
		var arg_next string = ""
		index_next := i+1
		if (index_next < len(os.Args) && os.Args[index_next] != "") {
			arg_next = os.Args[i+1]
		}

		//
		// If the user asked for help, bail
		//
		if arg == "-h" {
			PrintSyntax()

		} else if arg == "--size" {
			params["size"] = arg_next
			i++

		} else if arg == "--num-points" {
			params["num_points"] = arg_next
			i++

		} else if arg == "--num-cores" {
			params["num_cores"] = arg_next
			i++

		} else {
			PrintSyntax()

		}

	}

	return(params)

} // End of ParseArgs()


