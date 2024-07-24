class BoundedBlockingQueue {
    std::queue<int> queue_;
    std::condition_variable cv_;
    std::mutex mtx_;
    int capacity_;
public:
    BoundedBlockingQueue(int capacity) : capacity_(capacity) {}
    
    void enqueue(int element) {
        std::unique_lock lock(mtx_); // cannot use simpler guard_lock automatically acquires lock and unlocks when destroyed
        cv_.wait(lock, [this] { return queue_.size() < capacity_; });
        queue_.push(element);
        cv_.notify_all();
    }
    
    int dequeue() {
        std::unique_lock lock(mtx_);
        cv_.wait(lock, [this] { return !queue_.empty(); });
        int element = queue_.front();
        queue_.pop(); // doesn't return anything hence need to front
        cv_.notify_all();
        return element;
    }
    
    int size() {
        std::unique_lock lock(mtx_);
		return static_cast<int>(queue_.size()); // size return size_t which is unsigned integral type
    }
};