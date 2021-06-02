class InsertionSort {
  void sort(int xs[]{
    int n = xs.length;
    for(int itemToSort = 1; itemToSort < n; ++itemToSort){
      int key = xs[i];
      int itemSorted = itemToSort - 1;
      // go through already sorted items shifting them along to make space for 
      // itemToSort
      while(itemSorted >= 0 && xs[itemSorted] > key){
        xs[itemSorted + 1] = xs[itemSorted];
        itemSorted = itemSorted - 1;
      }
      xs[itemSorted + 1] = key;
    }
  })
}