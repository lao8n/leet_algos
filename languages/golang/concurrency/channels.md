***Channels***
```
ch := make(chan int)
ch <- x // send x
x = <- ch // receive x
close(ch)
```

***Unidirectional channels***
```
func squarer(out chan<- int, in <-chan int) {
    for v := range in {
        out <- v * v
    }
    close(out)
}
```

***Buffered channels***
```
ch = make(chan string, 3)
```
![Primitives vs channels decision tree](channels_vs_primitives.png)

Select
```
select {
    case <- ch1: 
    case x := <- ch2:
    default:
}
```

***Internal representation***
* Channels are used for communication between goroutines and can be used to send and receive values enabling synchronization and data exchange.
* Two main types 1. Unbuffered: block the sending goroutine until the other goroutine receives the value 2. buffered channels have capacity and only block when the buffer is full (on send) or empty (on receive)

Trade-off between buffered and unbuffered channels
* Advantages of unbuffered channels 1. direct synchronization 2. guaranteed communication 3. design clarity
* Advantages of buffered channels 1. increased deadlock potential

`hchan`
* References to goroutines that are waiting to send and receive from the channel
* Channel operations are guarded by a lock to manage synchronization
* There is a buffer to store the buffered channel

* Send = 1. check if waiting receiver if so - transfers the data resuming both. 2. If no receiver and buffered space available then store item and continue 3. o/w block
* Receive = 1. check if buffer empty if not take value and continue 2. if buffer empty and waiting sender then take value 3. if no sender then wait.

Closed channels
* No buffered data is lost it just means no more data will be sent
* You can still receive from a closed channel it just returns immediatley with the false value (or for single value return the zero-value)
* Be careful about closign a channel concurrently and then trying to still send data on it