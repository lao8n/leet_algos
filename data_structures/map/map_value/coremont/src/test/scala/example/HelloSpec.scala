package example

import org.scalatest._
import scala.io.Source

class SolutionSpec extends FlatSpec with Matchers {

  val file = Source.fromFile("src/test/scala/example/words.txt")
  // val solution = Solution.anagram(file.getLines(), Map.empty)
  val solution2 = Solution.anagram2(file.getLines())
  "The Solution object" should "say hello" in {
    solution2.size should equal (14) //should equal "Hello"
    // for (line <- file.getLines()){
    //   x = line
    // }
  }
  // file.close()
}
