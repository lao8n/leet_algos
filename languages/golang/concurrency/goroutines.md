Goroutines

Vs Threads = Goroutines are not OS threads, or green threads (managed by a language's runtime), they are a higher level of abstraction

Vs Coroutines = coroutines are concurrent subroutines (functions, closures or methods in Go) that are non preemptive - that is they cannot be interrupted. Instead, coroutines have multiple points throughout which allow for suspension or reentry

Goroutines = don't define their own suspension or reentry points - Go's runtime observes the runtime behaviour of goroutines and automatically suspends them when they block and resumes them when they become unblocked.

M:N scheduler = it maps M green threads to N Os threads. Goroutines are then scheduled onto the green threads. 

Fork-join model = at any point in the program it can fork, splitting off a child branch of execution before joining at some point in the future. Forks are implemented with goroutines and joins with sync.WaitGroups. 

Closures = be careful of closures - if the goroutine finishes too fast it might look for some variable on the heap instead. Goroutines can access the same memory so if want to have separate memory make sure to copy or need to synchronise with mutexes or communicate via channels

Done channel - the creator of a goroutine is also responsible for ensuring it can stop the goroutine
```
func newRandStream(done <- chan interface{}) <- chan interface{} {
    c := make(chan interface{})
    go func() {
        defer close(c)
        for {
            select {
                case c <- rand.Int():
                case <- done:
                    return
            }
        }
    }()
    return c
} 
done := make(chan interface{})
randStream := newRandStream(done)
for i := 0; i < 3; i++ {
    fmt.Println(<- randStream)
}
close(done)
```