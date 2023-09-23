package data_structures

func canAttendMeetings(intervals [][]int) bool {
    // sort by start time and end-time
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] { // start time same
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    for i := 1; i < len(intervals); i++ {
        if intervals[i-1][1] > intervals[i][0] { // previous meeting ends after current meeting starts
            return false
        }
    }
    return true
}