Escape analysis 
* used to decide whether a variable is allocated on the stack or the heap. 
* allocated on stack if the local variable does not escape the scope of the function
* some data structures are always allocated on the heap like slices and maps that have underlying components that can be updated.