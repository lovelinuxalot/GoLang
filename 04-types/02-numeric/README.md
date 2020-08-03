# Numeric types
## Integers
- numbers without decimals
  - aka, whole number
- int & uint
  - "implementation-specific sizes"
- all numeric types are distinct except 
  - ***byte*** which is an alias for ***uint8***
  - ***rune*** which is an alias for ***int32***
- types are unique
  - "this is static programming language"
  - "Conversions are required when different numeric types are mixed in an expression or assignment. For instance, int32 and int are not the same type even though they may have the same size on a particular architecture” [source](https://golang.org/ref/spec#Numeric_types)"
- rule of thumb: just use ***int***
## Floating point
- numbers with decimals
  - aka, real numbers
- rule of thumb: just use ***float64*** 
## Nice reading
 - [Caleb Doxsey’s book](https://www.golang-book.com/books/intro/3)
### code:
- https://play.golang.org/p/OdWUH8uva6
- https://play.golang.org/p/0JpmCYezs1 
- this does not run: https://play.golang.org/p/O7nFEn8nXz 
- int8
  - https://play.golang.org/p/IcOtgm6YKA 
  - does not work: https://play.golang.org/p/YbwTa1YT4i 
  - https://play.golang.org/p/exwG0ijjRf 
  - does not work: https://play.golang.org/p/sy16rgifWF 
## Runtime package
- GOOS
- GOARCH
https://play.golang.org/p/1vp5DImIMM 