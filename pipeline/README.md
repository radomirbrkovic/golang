# Pipeline - Goroutine and Channel

**Pipeline** is process of connected instructions or programs which one's output becomes input of next program. In concrete case pipeline is virtual process for connection `goroutine` and `channel`, where output one `goroutine` becomes input to next `goroutine` with using `channel` for broadcasting data.  

One of advantage of using pipeline is that there is consistent data stream through program, because any goroutine and channel don't need to wait everything be finished before they start. Also using pipeline use less variables and less memory, and makes code architecture simpler and easier for maintaining.    