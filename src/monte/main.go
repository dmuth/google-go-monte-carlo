
/**
* This package solves for Pi using the Monte Carlo method.
* Instead of creating an entire circle, it only creates the 
* upper-right quadrant.  This makes working with the 
* numbers (all >= 0) easier. :-)
*/
package monte

//import "log"
import "math"

import "../random"


//
// Our data structure
//
type monte struct {
	size int // Size of the grid we're creating
	size_squared int64 // Size squared, for checking with the Pythagorean thereom
	num_points int
	num_points_left int
	num_cores int
	num_points_in_circle int
	num_points_not_in_circle int
}


/**
* Create a new instance of our data structure.
* 
* @param {int} size How big is our grid for Monte Carlo?
* @param {int} num_numbers How many points to do we want to generate?
* @param {int} num_cores How many cores do we want to use for each major task?
* 
* @return {monte} Our structure
*/
func New(size int, num_points int, num_cores int) (monte) {

	size_squared := math.Pow(float64(size), 2)
	retval := monte{size, int64(size_squared), num_points, num_points, num_cores, 0, 0}

	return retval

} // End of New()


/**
* Our main entry point.
*/
func (m monte) Main() float64 {

	out_check_points := make(chan int)
	in_check_points := make(chan []int)
	in_calculate_pi := make(chan bool)
	out := make(chan float64)


	//
	// Goroutine to create get points from random numbers
	//
	go m.getPoints(out_check_points, in_check_points)

	//
	// Goroutine to check each point and determine if it's in the
	// circle or not.
	//
	go m.checkPoints(in_check_points, in_calculate_pi)

	//
	// Goroutine to calculate our value of Pi
	//
	go m.calculatePi(in_calculate_pi, out)

	//
	// Start generating our points!
	//
	num_numbers := m.num_points * 2;
	random.IntnBackground(out_check_points, m.size, num_numbers, m.num_cores)

	//
	// Read our value of Pi when we're all done!
	//
	retval := <- out
	return(retval)

} // End of Main()


/**
* Grab random numbers 2 at a time and pass them into our channel for checking.
* @param {chan} in Inbound channel which feeds us random numbers.
* @param {chan} out Outbound channel which takes an array of two points.
*/
func (m monte) getPoints(in chan int, out chan []int) {

	for {
		x := <- in
		y := <- in

		var points []int
		points = append(points, x)
		points = append(points, y)
		out <- points;

	}

}


/**
* Continuously read sets of points from our channel and check to see if 
*	each is in the circle.
*
* @param {chan} in Our channel that we're reading sets of points from
* @param {chan} out The channel we're writing to which signals calculatePi() 
*	to perform the Pi calculation.
*
*/
func (m *monte) checkPoints(in chan []int, out chan bool) {

	for {
		points := <- in
		x := points[0]
		y := points[1]
		x2 := math.Pow(float64(x), 2)
		y2 := math.Pow(float64(y), 2)
		c := int64(x2 + y2)

		if (c <= m.size_squared) {
			m.num_points_in_circle++
		} else {
			m.num_points_not_in_circle++
		}

		out <- true

	}

} // End of checkPoints()


/**
* Calculate Pi from our set of points.
*
* @param {chan} in Useless booleans are read from this. 
*	On each read, we know that another set of points 
*	has been generated.
* @param {chan} out When we're all done, send our value of Pi here.
*/
func (m *monte) calculatePi(in chan bool, out chan float64) {

	for {

		c := <- in
		if (c) {}
		m.num_points_left--

		//log.Println("PI 2: ", pi)

		//
		// Do we have all of our points? Send our result out then.
		//
		if (m.num_points_left == 0) {
			total := m.num_points_in_circle + m.num_points_not_in_circle
			pi := ( float64(m.num_points_in_circle) / float64(total) ) * 4
			out <- pi
		}

	}

} // End of calculatePIi()



