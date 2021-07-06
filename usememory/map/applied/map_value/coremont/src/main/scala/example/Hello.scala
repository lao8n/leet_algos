package example
// to set up sbt package run 
// sbt new sbt/scala-seed.g8

object Solution extends App {
  // brute force? 
  // greedy approach?   
  // dynamic programming?

  // go through each line and either calculate a number maybe multiplying
  // by position in alphabet or create a histogram 
  // or i could just sort but that costs O(n logn)

  // immutable variables or mutable?
  def anagram(reader : Iterator[String], anagramMap : Map[String, List[String]]): List[String]= {
    if(reader.hasNext){
      val line = reader.next()
      // process anagram 
      val sortedLine = line.sorted
      anagramMap get sortedLine match {
        case Some(listAnagrams) => anagram(reader, anagramMap + (sortedLine -> (line :: listAnagrams)))
        case None => anagram(reader, anagramMap + (sortedLine -> List(line)))
      }
    }
    else {
      // base case return 
      // [["ab", "ba"], ["ac"]]
      // 
      anagramMap.values.toList.maxBy(_.size)
    }
  }

  def anagram2(reader: Iterator[String]) : List[String] = {
    // read lines
    val words = reader.toList
    // List[String] 
    // Map[sorted, [words]]
    // [[words1], [words2]]
    words.groupBy(_.sorted).values.maxBy(_.size)
  }
}
