
/**
* This package is for parsing command line arguments.
*/
package args;

import "flag"
import "fmt"
import "os"


/**
* Our config structure to hold whatever we parse
*/
type Config struct {
	Size uint64
	Num_points int
	Num_goroutines int
	Random_md5 bool
	Chunk_size uint64
}


/**
* Loop through our arguments and parse them.
* @return {map} A map of our array values
*/
func ParseArgs() (Config) {

	config := Config{}

	flag.Uint64Var(&config.Size, "size", 10, "How big to make the grid for the circle quadrant")
	flag.IntVar(&config.Num_points, "num-points", 10, "How many points to plot?")
	flag.IntVar(&config.Num_goroutines, "num-goroutines", 10, "How many goroutines to use for generating random numbers")
	flag.BoolVar(&config.Random_md5, "random-md5", false, "Set to use MD5 for faux random number generation")
	flag.Uint64Var(&config.Chunk_size, "chunk-size", 10000, "Create random numbers in batches of this size (must be a multiple of 2)")
	help := flag.Bool("help", false, "test2")
	h := flag.Bool("h", false, "To get this help")
	flag.Parse()

	if (config.Chunk_size % 2 != 0) {
		fmt.Println("chunk-size must be a multiple of 2!")
		os.Exit(1)
	}

	if (*help || *h) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return(config)

} // End of ParseArgs()


