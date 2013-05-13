
package random

//import "fmt"
import "math/rand"
import "time"


//
// Okay, this structure is a bit silly since there is a global seed 
// for the random number generator.  I did this mostly as an exercise 
// so that I could more easily implement this in the MD5 version.  
// The performance impact is negligible, seeing that rand.Seed() 
// is only called the number of goroutines we're using.
//
type random_struct struct {
	//
	// Have we seeded the random number generator?
	//
	seeded bool
}


/**
* Read a request off of a channel, generate a random value, and write 
* it back out.
*
* @param {chan int} in Our channel to read in requests. Each value 
*	read is the maximum random number.
* @param {chan int} out The channel to write results out to.
*/
func (r random_struct) intNChannel(in chan uint64, out chan uint64) {

	//timeout := time.Millisecond * 100

	//
	// This will loop forever
	//
	for {

		select {
			case max := <-in:
				i := r.intn(max)
				out <- i

			//case <-time.After(timeout):
			//	log.Println("TIMEOUT")

		}

	}

} // End of intNChannel()


/**
* Return a random number between 1 and n
* @param {integer} n The maximum random number.
* @return {integer} retval The random value
*/
func (r *random_struct) intn(n uint64) (retval uint64) {

	if (!r.seeded) {
		rand.Seed(time.Now().UnixNano())
		r.seeded = true
	}

	retval = uint64(rand.Int63n(int64(n)))
	return(retval)

} // End of Intn()


