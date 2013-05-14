

package random_md5

import "bytes"
import "crypto/md5"
import "encoding/binary"
//import "fmt"


type random_struct struct {
	//
	// The results of the previous hash.
	//
	seed string
	//
	// Our bitmask for pulling values from MD5 responses
	//
	bitmask uint64
}


/**
* Return a random number between 1 and n
* @param {integer} max The maximum random number.
* @return {integer} retval The random value
*/
func (r *random_struct) intn(max uint64) (retval uint64) {

	//
	// Create a hash based on our current seed
	//
	hash := md5.New()
	hash.Write([]byte(r.seed))
	md5_value := hash.Sum(nil)
	//fmt.Println("MD5:", fmt.Sprintf("%x", md5_value))

	//
	// Make the current hash our new seed
	//
	r.seed = string(md5_value)

	//
	// Get our bitmask if we don't already have it
	//
	if (r.bitmask == 0) {
		r.bitmask = getBitmask(max)
	}

	//
	// Grab 8 bytes and put them into a uint64
	//
	buf := bytes.NewBuffer(md5_value)
	binary.Read(buf, binary.LittleEndian, &retval)
	retval = retval & r.bitmask
	//fmt.Println(md5_value, retval)

	//
	// If the value is too big (e.g. 32 when the max is 17), call ourself
	// again and hope we get lucky.
	// (And I hope this never causes a stack overflow...)
	//
	if (retval >= max) {
		retval = r.intn(max)
	}

	return(retval)

} // End of intn()


/**
* Read a request off of a channel, generate a random value, and write 
* it back out.
*
* @param {chan int} in Our channel to read in requests. Each value 
*	read is an array of the maximum random number and the number of 
*	random numbers we want.
* @param {chan int} out The channel to write results out to in groups of 2
*/
func (r *random_struct) intNChannel(in chan []uint64, out chan [][]uint64) {

	var values []uint64
	var retval [][]uint64

	for {

		args := <- in
		max := args[0]
		num_random := args[1]

		for i:=uint64(0); i<num_random; i++ {
			num := r.intn(max)
			values = append(values, num)
			//fmt.Println("num_random, max, result, values:", num_random, max, num, values)
			if (len(values) == 2) {
				retval = append(retval, values)
				values = []uint64{}
			}
		}

		out <- retval
		retval = [][]uint64{}

	}

} // End of intNChannel()


/**
* Create a bitmask from our max value.  This is for extracting that 
* value from an MD5 hash.
*
* @param {uint64} max Our maximum random value
* @return {uint64} A value which is 2*n-1.
*/
func getBitmask(max uint64) (retval uint64) {

	retval = 1
	for i:=1; i<64; i++ {
		retval *= 2
		if (retval >= max) {
			break
		}
	}

	retval--

	return(retval)

} // End of getBitmask()


