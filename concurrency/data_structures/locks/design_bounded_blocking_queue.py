from collections import deque # already support thread-safe performance
from threading import Lock

# enqueue
class BoundedBlockingQueue(object):
    def __init__(self, capacity: int):
        # blocks by putting thread into wait or sleep state - then wait queue is woken up and tested.
        # need two locks because need to be able to say block enqueues but not dequeues
        self.enqueue_lock, self.dequeue_lock = Lock(), Lock() 
        self.queue = deque() # list is also possible though not thread-safe
        self.capacity = capacity
        self.dequeue_lock.acquire()

    # enqueue_lock: can only acquire if queue not full and not in use
    def enqueue(self, element: int) -> None:
        self.enqueue_lock.acquire()
        self.queue.append(element) # with deque this is atomic
        if len(self.queue) < self.capacity:
            self.enqueue_lock.release()
        if self.dequeue_lock.locked(): # dangerous - could lead to race condition
            self.dequeue_lock.release()

    def dequeue(self) -> int:
        self.dequeue_lock.acquire()
        val = self.queue.popleft() # if list then pop(0)
        if len(self.queue):
            self.dequeue_lock.release()
        if val and self.enqueue_lock.locked(): 
            self.enqueue_lock.release()
        return val

    def size(self) -> int:
        return len(self.queue)
        