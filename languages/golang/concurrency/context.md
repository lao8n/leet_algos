***Context***

Alternative to done is sharing context. Standard pattern to communicate extra information alongside the information to cancel. 

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <- chan struct{}
    Err() error
    Value(key interface{}) interface{} // request specific information to be passed along as well as preemption
}
```

Purpose
1. an API for canceling branches of your call-graph
2. Data-bag for transporting request-scoped data through your call graph

Context is immutable but functions like `WithCancel`, `WithDeadline` and `WithTimeout` allow changes to them.
Create context with `Background()`. 

```
func main() {
    var wg sync.WaitGroup
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    wg.Add(1)
    go func() {
        defer wg.Done() 
        if err := printGreeting(ctx); err != nil {
            fmt.Printf("cannot print greeting: %v\n", err)
            cancel()
        }
    }()
    ...
    func printGreeting(ctx context.Context) error {
        greeing, err := genGreeting(ctx)
        if err != nil {
            return err
        }
        fmt.Printf("%s world!\n", greeting)
        return nil
    }
    ...
    func genGreeting(ctx context.Context) (string, error) {
        ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
        defer cancel()

        switch locale, err := locale(ctx); {
            case err != nil:
                return "", err
            case locale == "EN/US":
                return "hello", nil
            }
        }
        return "", fmt.Errorf("unsupported locale")
    }
    ...
    func locale(ctx context.Context) (string, error) {
        if deadline, ok := ctx.Deadline(); ok {
            if deadline.Sub(time.Now().Add(1 * time.Minute)) <= 0 {
                return "", context.DeadlineExceeded
            }
        }
        select {
            case <- ctx.Done():
            return "", ctx.Err()
            case <- time.After(1 * time.Minute):
        }
        return "EN/US", nil
    }
    wg.Wait()
}
```