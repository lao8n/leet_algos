```
start := time.Now()
select {
case <- c:
case <- time.After(1 * time.Second):
    fmt.Println("timed out")
default:
    time.Since(start)
}
```