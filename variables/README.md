# Variable in GO Lang

The Go language is static type programing language, that mean when a variable is **DECLARED** the certain **TYPE can't be changed**. There is two ways how a variable can be initialized: 
- `var z string = "value"` (use for global variables that are outside of function body)
- `z := "value"` (use for local variable into function body)

>If a variable doesn't have defined value on declaration it will have **Zero Value**, please take a look for more info on [HERE](https://golang.org/ref/spec#The_zero_value)

>The best practice: keep scope as “narrow” as possible