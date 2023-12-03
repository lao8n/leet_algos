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

```
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
for now := range ticker.C {
    fmt.Println("Tick a", now)
}
```