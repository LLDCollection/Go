# LRU Cache Problem

## Problem Statement
Design and implement a data structure for a Least Recently Used (LRU) Cache. It should support the following operations: `get` and `put`.

1. `get(key)` - If the key exists in the cache, return its value; otherwise, return `-1`.
2. `put(key, value)` - Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The cache should be initialized with a positive size capacity, and it should work efficiently with respect to time complexity.

### Example
Consider the LRU Cache initialized with a capacity of `2`.

- `put(1, 1)`
- `put(2, 2)`
- `get(1)` returns `1`
- `put(3, 3)` evicts key `2`
- `get(2)` returns `-1` (not found)
- `put(4, 4)` evicts key `1`
- `get(1)` returns `-1` (not found)
- `get(3)` returns `3`
- `get(4)` returns `4`

### Problem Extension
1. All operations (`get` and `put`) must be performed in O(1) time complexity.

## Design Pattern
Use the *Composite Pattern* for this solution:
- The `LRUCache` class acts as a composite that can grow dynamically and handles multiple key-value pairs efficiently.
- The combination of a hash map and a doubly-linked list allows for constant time complexities for both `put` and `get` operations by quickly managing the order of elements based on their usage.
