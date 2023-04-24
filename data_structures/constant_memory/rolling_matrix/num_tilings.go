package data_structures

// interconnected problems combined for a solution - esp as num of ways -> dynamic programming

// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// try bottom up - as potential to optimise space into perhaps just the previous column

// 3 components to dp
// 1 base cases
// 2 recurrence
// 3 state -> as we don't care about horizontal symmetry just need to know 1. num of n left 2. whether 
// 1 or 2 spaces to fill

// choose one space way 
// [] + [][] -> [][][] 
// []   []      [][]
// or
// [] +      -> []
//      [][]    [][]
// oneFilledWays[0] = (oneFilledWays[1] + twoFilledWays[2]) % mod
// choose two space ways
// [] + [] -> []
// []   []    []
// or
// [] + [][] -> [][][]
// []   [][]    [][][]
// or 
// [] +   [] -> [][] * 2
//      [][]    [][]
// twoFilledWays[0] = (oneFilledWays[1] * 2 + twoFilledWays[1] + twoFilledWays[2]) % mod

// we also update
// oneFilledWays[1], oneFilledWays[2] = oneFilledWays[0], oneFilledWays[1]
// twoFilledWays[1], twoFilledWays[2] = twoFilledWays[0], twoFilledWays[1]		

func numTilings(n int) int {
    // base cases
    if n <= 2{
        return n
    }

    mod := 1000000007
    // we can represent this in a 3 x 3 transition matrix
    // twoFilledWays[k]   = 1 x twoFilledWays[k-1] + 1 x twoFilledWays[k-2] + 2 x oneFilledWays[k-1]
    // twoFilledWays[k-1] = 1 x twoFilledWays[k-1] 
    // oneFilledWays[k]   =                          1 x twoFilledWays[k-2] + 1 x oneFilledWays[k-1]
    // note we can simplify by the fact that we also update 
    base_case := [][]int {
        []int{2},
        []int{1},
        []int{1},
    }
    matrix := [][]int{
        []int{1, 1, 2},
        []int{1, 0, 0},
        []int{0, 1, 1},
    }

    matrix_exp := matrix
    for i := 0; i < n - 3; i++ {
        matrix_exp = matrix_multiply(matrix_exp, matrix, mod)
    }
    result := matrix_multiply(matrix_exp, base_case, mod)
    return result[0][0]
}

func matrix_multiply(x, y [][]int, mod int) [][]int {
    rows, cols, shared := len(x), len(y[0]), len(x[0])
    result := make([][]int, len(x))
    for i := 0; i < rows; i++ {
        result[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            sum := 0
            for k := 0; k < shared; k++ {
                sum = (sum + x[i][k] * y[k][j]) % mod
            } 
            result[i][j] = sum
        }
    }
    return result
}
Console
