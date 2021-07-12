object Solution {
    def merge(intervals: Array[Array[Int]]): Array[Array[Int]] = {
        intervals.toList.sortBy(arr => arr(0)).foldLeft(List[Array[Int]]()){
            // guaranteed that last.start <= current.start
            // need to test if there is overlap i.e. current.start <= last.end // we prepend head
            // and we need to check that last.end < current.end
            case (head :: tailIntervals, interval) if interval(0) <= head(1) && head(0) <= interval(0)=> 
                Array(head(0), head(1) max interval(1)) :: tailIntervals
            // new interval is needed because no overlap
            case (listIntervals, interval) =>
                interval :: listIntervals
        }.toArray
    }
}
