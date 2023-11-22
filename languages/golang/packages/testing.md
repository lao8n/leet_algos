```
go test -bench=.
```

```
func Benchmark(b *testing.B){
    b.StartTimer() 
    b.ResetTimer()
}
```