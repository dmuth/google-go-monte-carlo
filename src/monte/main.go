
/**
* This package solves for Pi using the Monte Carlo method.
* Instead of creating an entire circle, it only creates the 
* upper-right quadrant.  This makes working with the 
* numbers (all >= 0) easier. :-)
*/
package monte

//import "fmt"
import "math"

import "../args"
import "../random"
import "../random_md5"


//
// Our data structure
//
type monte struct {
	size uint64 // Size of the grid we're creating
	size_squared int64 // Size squared, for checking with the Pythagorean thereom
	num_points int
	num_points_left int
	num_goroutines int
	num_points_in_circle int
	num_points_not_in_circle int
}


/**
* Create a new instance of our data structure.
* 
* @param {int} size How big is our grid for Monte Carlo?
* @param {int} num_numbers How many points to do we want to generate?
* @param {int} num_goroutines How many goroutines do we want to use 
*	for generating random numbers?
* 
* @return {monte} Our structure
*/
func New(size uint64, num_points int, num_goroutines int) (monte) {

	size_squared := math.Pow(float64(size), 2)
	retval := monte{size, int64(size_squared), num_points, 
		num_points, num_goroutines, 0, 0}

	return retval

} // End of New()


/**
* Our main entry point.
*/
func (m monte) Main(config args.Config) float64 {

	out_check_points := make(chan []uint64)
	pi := make(chan float64)


	//
	// Goroutine to create points from random numbers
	//
	go m.getPoints(out_check_points, pi)

	//
	// Start generating our points!
	//
	num_numbers := m.num_points * 2;
	if (!config.Random_md5) {
		random.IntnBackground(out_check_points, m.size, num_numbers, 
			m.num_goroutines)
	} else {
		random_md5.IntnBackground(out_check_points, m.size, num_numbers, 
			m.num_goroutines)
	}

	//
	// Read our value of Pi when we're all done!
	//
	retval := <- pi

	return(retval)

} // End of Main()


/**
* Grab random numbers 2 at a time and pass them into our channel for checking.
* @param {chan} in Inbound channel which feeds us random numbers.
* @param {chan} out Outbound channel which takes an array of two points.
*/
func (m *monte) getPoints(in chan []uint64, out chan float64) {

	for {
		values := <- in
		x := values[0]
		y := values[1]

		x2 := math.Pow(float64(x), 2)
		y2 := math.Pow(float64(y), 2)
		c := int64(x2 + y2)

		if (c <= m.size_squared) {
			m.num_points_in_circle++
		} else {
			m.num_points_not_in_circle++
		}

		m.num_points_left--
		if (m.num_points_left == 0) {
			pi := m.calculatePi()
			out <- pi
		}

	}

}


/**
* Calculate Pi based on our points in or out of the circle
*
* @return {float64} The value of Pi
*/
func (m *monte) calculatePi() (float64) {

	total := m.num_points_in_circle + m.num_points_not_in_circle
	retval := ( float64(m.num_points_in_circle) / float64(total) ) * 4
	return(retval)

} // End of calculatePIi()



