# Golang's garbage collector

**Garbage collector** is the process for freeing up memory space that is not being used. In other words, the garbage collector sees which objects are out of scope and cannot be referenced anymore and frees the memory space they consume. The process happens in a concurrent way while a Go program is running and not before or after execution of the program. Go implements [tri-color mark-and-sweep algorithm](https://en.wikipedia.org/wiki/Tracing_garbage_collection#Tri-color_marking) 

