public class BubbleSort<T extends Comparable<? super T>> {
  public static void main(String[] args){
    String[] arrayOfString = {"example", "array"};
    BubbleSort<String> stringSorter = new BubbleSort<>();
    stringSorter.sortBubble(arrayOfStrings, 0 arrayOfStrings.length - 1);
  }

  private void sortBubble(T[] xs, int start, int end){
    for(int loop = start; loop < end; loop++){ // number of loops
      for(int i = start; i < end - loop; i++){ 
        if(xs[i].compareTo(xs[i+1]) > 0 ){
          swap(xs, i, i+1)
        }
      }
    }
  }

  private void swap(T[] xs, int i, int j){
    T t = xs[i];
    xs[i] = xs[j];
    xs[j] = t;
  }
}