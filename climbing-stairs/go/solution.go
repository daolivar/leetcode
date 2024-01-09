// climbStairs uses dynamic programming to calculate the result by storing intermediate
// results in a slice (`dp`). The slice is iteratively filled with the number of ways
// to reach each step, starting from the base cases for steps 0 and 1.
//
// Time Complexity: O(n), this is because the function uses a bottom-up dynamic
// programming approach (tabulation) to fill the array dp from the base cases
// (0  and 1) up to n. The loop iterates n times, and at each iteration, a
// constant amount of work is done.
//
// Space Complexity: O(n), the array dp of size n + 1 is used to store the solutions to
// subproblems. Therefore, the space required is linearly proportional to the input n.
//
// Parameters:
//   - n: The total number of steps in the staircase.
//
// Returns:
//   - int: The number of distinct ways to climb to the top.
func climbStairs(n int) int {
	// Initialize a slice to store intermediate results for dynamic programming
	dp := make([]int, n+1)

	// Base cases: there is 1 way to reach step 0 and 1 way to reach step 1
	dp[0] = 1
	dp[1] = 1

	// Fill the dp slice iteratively for steps 2 to n
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// Return the number of distinct ways to climb to the top
	return dp[n]
}
