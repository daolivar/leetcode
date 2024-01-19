/**
 * Solution class provides a method to convert temperature from Celsius to
 * Kelvin and Fahrenheit.
 */
class Solution {
    /**
     * Converts temperature from Celsius to Kelvin and Fahrenheit.
     * 
     * Time Complexity: O(1), as the conversion involves basic arithmetic operations
     * that are constant time.
     * 
     * Space Complexity: O(1), as there is no additional space used that grows with
     * the input.
     * 
     * @param celsius The temperature in Celsius.
     * @return An array containing the converted temperature values in Kelvin and
     *         Fahrenheit.
     *         Index 0: Temperature in Kelvin.
     *         Index 1: Temperature in Fahrenheit.
     */
    public double[] convertTemperature(double celsius) {
        return new double[] { celsius + 273.15, celsius * 1.80 + 32.00 };
    }
}
