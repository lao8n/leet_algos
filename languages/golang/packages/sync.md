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

Specialized map safe for concurrent access
```
sync.Map 
```
Good if entry for a given key is only written once but read many times or if multiple goroutines read, write and overwrite entries for disjoint sets of keys