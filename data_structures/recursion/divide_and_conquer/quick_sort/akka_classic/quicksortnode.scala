

import akka.actor.{Props, ActorRef, Actor}

// not sure how to make this generic

class QuickSortNode extends Actor {

  override def receive: Actor.Receive = {
    case Sort(xs) => {
      xs match {
        case x :: Nil => sender ! Result(List(x))
        case Nil => sender ! Result(List())
        case pivot :: tail => {
          val (lesser, greater) = tail.partition(_ < pivot)
          val leftSorter = context.actorOf(Props[QuickSortNode])
          val rightSorter = context.actorOf(Props[QuickSortNode])
          leftSorter ! Sort(lesser)
          rightSorter ! Sort(greater)
          context.become(waitingBothResults(pivot, leftSorter, rightSorter))
        }
      }
    }
  }

  def waitingBothResults(pivot: Int, leftSorter: ActorRef, rightSorter: ActorRef): Receive = {
    case Result(firstRes) => {
      if (sender() == leftSorter) {
        context.become(waitingRightResult(pivot, firstRes))
      } else {
        context.become(waitingLeftResult(pivot, firstRes))
      }
    }
  }

  def waitingRightResult(pivot: Int, leftRes: List[Int]): Receive = {
    case Result(rightRes) => {
      context.parent ! Result(leftRes ++ List(pivot) ++ rightRes)
      context.stop(self)
    }
  }

  def waitingLeftResult(pivot: Int, rightRes: List[Int]): Receive = {
    case Result(leftRes) => {
      context.parent ! Result(leftRes ++ List(pivot) ++ rightRes)
      context.stop(self)
    }
  }

}