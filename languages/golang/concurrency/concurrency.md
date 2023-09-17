Goroutines
```
go echo(c, input.Text(), 1 *time.Second)
go func() { ... }()
```

Channels
```
ch := make(chan int)
ch <- x // send x
x = <- ch // receive x
close(ch)
```
Unidirectional channels
```
func squarer(out chan<- int, in <-chan int) {
    for v := range in {
        out <- v * v
    }
    close(out)
}
```
Buffered channels
```
ch = make(chan string, 3)
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
Mutex - easiet thatn using a channel with single buffered amount of acquiring and releasing token
```
mu sync.Mutex
mu.Lock()
mu.Unlock()
```