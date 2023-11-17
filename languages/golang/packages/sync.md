```
var wg sync.WaitGroup
```

Specialized map safe for concurrent access
```
sync.Map 
```
Good if entry for a given key is only written once but read many times or if multiple goroutines read, write and overwrite entries for disjoint sets of keys