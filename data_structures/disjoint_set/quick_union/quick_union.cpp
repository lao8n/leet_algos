// in disjoint set the key data structure is an array where the index is the node and the value is the parent
// key choice: is that value of the parent the root node or just the parent?
// if it is the parent then -> quick union
// we can optimise slightly where given two trees we always have the root of one tree connect to the root of another
// this means if it is a single node being connected they are all connected to the root - but it does mean
// if there is a tree being connected you have to search through