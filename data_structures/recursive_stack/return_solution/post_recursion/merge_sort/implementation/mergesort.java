public class SortRecursiveInPlace<T extends Comparable<? super T>>{
  public static void main(String[] args){
    String[] arrayOfString = {"example", "array"};
    SortRecursiveInPlace<String> stringSorter = new SortRecursiveInPlace<>();
    stringSorter.sortRecursiveWithSwap(arrayOfStrings, 0, arrayOfStrings.length - 1);
  }

  public void sortWithTempArrays(T[] xs){
    if(xs.length < 2){
      return;
    }
    int mid = xs.length / 2;
    int[] leftTemp = new int[mid];
    int[] rightTemp = new int[xs.length - mid];

    // copy into leftTemp
    for(int i = 0; i < mid; i++){
      leftTemp[i] = xs[i];
    }
    for(int i = mid; i < xs.length; i++){
      rightTemp[i - mid] = xs[i];
    }

    // recursively divide in half
    mergeSort(leftTemp);
    mergeSort(rightTemp);
    merge(xs, leftTemp, rightTemp);
  }

  private void merge(int[] xs, int[] leftTemp, int[] rightTemp){
    int xsIndex, leftIndex, rightIndex = 0;
    while(leftIndex < leftTemp.length && rightIndex < rightTemp.length){
      if(leftTemp[leftIndex] <= rightTemp[rightIndex]){
        xs[xsIndex++] = leftTemp[leftIndex++];
      }
      else {
        xs[xsIndex++] = rightTemp[rightIndex++];
      }
    }
    while(leftIndex < leftTemp.length){
      xs[xsIndex++] = leftTemp[leftIndex++];
    }
    while(rightIndex < rightTemp.length){
      xs[xsIndex++] = rightTemp[rightIndex++];
    }
  }

  // TODO: in-place
  public void sortInPlace(T[] xs, int start, int end){
    if(xs.length < 2){
      return;
    }
    int midpoint = (end - start) / 2;
    sortRecursiveWithSwap(xs, start, midpoint);
    sortRecursiveWithSwap(xs, midpoint, end);
    merge(xs, start, end);
  }

  private void merge(T[] xs, int leftStart, int leftEnd, int rightStart, rightEnd){
   
  }

  private void swap(T[] xs, int i, int j){
    T t = xs[i];
    xs[i] = xs[j];
    xs[j] = t;
  }
}