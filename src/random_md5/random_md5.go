
package random_md5

//import "fmt"
import "strconv"
import "time"


/**
* Create background processes and pass out channels into them.
*
* @param {channel} out Where our random values will be written to
* @param {int} max All random numbers are < this number.
* @param {int} num_numbers How many random numbers to generate?
* @param {int} num_goroutines How many goroutines to use?
*/
func IntnBackground(out chan []uint64, max uint64, num_numbers int, num_goroutines int) {

	in := make(chan []uint64)

	//
	// Create a number of background processes.
	//
	for i := 0; i < num_goroutines; i++ {
		seed := strconv.FormatInt(time.Now().UnixNano(), 10)
		random_struct := random_struct{seed, 0}
		go random_struct.intNChannel(in, out)
	}

	//
	// Now stuff our input channel with all of our requests.
	// We'll do it in chunks so as not to destroy our CPUs.
	// See my blog post at http://www.dmuth.org/node/1414/multi-core-cpu-performance-google-go
	// for a further explanation.
	//
	num_left := uint64(num_numbers)
	chunk_size := uint64(10000)

	for {
		if (num_left < chunk_size) {
			chunk_size = num_left
		}

		num_left -= chunk_size

		var args []uint64
		args = append(args, max, chunk_size)
		in <- args

		if (num_left <= 0) {
			break
		}

	}

} // End of IntnBackground()


