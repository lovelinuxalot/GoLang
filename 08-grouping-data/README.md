# Grouping data
- **Array**
  - a numbered sequence of elements of a single type
  - does not change in size
  - **used for Go internals; generally not recommended for your code**
  - https://golang.org/ref/spec#Array_types 
- **Slice**
  - built on top of an array
  - **holds values of the same type**
  - changes in size
  - has a length and a capacity
  - https://golang.org/ref/spec#Slice_types 
- **Map**
  - **key / value storage**
  - an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.
  - https://golang.org/ref/spec#Map_types 
- **Struct**
  - a data structure
  - a composite type
  - allows us to collect values of different types together
  - https://golang.org/ref/spec#Struct_types

# Slice - underlying array
Underlying every slice is an array. A slice is actually a data structure which has three parts: 
 - a pointer to an array
 - len
 - cap
In this video, we will explore the relationship between the slice and the underlying array.

### code
- a new underlying array is allocated
  - https://play.golang.org/p/XE9l_on3Va 
- the same array is used for TWO slices
  - https://play.golang.org/p/Ah2t7mgTn4 
- you can also “throw away” the variable and the same thing happens
  - https://play.golang.org/p/NL8p7Z8R3E 
- looking at the LEN & CAP of a slice
  - https://play.golang.org/p/GvEnKl3x-A 
