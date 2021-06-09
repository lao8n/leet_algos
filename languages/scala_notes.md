Scala coding choices
* greedy vs dynamic programming
* O(log) by replacing for with recursion
* Can we sort (or is it already sorted) e.g. two indices
* Can we use memory to save time e.g. Maps
* Mutable data  structures or recursion with immutable data structures
* Do we need all the solutions or the first solution (if former need accumulator)
* Rather than foldleft is there a way to aggregate with map and group by instead?

Scala language
`:::` vs `::` vs `++` vs `add`

`implicit ord: Ordering[T]`

`T <: Ordered[T]`

`xs.foldstart`

`tail.partition`

`nums.zipWithIndex.foldLeft(Map[Int, Int]()){ case (numsMap, numsTuple @(x, complementIndex))` zipWithIndex gives Array of tuples where each tuple is nums and index. foldLeft then initializes an empty map. and we pattern match on map and tuple 

`foldLeft` vs `foldRight` does foldLeft only work with mutable data structures? what is it equivalent to?

`numsTuple @(x, complementIndex)` a way to name the tuple 

`scala.collection.Mutable`

`scala.collection.Immutable`

`numsMap.find(_._ == complement)` vs `numsMap.get(complement)` 

`.sorted` vs `sortWith(_>_)` 

`numsMap + (num -> numIndex)` is deprecated

`new ArrayList<>(setOfSomething);` convert set to list