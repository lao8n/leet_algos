package data_structures

import "math"

// sieve of eratosthenes -> only need to go to sqrt of n
// approaches
// * for loop over all possible integers - each integer need to check all numbers less than square root...O(n^2)
// * loop from i to sqrt n -> multiply all possible candidates - remove from set costs O(1) - still O(n^2) to check everything
// * loop through numbers from 1 to n-1 - but just need to find one divisor -> work up from 2 onwards? i guess factorial costs time?
// * solution
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	numbers := make([]bool, n)
	for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
		// if it is a prime
		if !numbers[p] {
			// multiples smaller than p * p are covered by other primes
			for m := p * p; m < n; m += p {
				numbers[m] = true
			}
		}
	}

	numPrimes := 0
	for i := 2; i < n; i++ {
		if !numbers[i] {
			numPrimes++
		}
	}
	return numPrimes
}
