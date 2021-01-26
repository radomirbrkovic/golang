# Map

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is nil.

A map is a key-value store. This means that you store some value and you access that value by a key.
For instance, I might store the value “phoneNumber” and access it be the key “friendName”.
This way I could enter my friend’s name and have the map return their phone number. 
The syntax for a map is map[key]value. The key can be of any type which allows comparison 
(eg, I could use a string or an int, for example, as I can compare if two strings are equal,
or if two ints are equal). It is important to note that maps are unordered. 
If you print out all of the keys and values in a map, they will print out in random order.  
A map is the perfect data structure when you need to look up data fast.

[Official documentation](https://golang.org/ref/spec#Map_types)

# Declaration

`map[KEY_TYPE]VALUE_TYPE{KEY:VALUE,}`

# Insert and update

In a map the element can be easy added or updated via the following syntax: `MAP_VARIABLE[KEY] = VALUE`

# Delete

You delete an entry from a map using `delete(<map name>, “key”)`. No error is thrown if you use a key which does not exist. To confirm you delete a key/value, verify that the key/value exists with the comma ok idiom.
