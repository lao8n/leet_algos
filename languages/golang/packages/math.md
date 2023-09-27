Constants
```
math.Pi
```
math.Floor()
math.Ceil
```

```
Random
```
import "math/rand"
rand.Seed(time.Now().UnixNano())
// O(n) complexity based upon Fisher-Yates shuffle algorithm (Knuth)
rand.Shuffle(len(d.cards), func(i, j int) {
    d.card[i], d.card[j] = d.card[j], d.card[i]
})
rand.Int() // up to MaxInt64
rand.Intn(len(slice))
```