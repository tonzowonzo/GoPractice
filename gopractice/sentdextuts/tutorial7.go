// Methods.
/* Two method types in Go: value receivers and pointer receivers.
 - Value receivers - receive values and do calculations on them.
 - Pointer receivers - Allows you to change the value in the struct.
*/
// Instantiate package.
package main

// Import libraries.
import ("fmt")

 // Add constants.
 const usixteenbitmax float64 = 65535
 const kmh_multiple float64 = 1.60934

// Instantiate car class.
 type car struct {
 	gas_pedal uint16
 	brake_pedal uint16
 	steering_wheel int16
 	top_speed_kmh float64
 }

// Method to calculate kmh, a value receiver. We can reference the car struct with c.
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

// Function to calculate mph.
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

// Main function.
func main() {
	a_car := car{gas_pedal: 22341,
				 brake_pedal: 0,
				 steering_wheel: 12561,
				 top_speed_kmh: 225.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}

