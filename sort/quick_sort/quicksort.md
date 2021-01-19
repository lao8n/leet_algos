quicksort is a conquer-than-divide algorithm, which does most of the work during the partitioning
and recursive calls. the merging of the sorted partitions is trivial
this is in contrast to merge sort which is a divide-than-conquer algorithm, with the partitioning
being trivial - just splitting the array in half but most of the work happens in the recursive 
calls and merge phase

quicksort algorithm can be summarised as:
1. choose the pivot
2. partition elements as <, = or > the pivot
3. recursively sort the partitions
4. join the sorted partitions together

complexity
for good pivots we have O(n logn) but potentialy O(n^2) for the worst pivots equivalent to bubble
sort.

improvements
1. the key improvement from the naive algorithm is to do things in place, saving on the memory allocation
cost of more arrays
2. choose the pivot more effectively e.g. the median of 3 randomly chosen elements
3. for sufficiently small arrays switch to insertion sort avoiding too much recursion at the leaf of tree

differences across languages
essentially the languages are the same except for the languages with immutable data structures like erlang 
which cannot do the in-place approach