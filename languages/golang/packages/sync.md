***Wait Group***
```
var wg sync.WaitGroup
wg.Add(1)
go func(){
    defer wg.Done()
    // do something
}()
wg.Wait()
```
Implemented as a counter where the method blocks until the wait group counter is 0.
***Mutex***
```
var lock sync.Mutex
increment := func(){
    lock.Lock()
    defer lock.Unlock()
    count++
}
```
Allows multiple readers or a single writer.
```
var m sync.RWMutex
```
***Condition***
```
for conditionTrue() == false { } // problem uses all cycles of one core
for conditionTrue() == false { time.Sleep() } // problem inefficient and potential wait
c := sync.NewCond(&sync.Mutex{})
c.L.Lock()
for conditionTrue() == false {
    c.Wait() // doesn't just block the goroutine but suspends it allowing other threads to run
} 
c.L.Unlock()

// also have 
c.Signal() // one of two methods to signify to blocked goroutine waiting that condition has been triggered
c.Broadcast() // sends it to all goroutines waiting
```
Click a button and broadcast that it has been clicked to all subscribers
```
type Button struct {
    Clicked *sync.Cond
}
button := Button{ Clicked: sync.NewCond(&sync.Mutex{})}
subscribe := func(c *sync.Cond, fn func()){
    var goroutineRunning sync.WaitGroup
    goroutineRunning.Add(1)
    go func() {
        goroutineRunning.Done()
        c.L.Lock()
        defer c.L.Unlock()
        c.Wait()
        fn()
    }()
    goroutineRunning.Wait()
}

var clickRegistered sync.WaitGroup
clickRegistered.Add(3)
subscribe(button.Clicked, func() {
    fmt.Println("do a")
    clickRegistered.Done()
})
...
button.Clicked.Broadcast()
clickRegisted.Wait()
```
***Once*** 

ensures something is only called once even on different goroutines. 
```
var once sync.Once
var increments sync.WaitGroup
increments.Add(100)
for i := 0; i < 100; i++ {
    go func(){
        defer increments.Done()
        once.Do(increment)
    }()
}
increments.Wait()
// Count is 1
```

***Pool***

```
myPool := &sync.Pool {
    New : func() interface{} {
        return struct{}{}
    }
}
instance := myPool.Get()
myPool.Put(instance) // returns the instance to be available to get
```
Pool will garbage collect put memory

***Data structures***

Concurrent safe implementation of the pool pattern where you can constrain the creation of things that are expensive (such as database connections)

Specialized map safe for concurrent access
```
sync.Map 
```
Good if entry for a given key is only written once but read many times or if multiple goroutines read, write and overwrite entries for disjoint sets of keys