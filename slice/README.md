# Slice

A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is nil.

[Official documentation]( https://golang.org/ref/spec#Slice_types)

**A SLICE HOLDS VALUES OF THE SAME TYPE.**

## Declaration 
 A composite literal is created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.

 `VARIABLE := []TYPE{VALUES}`


## Append

To append values to a slice, we use the special built in function append. This function returns a slice of the same type. 

### Example:

    ```
    x := []int{4, 5, 7, 8, 42}    
    y := []int{100, 101, 108, 128}
    x = append(x, 43, 44, 88, 21, 46)
    x = append(x, y...)
    fmt.Println(x)
    // OUTPUT: [4 5 7 8 42 43 44 88 21 46 100 101 108 128]
    ```

## Slicing a slice

We can slice a slice which means that we can cut parts of the slice away. We do this with the colon operator.

### Example:
```
 x := []int{4, 5, 7, 8, 42}
 fmt.Println(x[1:])
 // OUTPUT[5 7 8 42]
 fmt.Println(x[1:3])
 // OUTPUT[5 7]
```

## Deleting from a slice

We can delete from a slice using both append and slicing. This is a wonderful and elegant example of why Go is great and how Go provides ease of programming.

### Example:

```
x:= []int[1, 2, 3, 4, 5]
// Remove 2nd and 3th element from slice
x = append(x[:2], x[4:]...)
// OUTPUT [1 4 5]
```

## Make

Slices are built on top of arrays. A slice is dynamic in that it will grow in size. The underlying array, however, does not grow in size. When we create a slice, we can use the built in function make to specify how large our slice should be and also how large the underlying array should be. This can enhance performance a little bit. `make([]T, length, capacity)` `make([]int, 50, 100)` 

## Multi-dimensional slice

A multi-dimensional slice is like a spreadsheet. You can have a slice of a slice of some type. Does that sound confusing? Watch this video and it will all be clarified. `[][]string{[]string{'a', 'b'}, []string{'c', 'd'}, []string{'e', 'f'}}`