`:::` 
`::`
`implicit ord: Ordering[T]`
`T <: Ordered[T]`
`xs.foldstart`
`tail.partition`
`nums.zipWithIndex.foldLeft(Map[Int, Int]()){ case (numsMap, numsTuple @(x, complementIndex))` zipWithIndex gives Array of tuples where each tuple is nums and index. foldLeft then initializes an empty map. and we pattern match on map and tuple  
`foldLeft` vs `foldRight` 
`numsTuple @(x, complementIndex)` a way to name the tuple 
`scala.collection.Mutable`
`scala.collection.Immutable`