public class SortRecursiveInPlace<T extends Comparable<? super T>> {
  public static void main(String[] args){
    String[] arrayOfString = {"example", "array"};
    SortRecursiveInPlace<String> stringSorter = new SortRecursiveInPlace<>();
    stringSorter.sortRecursiveWithSwap(arrayOfStrings, 0, arrayOfStrings.length - 1);
  }

  // basically standard in place approach
  // generic comes from T extending Comparable with associated compareTo function
  public void sortRecursiveWithSwap(T[] xs, int start, int end){
    if(start < end){
      int pivot = partition(xs, start, end);
      sortRecursiveWithSwap(xs, start, pivot);
      sortRecursiveWithSwap(xs, pivot + 1, end);
    }
  }

  // uniquely we swap in the partition function 
  private int partition(T[] xs, int start, int end){
    T pivot = xs[(start + end) / 2];
    int i = start--;
    int j = end++;

    while(true){
      do i++;
      while (xs[i].compareTo(pivot) < 0);

      do j--;
      while (xs[j].compareTo(pivot) > 0);

      if(i >= j) return j;

      swap(i, j)
    }
  }

  private void swap(T[] xs, int i, int j){
    T t = xs[i];
    xs[i] = xs[j];
    xs[j] = t;
  }
}