# Functions

A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil.

Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique. If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

The final incoming parameter in a function signature may have a type prefixed with .... A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.

# Declaration


`func (receiver) identifier(parameters) (returns) { code }`
	
know the difference between parameters and arguments:
- we define our func with parameters (if any)
- we call our func and pass in arguments (in any)

Everything in Go is PASS BY VALUE purpose of functions
- abstract code
- code reusability

### Examples:

```
func() {}
func(x int) int {}
func(a, _ int, z float32) bool {}
func(a, b int, z float32) (bool) {}
func(prefix string, values ...int) {}
func(a, b int, z float64, opt ...interface{}) (success bool) {}
func(int, int, float64) (float64, *[]int) {}
func(n int) func(p *T) {}
```

## Variadic parameter

You can create a func which takes an unlimited number of arguments. When you do this, this is known as a “variadic parameter.” When use the lexical element operator “...T” to signify a variadic parameter (there “T” represents some type) `func(x ...int) {}` - unlimited number of int arguments

## Defer

A `defer` statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.  

## Methods

A method is nothing more than a FUNC attached to a TYPE. When you attach a func to a type it is a method of that type. You attach a func to a type with a RECEIVER. 

### Example:

```
func (s string) println() {
    fmt.Println(s)
}
```

## Callback

- **Passing a func as an argument**
- functional programming not something that is recommended in go, however, it is good to be aware of callbacks
- idiomatic go: write clear, simple, readable code

### Example

```
func main() {
	math(23, 7, add)
}

func add(x, y int) {
	fmt.Println(x + y)
}

func math(x int, y int, f func(a, b int)) {
	f(x, y)
}
```
