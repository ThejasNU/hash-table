# hash-table

An [open-addressed](https://en.wikipedia.org/wiki/Open_addressing), [double-hashed](https://en.wikipedia.org/wiki/Double_hashing) hash table in GO. This is also known as map,dictionary or key-value store. Inspired by tutorial [write-a-hash-table](https://github.com/jamesroutley/write-a-hash-table).

## Features of the implementation

- It is similar to unordered_map in C++, which also uses hashing for storing values of each key.
- Hashfunction used here is `polynomial rolling hash function`.
- Collision has been handled using double-hashing method.
- The table dynamically resizes itself when the load is above 70% or below 10%.
- All operations on hash-table have constant time complexity .i.e., O(1).
