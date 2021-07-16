# Channel Synchronization | [Source](https://gobyexample.com/channel-synchronization)

We can use channels to synchronize execution across goroutines. Here's an example of using a blocking receive ti wait for a goroutine to finish. When waiting for multiple goroutines to finish, better solutions is to use a `WaitGroup`.
