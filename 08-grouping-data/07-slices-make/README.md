# Slice - make
- Slices are built on top of arrays. 
- A slice is dynamic in that it will grow in size.
- The underlying array, however, does not grow in size. 
- When we create a slice, **we can use the built in function make to specify how large our slice should be and also how large the underlying array should be.**
- This can enhance performance a little bit.
  - make([]T, length, capacity) 
  - make([]int, 50, 100)
- if the capacity of the slice is full, and when a new value is added,
  - a new array is created with **double the capacity of the old array**
  - the **contents of the old array is copied to the new one**
  - the old array gets deleted
  - this is the reason there is performance issues when creating dynamic arrays

### code
- https://play.golang.org/p/07hH1b-hvD 