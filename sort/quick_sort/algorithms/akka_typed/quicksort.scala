package akka_typed

import akka.actor.typed._
import akka.actor.typed.scaladsl._

class QuickSort extends App {
    sealed trait SortMessage
    final case class Sort(xs: List[Int]) extends SortMessage
    final case object Result(sorted: List[Int]) extends SortMessage

    val sort: Behavior[SortMessage] =
        Behaviors.setup { ctx =>   
            Behaviors.receiveMessage[SortMessage] {
                case Sort(xs) => 
                    (pivot, leftSorter, rightSorter) = sort(xs)    
                    waitingBothResults(p)
                case Result =>
                    println("shutting down ...")
                    Behaviors.stopped
            }
        }
        
    val waitingBothResults(parent: ActorRef[SortMessage], pivot: Int, leftSorter: ActorRef[SortMessage], 
        rightSorter: ActorRef[SortMessage]): Behavior[SortMessage] = 
        Behaviors.setup { ctx =>   
            Behaviors.receiveMessage[SortMessage] {
                case Result(firstRes) => {
                    if (sender() == leftSorter) {
                        waitingRightResult(pivot, firstRes)
                    } else {
                        waitingLeftResult(pivot, firstRes)
                    }
                }
            }
        }

    val waitingLeftResult(pivot: Int, rightRes: List[Int]): Behavior[SortMessage] = 
        Behaviors.setup { ctx =>   
            Behaviors.receiveMessage[SortMessage] {
                case Result(leftRes) => {
                context.parent ! Result(leftRes ++ List(pivot) ++ rightRes)
                context.stop(self)
                }
            }
        }

    val waitingRightResult(pivot: Int, leftRes: List[Int]): Behavior[SortMessage] = 
        Behaviors.setup { ctx =>   
            Behaviors.receiveMessage[SortMessage] {
                
            }
        }

    ActorSystem(Behaviors.setup[Any] { ctx =>
        val greeterRef = ctx.spawn(greeter, "greeter")
        ctx.watch(greeterRef) // sign death pact
    
        greeterRef ! Greet("world")
        greeterRef ! Stop

        Behaviors.empty
    }, "helloworld")
}