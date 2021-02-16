# Hash Tables

One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.

## Implementation

Internally the hash table will be represented with a `map` type, and I’ll expose the following methods: 
- `Put()`
- `Remove()`
- `Get()`
- `Size()`

I’ll create a `ValueHashTable` generic type, concurrency safe, that can generate hash tables containing any type. Keys can be any type as long as it implements the `Stringer` interface. By running `go generate` the code will create a type-specific Hash Table implementation, encapsulating the actual value-specific data structure containing the data.