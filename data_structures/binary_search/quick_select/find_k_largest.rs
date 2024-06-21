impl Solution {
    // had to make nums mut o/w would need to clone
    pub fn find_kth_largest(mut nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        *nums.select_nth_unstable(n - k as usize).1
    }
}