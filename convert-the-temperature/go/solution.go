// convertTemperature converts a temperature value from Celsius to Kelvin and Fahrenheit.
//
// Time Complexity: O(1) as the conversion involves simple arithmetic operations.
//
// Space Complexity: O(1) as the result is returned in a fixed-size slice.
//
// Parameters:
//   - celsius: The temperature value in Celsius.
//
// Returns:
//   - A slice containing the converted temperature values in Kelvin and Fahrenheit.
//     Index 0: Temperature in Kelvin
//     Index 1: Temperature in Fahrenheit
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}
