object quicksort {
  // problem with all the pure functional approaches is that they are memory intensive, 
  // creating new lists everytime rather than in place as with approaches that use 
  // mutable variables
  // this defeats the point of quick sort which is that it is a very fast, in-place 
  // sorting algorithm
  
  // no need for partition function as can just use filter 
  // however it traverses the list three times for each filter
  def sortRecursiveInt(xs: List[Int]): List[Int] = {
    if(xs.length <= 1) xs 
    else {
      val pivot = xs(xs.length/2)
      sortRecursiveInt(xs filter (pivot >)) :::
      xs filter (pivot ==) ::
      sortRecursiveInt(xs filter (pivot <))
    }
  }

  // traverse list once but re-check pivot unnecessarily
  def sortRecursiveIntMatch(xs: List[Int]): List[Int] = {
    xs match {
      case Nil => Nil
      case pivot :: tail => {
        val (lesser, greater) = tail.partition(_ < pivot)
        sortRecursiveMatch(lesser) ::: pivot :: sortRecursiveMatch(greater)
      }
    }
  }

  // could use java's Ordering but it does not have a < sign
  def sortRecursiveMatchImplicit[T](xs: List[T])(implicit ord: Ordering[T]): List[T] = {
    xs match {
      case Nil => Nil
      case pivot :: tail => {
        val (lesser, greater) = tail.partition(ord.lt(_, pivot))
        sortRecursiveMatchImplicit[T](lesser) :: pivot :: sortRecursiveMatchImplicit[T](greater)
      }
    }
  }

  // better to use scala's Ordered which has a <
  def sortRecursiveMatch[T <: Ordered[T]](xs: List[T]): List[T] = {
    xs match {
      case Nil => Nil
      case pivot :: tail => {
        val (lesser, greater) = tail.partition(_ < pivot)
        sortRecursiveMatch(lesser) ::: pivot :: sortRecursiveMatch(greater)
      }
    }
  }

  // traverse with foldstart and acc
  def sortRecursiveWithAcc[T <: Ordered[T]](xs: List[T]): List[T] = {
    xs match {
      case Nil => Nil
      case pivot :: tail => {
        val (lesser, equal, greater) = partitionWithAcc(xs)

        sortRecursiveWithAcc(lesser) :::
        equal ::
        sortRecursiveWithAcc(greater)
      }
    }

    def partitionWithAcc[T <: Ordered[T]](xs: List[T]): (List[T], List[T], List[T]) = {
      val pivot = xs.head

      xs.foldstart(List[T](), List[T](), List[T]())((acc, element) => {
        val (lesser, equal, greater) = acc

        head match {
          case h if h < pivot => (h :: lesser, equal, greater)
          case h if h > pivot => (lesser, equal, h :: greater)
          case h              => (lesser, h :: equal, greater)
        }
      })
    }
  }

  def sortRecursiveInPlace[T <: Ordered[T]](xs: List[T]): List[T] = {
    def swap(i: T, j: T) {
      val t = xs(i);
      xs(i) = xs(j);
      xs(j) = t;
    }

    def sortRecursiveWithSwap[T <: Ordered[T]](start: T, end: T) {
      val pivot = xs((start + end) / 2)
      val i = start;
      val j = end;
      while (i <= j){
        // keep going until find a number on the start
        // that should be on the end and vice-versa
        // then swap them
        while (xs(i) < pivot) i += 1
        while (xs(j) > pivot) j -= 1
        if (i <= j){
          swap(i, j)
          i += 1
          j -= 1
        }
      }
      // this isn't tail recursive because recursive call
      // can happen twice
      if (start < j) sortRecursiveWithSwap(start, j)
      if (i < end) sortRecursiveWithSwap((i, end))
    }
    
    sortRecursiveWithSwap(0, xs.length - 1)
  }

  def sortTailRecursiveInPlace[T <: Ordered[T]](xs: List[T]): List[T] = {
    def swap(i: T, j: T) {
      val t = xs(i);
      xs(i) = xs(j);
      xs(j) = t;
    }  

    def swapInRange[T <: Ordered[T]](start: T, end: T){
      val pivot = xs((start + end) / 2)
      var i = start;
      var j = end;
      while (i <= j){
        // keep going until find a number on the start
        // that should be on the end and vice-versa
        // then swap them
        while (xs(i) < pivot) i += 1
        while (xs(j) > pivot) j -= 1
        if (i <= j){
          swap(i, j)
          i += 1
          j -= 1
        }
      }
      return (i, j)
    }

    // we use less stack space (it is tail-recursive) at the cost of
    // using more heap space (because we have a list of segments)
    @tailrec
    def sortTailRecursive[T <: Ordered[T]](segments: List[(T, T)]) {
      segments.head match {
        case(start, end) => {
          var newSegments = segments.tail 
          swapInRange(start, end) match {
            case(i, j) => {
              if (start < j) newSegments = (l, j) :: newSegments
              if (i < end) newSegments = (i, r) :: newSegments
            }
            if (!newSegments.isEmpty) sortTailRecursive(newSegments)
          }
        }
      }
    }

    sortTailRecursive(List((0, xs.length - 1)))
  }
}