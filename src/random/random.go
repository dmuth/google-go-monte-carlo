
package random

import "math/rand"
import "time"
//import "log"


//
// Have we seeded the random number generator?
//
var seeded bool = false;


/**
* Return a random number between 1 and n
* @param {integer} n The maximum random number.
* @return {integer} retval The random value
*/
func Intn(n int) (retval int) {

	if (!seeded) {
		rand.Seed(time.Now().UnixNano())
	}

	retval = rand.Intn(n)
	return(retval)

} // End of Intn()


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
		go intNChannel(in, out)
	}

	//
	// Now stuff our input channel with all of our requests.
	//
	for i := 0; i<num_numbers; i++ {
		in <- max
	}

} // End of IntnBackground()


/**
* Read a request off of a channel, generate a random value, and write 
* it back out.
*
* @param {chan int} in Our channel to read in requests. Each value 
*	read is the maximum random number.
* @param {chan int} out The channel to write results out to.
*/
func intNChannel(in chan int, out chan int) {

	//timeout := time.Millisecond * 100

	//
	// This will loop forever
	//
	for {

		select {
			case max := <-in:
				i := Intn(max)
				out <- i

			//case <-time.After(timeout):
			//	log.Println("TIMEOUT")

		}

	}

} // End of intNChannel()



