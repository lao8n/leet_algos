3 approaches to managing memory
1. Garbage collector
2. Explicit allocation and free-ing
3. Ownership - Rust

Ownership rules
1. Each value in rust has an owner
2. There can only be one owner at a time
3. When the owner goes out of scope the value will be dropped

C++ unique pointers
* Same concept except Rust enforces it at compile time