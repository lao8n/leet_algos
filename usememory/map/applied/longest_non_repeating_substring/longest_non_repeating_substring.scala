object Solution {
  def lengthOfLongestSubstring(s: String){
    var longestSubstring = 0;
    val substringMap = scala.collection.mutable.HashMap[Char, Int]()
    
    var substringStart = 0;
    for((c, substringFinish) <- s.zipWithIndex){
      substringMap get c match {
        case Some(k) => substringStart = substringStart max (k + 1)
        case None => 
      }
      longestSubstring = longestSubstring max (substringFinish - substringStart + 1)
      substringMap put(c, substringFinish);
    }
    return longestSubstring;
  }

  import scala.collection.immutable.HashMap
  def lengthOfLongestSubstring(s: String){
    @annotation.tailrec
      def recursiveLengthOfLongestSubstring(substringMap: HashMap[Char, Int], 
        substringStart: Int, substringFinish: Int, longestSubstring: Int): Int = {
          
        if(substringFinish == s.length()){ longestSubstring }
        else {
          val c = s.charAt(substringFinish)
          substringMap get c match {
            case Some(k) => 
              val updatedSubstringStart = substringStart max (k + 1)
              recursiveLengthOfLongestSubstring(
              substringMap + (c -> substringFinish),
              updatedSubstringStart,
              substringFinish + 1,
              longestSubstring max (substringFinish + 1 - updatedSubstringStart)
            )
            case None => recursiveLengthOfLongestSubstring(
              substringMap + (c -> substringFinish), 
              substringStart, 
              substringFinish + 1,
              longestSubstring max (substringFinish + 1 - substringStart),
            )
          }
        }
      }
      recursiveLengthOfLongestSubstring(HashMap(), 0, 0, 0)
  }
}