```
func add(x int, y int) int { return x + y}
```
Function literals
```
return func(x int) int { return x++ }
```
Deferred function calls
```
defer f.Close() // guaranteed to run after return statement
```
Methods
```
type Point struct { X, Y float64 }
func (p Point) Distance(q Point) float64 {
    return math.Hypo(q.X-p.X, q.Y-p.Y)
}
// pointer receives = either want to update or want to avoid copying because too large
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
``` 
Interfaces
```
type Reader interface {
    Read(p []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}
var w io.Writer
w = os.Stdout // now has interface type 
```
Interface satisfaction
```
// it is okay for *T pointer receiver method to satisfy T interface method but not the other way around
// it will automatically reference
```