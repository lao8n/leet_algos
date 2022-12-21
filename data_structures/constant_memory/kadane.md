Kadane's algorithm
* dynamic programming algorithm where it calculates rolling sum, and if sum > max set max to sum otherwise set sum = 0
* time complexity O(n^2) space complexity O(1)

```
func kadane(int[] input) int {
  max, sum := 0, 0
  for i:=1; i < len(input); i++ {
    sum += input[i] - input[i-1];
    if sum > max {
        max = sum
    } else if sum < 0 {
        sum = 0
    }
  }
  return max
}
```