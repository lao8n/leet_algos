object Solution {
    def isValid(s: String): Boolean = {
        // brute force
        // cannot count brackets because of example 4
        
        // use memory stack
        import scala.collection.mutable
        s.foldLeft(mutable.Stack[Character]()){
            case (stack, lb) if (lb == '(' | lb == '[' | lb == '{') => stack.push(lb)
            case (mutable.Stack(), _) => return false
            case (stack, ')') if(stack.pop() == '(') => stack
            case (stack, ']') if(stack.pop() == '[') => stack
            case (stack, '}') if(stack.pop() == '{') => stack
            case (stack, _) => return false
        } match {
            case mutable.Stack() => true
            case _ => false
        }
    }
}