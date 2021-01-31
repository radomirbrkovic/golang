# Struct

A struct is an composite data type. (composite data types, aka, aggregate data types, aka, complex data types). **Structs allow us to compose together values of different types.**

All properties in struct must be initialized

[Official documentation](https://golang.org/ref/spec#Struct_types)

## Declaration 

The struct is declared via the following statement:
 ```
 type <struct-name> struct 
 {
     key type
 }
 ```

 ## Embedded (nested) structs

We can **take one struct and embed it in another struct**. When you do this the inner type gets promoted to the outer type. 

### Example

```
type person struct {
    firstName string
    lastName string
}

type basketballPlayer struct {
    person
    position string
}

....

// Set up VALUE of basketballPlayer TYPE
player := basketballPlayer {
    person:  person {
		firstName: "Micheal",
		lastName: "Jordan",
	},
    position: "Outfielder, Small forward, Shooting guard",
}

```

## Anonymous structs

We can create anonymous structs also. **An anonymous struct is a struct which is not associated with a specific identifier**.

### Example

```
    myStruct := struct {
		firstName string
		lastName  string
	}{
		firstName: "Micheal",
		lastName: "Jordan",
	}
```

## Is go an object oriented language?

Answer on this question is YES and NO, yes because golang provides all that provide other OOP languages, but NO because of in golang all is about TYPE and it hasn't classes and objects.

We create TYPES in Go; user-defined TYPES. We can then have VALUES of that type. We can assign VALUES of a user-defined TYPE to VARIABLES. Anecdote: makes me think of that song, “It’s all about the bass, all about the bass” except “it’s all about the TYPE, all about the TYPE”

More about that topic you can read [on here](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)


### Go is Object Oriented

1. Encapsulation
    - state ("fields")
    - behavior ("methods")
    - exported & unexported; viewable & not viewable
2. Reusability
    - inheritance ("embedded types")
3. Polymorphism
    - interfaces
4. Overriding
    -  "promotion"   
