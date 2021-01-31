# Concurrency

## Concurrency vs Parallelism

Foremost we have to clarify difference between **concurrency** and **parallelism**. 
- **Concurrency is the composition of independently executing processes** (Concurrency is about *dealing with* lots of things at once)
- **Parallelism is the simultaneous execution of computations, which can be related** (Parallelism is about *doing* lots of things at once. (Todd McLeod))

## Share by communicating

Go provides concurrent programming by encouraging a different approach in which shared values are passed around on chanel and, in fact, never actively shared by separated threads of execution. Only one *goroutine* has access to the value at any given time. Data races can't occur, by design. So there is a slogan:
>Do not communicate by sharing memory; instead, share memory by communicating. 

This approach can be taken too far. Reference counts may be best done by putting a mutex around an integer variable, for instance. But as a high-level approach, using channels to control access makes it easier to write clear, correct programs.

More about concurrent programing in go you can find [here](https://golang.org/doc/effective_go.html#concurrency)

## Goroutines

They're called **goroutines** because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations. A goroutine has a simple model: **it is a function executing concurrently with other goroutines in the same address space**. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.


To create a goroutine we use the keyword `go` as up front of calling the function. Then teh function will be executed concurrently and main code flow won't wait for it. 

## WaitGroup

But if we need to wait of goroutine we can use **WaitGroup**, which waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished

## Mutex

A **mutex** is a mutual exclusion lock. **Mutexes allow us to lock our code so that only one goroutine can access that locked chunk of code at a time.**