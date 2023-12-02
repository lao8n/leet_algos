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
for _, s := range []string{"a", "b", "c"} {
    select {
        case <- done: // way to interrupt the for loop on the done channel - best practice to have a done channel
            // so parent can cancel child
            return
        case stringStream <- s:
    }
}
```
Mutex - easier thatn using a channel with single buffered amount of acquiring and releasing token
```
mu sync.Mutex
mu.Lock()
mu.Unlock()
```

Queues
* Doesn't make the system faster just decouples
* Most useful 1. allows for batching to make system faster 2. entrance of pipeline to prevent death spiral
* Queuing theory = Little's Law L = lambda x W where L average number of units in the system, lambda is the average arrival rate of units and W is the average time a unit spends in the system 