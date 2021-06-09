### scala.collection
* either mutable or immutable
* operations that affect the whole collection
* Collections
  * Seq, Set & Map

### scala.collection.immutable
* add or remove single values
* Collections 
  * Set: HashSet
  * Map: HashMap
  * IndexedSeq: Vector, String
  * LinearSeq: List

### scala.collection.mutable
* side-effecting modifications
* import scala.collection.mutable then call mutable.HashSet etc.
* Collections
  * Set: HashSet
  * Map: HashMap
  * IndexedSeq: ArrayBuffer, ArrayDeque, StringBuilder
  * Buffer: ArrayBuffer

### Initialization
```scala
Iterable("x", "y", "z")
Map("x" -> 24, "y" -> 25, "z" -> 26)
Set(Color.red, Color.green, Color.blue)
SortedSet("hello", "world")
Buffer(x, y, z)
IndexedSeq(1.0, 2.0)
LinearSeq(a, b, c)
```

### Iterable methods
```scala
val x = Iterable(), val y = Iterable()
x.concat() x ++ y
x.map() x.flatMap() x.collect() 
x.toList(), x.toMap(), x.toSeq() x.to(mutable.Set)
x.isEmpty, x.size // no () is no side effects
x.head, x.last, x.headOption, x.lastOption x.tail, x.filter
x.partition, x.split, x.groupBy
x.exists, x.forall, x.count
x.foldLeft, x.foldRight, x.reduceLeft, x.reduceRight, x.sum, x.min, x.max
```