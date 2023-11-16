Goroutines
```
go echo(c, input.Text(), 1 *time.Second)
go func() { ... }()
```

Wait groups
```
var wg sync.WaitGroup
...
wg.Add(1)
defer wg.Done()
wg.Wait()
```
Select
```
select {
    case <- ch1: 
    case x := <- ch2:
    default:
}
```
Mutex - easier thatn using a channel with single buffered amount of acquiring and releasing token
```
mu sync.Mutex
mu.Lock()
mu.Unlock()
```