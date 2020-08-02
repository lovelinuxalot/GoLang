# The fmt package

### basic code setup
 - using var
  - using zero value
 - using short declaration operator
 - https://play.golang.org/p/oLxC8gHE0D 
### printing default value
 - %v
 - https://play.golang.org/p/51WOG55riI 
### escaped characters
 - like \n or \t
 - https://golang.org/ref/spec#Rune_literals 
### format printing
 - https://play.golang.org/p/rqGVNVP5kl 
 - we look at the difference between these funcs in the “fmt” package
   - ####group #1 
     - general printing to standard out
     - [func Print(a ...interface{}) (n int, err error)](http://godoc.org/fmt#Print)
     - [func Printf(format string, a ...interface{}) (n int, err error)](http://godoc.org/fmt#Printf)
     - [func Println(a ...interface{}) (n int, err error)](http://godoc.org/fmt#Println)
   - ####group #2
     - printing to a string which you can assign to a variable
     - [func Sprint(a ...interface{}) string](http://godoc.org/fmt#Sprint)
     - [func Sprintf(format string, a ...interface{}) string](http://godoc.org/fmt#Sprintf)
     - [func Sprintln(a ...interface{}) string](http://godoc.org/fmt#Sprintln)
   - ####group #3
     - printing to a file or a web server’s response
     - [func Fprint(w io.Writer, a ...interface{}) (n int, err error)](http://godoc.org/fmt#Fprint)
     - [func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)](http://godoc.org/fmt#Fprintf)
     - [func Fprintln(w io.Writer, a ...interface{}) (n int, err error)](http://godoc.org/fmt#Fprintln)

###code:
https://play.golang.org/p/o36sAgq5BkT 
