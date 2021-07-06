### scala.collection
* either mutable or immutable
* operations that affect the whole collection
* Collections
  * Seq, Set & Map
* Also
  * Array - O(n) for tail operations as copies whole array

### scala.collection.immutable
* add or remove single values
* Collections 
  * Set: HashSet
  * Map: HashMap
  * IndexedSeq: Vector, String
    * Efficient apply, length and update operations
    * Vector is a compromise between IndexedSeq and LinearSeq with constant time
      access for both implemented with a tree with high branching factor (32 elements)
  * LinearSeq: List
    * Efficient head & tail operations 

### scala.collection.mutable
* side-effecting modifications
* import scala.collection.mutable then call mutable.HashSet etc.
* Collections
  * Set: HashSet
  * Map: HashMap
  * IndexedSeq: ArrayBuffer, ArrayDeque, StringBuilder
    * Efficient apply, length and update operations
  * Buffer: ArrayBuffer, ListBuffer, StringBuffer
    * Allow appending - as opposed to copying of whole new Seq
    * good for building lists and arrays then convert once made

### Mutable and Immutable collections
* Can often replace a mutable collection stored in val with an immutable 
  collection stored in var and vice versa

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
x.concat() x ++ y // expensive because left associative unlike ::: which is right associative
// lists are immutable linked-lists where need to connect lists copying all of first lists 
// https://stackoverflow.com/questions/6559996/scala-list-concatenation-vs
// can always use ++ (unless want type safety) as ++ goes to ::: if list
x.map() x.flatMap() x.collect() 
x.toList(), x.toMap(), x.toSeq() x.to(mutable.Set)
x.isEmpty, x.size // no () is no side effects
x.head, x.last, x.headOption, x.lastOption x.tail, x.filter, x.drop
// filter vs withFilter -> withFilter better because lazy evaluation and single pass
x.partition, x.split, x.groupBy
x.exists, x.forall, x.count
x.foldLeft, x.foldRight, x.reduceLeft, x.reduceRight, x.sum, x.min, x.max
x.mkString
x zip y, x.zipWithIndex
```

### Seq methods
```scala
xs(i), xs.isDefinedAt(i), xs.length, xs.indices
xs.indexOf(x), xs.lastIndexOf(x), xs.indexOfSlice(ys)
x +: xs, ys ++: xs // prepended
xs :+ x, xs :++ ys // appended
xs.padTo(len, x)
xs.sorted, xs.sortWith(lt), xs.sortBy(f)
xs.updated(i, x), xs(i) = x // latter only for mutable.Seq
xs.reverse
xs sameElements ys, xs contains x, 
```

#### Immutable Lists
* Constant time operation for pre-pending, access to first element and to tail

#### Immutable LazyLists
* Can be infinitely long, 
```scala
1 #:: LazyList.empty // instead of stand cons :: 
// example of fibonacci seq
def fibFrom(a: Int, b: Int): LazyList[Int] = a #:: fibFrom(b, a + b)
// with :: operator this would cause an infinite recursion but with #:: it is fine
// then call 
val fibs = fibFrom(1, 1).take(7)
```

#### Immutable Queue
* FIFO
```scala
Queue().enqueue(1), :+ // returns Queue(1)
Queue().enqueueAll(List(2, 3))
q.dequeue // removes first element
```

#### Mutable Array
```scala
val a1 = Array(1, 2, 3)
val a1 = Array.fill[Byte](5)(0)
val a1 = Array.ofDim[String](5)
val a1 = Array.ofDim[Int](5, 4) // 2d array
val a1 = Array.fill[Int](5, 4)(defaultVal) // 2d array
```

#### Mutable.IndexedSeq
```scala
xs.mapInPlace(f), xs.sortInPlace()
```

#### Mutable.Buffer
```scala
buf += x, buf ++= xs // append 
val buf = mutable.ArrayBuffer.empty[Int] // ArrayDequeue has efficient prepend and append
buf += 10
buf.toArray
val buf = mutable.ListBuffer.empty[Int]
buf += 1
buf.toList
val buf = new StringBuilder
buf += 'a'
buf ++= "bcdef"
buf.toString
```

#### Mutable Queue
```scala
val queue = mutable.Queue[String]()
val queue = new mutable.Queue[String]
val queue = mutable.Queue.empty[String]
queue += "a", queue enqueue "a" // instead of enqueue
queue dequeue
```

#### Mutable Stack
```scala
val stack = new mutable.Stack[String]()
val stack = mutable.Stack[String]
val stack = mutable.Stack.empty[String]
stack.push(1)
stack.top
stack.pop
```

### Set methods
```scala
xs contains x, xs(x), xs subsetOf ys
xs concat ys, xs ++ ys
xs.empty
xs intersect ys, xs union ys, xs diff ys
```

#### Immutable Set
```scala
xs + x, xs - x
xs += x, xs -= x // if xs is a var
```

#### Mutable Set
```scala
xs += x, xs ++= ys, xs -= x, xs --= ys
xs(x) = b
```
### Map methods
```scala
ms get k // returns Option
ms(k) // exception if not found
ms getOrElse(k, d) 
ms contains k
ms.keys, ms.keySet, ms.keysIterator, ms.values, ms.valuesIterator
// can also get Iterator with .iterator and iterator.hasNext and iterator.next
```

### Immutable Map Methods
```scala
ms + (k -> v), ms.updated(k, v) // equivalent
ms remove k, ms - k // equivalent
```

### Mutable map methods
```scala
ms(k) = v
ms += (k -> v)
m.put(key, value) // returns an option of value previously associated with key
ms.getOrElseUpdate(k, d) // d is default if no value found and d is returned
```

#### Mutable HashMap
```scala
val map = mutable.HashMap.empty[Int, String]
map += (1 -> "make a website")
map(1)
map contains 2
// Weak HashMaps are good for caches because garbage collector does not follow
// links from the map to the keys stored in it - so if no reference to key then
// key and value entry in hashmap is lost
```

### IO
```scala
val file = Source.fromFile("src/test/.../x.txt")
for (line <- file.getLines()){
  println(line)
}
file.close()
// or 
val reader = file.getLines() // returns Iterator[String]
reader.hasNext
reader.next()
reader.toList
```