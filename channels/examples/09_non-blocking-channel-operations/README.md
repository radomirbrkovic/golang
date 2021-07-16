# Non-Blocking Channel Operations

Basic sends and receives on channels are blocking. However, we can use `select` with `default` clause to implement *non-blocking* sends, receives and even non-blocking multi-way `selects`.  