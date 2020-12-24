object quicksort {
  // problem with all the pure functional approaches is that they are memory intensive, 
  // creating new lists everytime rather than in place as with approaches that use 
  // mutable variables
  // this defeats the point of quick sort which is that it is a very fast, in-place 
  // sorting algorithm
  
  // no need for partition function as can just use filter 
  // however it traverses the list three times for each filter
  def sortRecursive(xs: List[A]): List[A] = {
    if(xs.length <= 1) xs 
    else {
      val pivot = xs(xs.length/2)
      sortRecursive(xs filter (pivot >)) :::
      xs filter (pivot ==) ::
      sortRecursive(xs filter (pivot <))
    }
  }

  // traverse list once but re-check pivot unnecessarily
  def sortRecursiveMatch(xs: List[A]): List[A] = {
    xs match {
      case Nil => Nil
      case pivot :: tail => {
        val (lesser, greater) = tail.partition(_ < pivot)
        sortRecursiveMatch(lesser) ::: pivot :: sortRecursiveMatch((greater))
      }
    }
  }

  // traverse with foldLeft and acc
  def sortRecursiveWithAcc(xs: List[A]): List[A] = {
    xs match {
      case Nil => Nil
      case pivot :: tail => {
        val (lesser, equal, greater) = partitionWithAcc(xs)

        sortRecursiveWithAcc(lesser) :::
        equal ::
        sortRecursiveWithAcc(greater)
      }
    }

    def partitionWithAcc(xs: List[A]): (List[A], List[A], List[A]) = {
      val pivot = xs.head

      xs.foldLeft(List[A](), List[A](), List[A]())((acc, element) => {
        val (lesser, equal, greater) = acc

        head match {
          case h if h < pivot => (h :: lesser, equal, greater)
          case h if h > pivot => (lesser, equal, h :: greater)
          case h              => (lesser, h :: equal, greater)
        }
      })
    }
  }

  def sortRecursiveInPlace(xs: List[A]): List[A] = {
    def swap(i: A, j: A) {
      val t = xs(i);
      xs(i) = xs(j);
      xs(j) = t;
    }

    def sortRecursiveWithSwap(left: A, right: A) {
      val pivot = xs((left + right) / 2)
      val i = left;
      val j = right;
      while (i <= j){
        // keep going until find a number on the left
        // that should be on the right and vice-versa
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
      if (left < j) sortRecursiveWithSwap(l, j)
      if (i < right) sortRecursiveWithSwap((i, r))
    }
    
    sortRecursiveWithSwap(0, xs.length - 1)
  }

  def sortTailRecursiveInPlace(xs: List[A]): List[A] = {
    def swap(i: A, j: A) {
      val t = xs(i);
      xs(i) = xs(j);
      xs(j) = t;
    }  

    def swapInRange(left: A, right: A){
      val pivot = xs((left + right) / 2)
      var i = left;
      var j = right;
      while (i <= j){
        // keep going until find a number on the left
        // that should be on the right and vice-versa
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
    def sortTailRecursive(segments: List[(A, A)]) {
      segments.head match {
        case(left, right) => {
          var newSegments = segments.tail 
          swapInRange(left, right) match {
            case(i, j) => {
              if (left < j) newSegments = (l, j) :: newSegments
              if (i < right) newSegments = (i, r) :: newSegments
            }
            if (!newSegments.isEmpty) sortTailRecursive(newSegments)
          }
        }
      }
    }

    sortTailRecursive(List((0, xs.length - 1)))
  }
}