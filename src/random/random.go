
package random


/**
* Create background processes and pass out channels into them.
*
* @param {channel} out Where our random values will be written to
* @param {int} max All random numbers are < this number.
* @param {int} num_numbers How many random numbers to generate?
* @param {int} num_goroutines How many goroutines to use?
*/
func IntnBackground(out chan int, max int, num_numbers int, num_goroutines int) {

	in := make(chan int)

	//
	// Create a number of background processes.
	//
	// Originally, I thought that this would get me lots of parallelism when
	// generating random numbers.  As it turns out, it did not.  I can
	// have *10* goroutines, and only one CPU core will spike to 100%.
	// So I learned something new about how random numbers are generated.
	//
	for i := 0; i < num_goroutines; i++ {
		random_struct := random_struct{false}
		go random_struct.intNChannel(in, out)
	}

	//
	// Now stuff our input channel with all of our requests.
	//
	for i := 0; i<num_numbers; i++ {
		in <- max
	}

} // End of IntnBackground()



