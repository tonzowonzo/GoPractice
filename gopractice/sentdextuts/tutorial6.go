// Structs (similar to classes).
package main 

// Imports.
import ("fmt")

// Create a struct (class).
type car struct {
	gas_pedal uint16 // min = 0, max = 65535.
	brake_pedal uint16
	steering_wheel int16 // min = -32k, max=+32k
	top_speed_kmh float64
}





// Main function.
func main() {
	// Instantiate a struct.
	a_car := car{gas_pedal: 22314, 
				 brake_pedal: 0, 
				 steering_wheel: 12561, 
				 top_speed_kmh: 225.0}
	// Reference values from the struct.
	fmt.Println(a_car.gas_pedal)
}