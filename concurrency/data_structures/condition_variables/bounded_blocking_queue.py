from collections import deque
from threading import Condition

class BoundedBlockingQueue(object):

    def __init__(self, capacity: int):
        self.c = Condition()
        self.q = deque()
        self.capacity = capacity

    def enqueue(self, element: int) -> None:
        with self.c: # context managers handle setup and tear down
            self.c.wait_for(lambda: len(self.q) < self.capacity) # better than polling, but still involves thread being woken up
            self.q.append(element)
            self.c.notify_all() # only two threads so notify is fine

    def dequeue(self) -> int:
        with self.c:
            self.c.wait_for(lambda: len(self.q) > 0)
            val = self.q.popleft()
            self.c.notify_all()
            return val
        
    def size(self) -> int:
        return len(self.q)