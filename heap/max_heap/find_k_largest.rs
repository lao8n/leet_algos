use std::collections::BinaryHeap; 
// max heap is default in binary collections
impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut heap = BinaryHeap::from(nums);
        let mut counter = k;
        while counter > 1 {
            heap.pop();
            counter -= 1;
        }
        *heap.peek().unwrap()
    }
}