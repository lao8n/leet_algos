package data_structures

// clarifying questions

// approaches
// * use recursion to acc a path - because we have slices we want to backtrack for the path

// specifics
// * may need to copy slice when putting it in the final solution
// * how keep track of remaining letters? typically we have an index, but here we need to have a remaining letters
func permute(nums []int) [][]int {
    d := data{
        nums: nums,
        solution: [][]int{},
    }
    d.backtrack([]int{}, map[int]bool{})
    return d.solution
}

type data struct {
    nums []int
    solution [][]int
}

func (d *data) backtrack(path []int, used map[int]bool) {
    // Base case
    if len(path) == len(d.nums) {
        // Create a copy of path to avoid slice reference issues
        pathCopy := make([]int, len(path))
        copy(pathCopy, path)
        d.solution = append(d.solution, pathCopy)
        return
    }
    
    // Recursive case
    for _, num := range d.nums {
        if _, ok := used[num]; !ok {
            used[num] = true
            path = append(path, num)
            d.backtrack(path, used)
            delete(used, num)
            path = path[:len(path) - 1]
        }
    }
}