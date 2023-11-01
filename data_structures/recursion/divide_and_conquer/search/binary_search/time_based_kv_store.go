package data_structures

// clarifying questions
// * multiple values for same key 
// * timestamp is int
// approaches
// * map with a list of items - should always be ordered so just pick the last one?
type TimeMap struct {
    m map[string][]entry
}

type entry struct{
    val string
    timestamp int
}

func Constructor() TimeMap {
    return TimeMap{
        m: make(map[string][]entry),
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.m[key] = append(this.m[key], entry{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if vals, ok := this.m[key]; ok {
        return this.binarySearch(vals, timestamp)
    }
    return ""
}

func (this *TimeMap) binarySearch(vals []entry, timestamp int) string {
    left, right := 0, len(vals) - 1
    result := ""
    for left <= right {
        mid := left + (right - left) / 2
        // we want the largest possible ts <= timestamp
        // if mid > we know we can discard right hand side
        if vals[mid].timestamp > timestamp {
            right = mid - 1
        // we might have overshot
        } else {
            result = vals[mid].val
            left = mid + 1
        }
    }
    return result
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */