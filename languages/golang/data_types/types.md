Ints
```
var apples int = 1
```

Floats
```
pi := 3.141
f := 1e100
const e = 2.71828
```

Booleans
```
flag := true
```

Strings
```
s := "hello world" // seqs of immutable runes
c := s[0] // ith- byte not necessarily i-th rune if non-ASCII as can be two bytes
r := `raw string` // do not interpret escape sequences etc
len(utf8.RuneCountInString("世界")) // 2 vs len() which is 5 but `range` loop decodes implicitly.
r := []rune(s) // convert string to slice of runes
b := []byte(s) // convert string to a slice of bytes
string(r) // convert from runes or bytes back to string
```

Types
```
type Celsius float64
type Fahreinheit float64
```

Constants
```
const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
)
type Weekday int
const (
    Sunday Weekday = iota // iota starts at 0
    Monday 
    ...
)
```

Arrays
```
var a [3]int
q := [...]int{1, 2, 3} // len 3
```

Slices
```
var s []int
make([]T, len) // slice has length and capacity
x := []int{1, 2, 3}
x = append(x, 1)
x = append(x, x...)
```
Note cannot compare slices too each with equals

Maps
```
ages := make(map[string]int)
ages["bob"] = 7
if age, ok := ages["bob"]; !ok { ... }
```
Note cannot compare maps to each other with equals

Structs
```
type Employee struct {
    ID int
    Name string
    var time.Time
}

var dilbert Employee
p := Point{1, 2}
pp := &Point{1, 2}
```
Struct embedding
```
type ColoredPoint struct {
    Point
    Color color.RGBA
}
```
Type switches
```
switch x.(type) {
    case nil:
    case int, uint:
    ...
    default:
}
```
Generics
```
type Number interface {
    int64 | float64
}
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```