# Select instruction 

The  `select` is really powerful command which can help us in many situations. The `select` instruction in go language is looks like `switch` instruction, but it works with channels, that mean the `select` instruction provides goroutines can communicate with many channels. The main benefit of using select instruction is non-blocking operations between channels. 

We should be careful with using this instruction because it can block all channels.  

## Example 

```
select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
```