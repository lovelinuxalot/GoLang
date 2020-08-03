# Hands-on exercise #1
- Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x" and "y" and "z"
  - 42
  - "James Bond"
  - true
- Now print the values stored in those variables using 
  - a single print statement
  - multiple print statements
### code:
here’s the solution: https://play.golang.org/p/yYXnWXIQNa
 
---
# Hands-on exercise #2
- Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
  - identifier "x" type int
  - identifier "y" type string
  - identifier "z" type bool
- in func main
  - print out the values for each identifier
  - The compiler assigned values to the variables. What are these values called?
###code:
here’s the solution: https://play.golang.org/p/jzHwSlles9 

---
# Hands-on exercise #3
- Using the code from the previous exercise,
  - At the package level scope, assign the following values to the three variables
    - for a assign 42
    - for b assign "James Bond"
    - for c assign true
  - in func main
    - use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER "s"
    - print out the value stored by variable "s"
###code:
here’s the solution: https://play.golang.org/p/QFctSQB_h3 

---
# Hands-on exercise #4
- FYI - nice documentation and new terminology "underlying type"
  - https://golang.org/ref/spec#Types 
- For this exercise
  - Create your own type. Have the underlying type be an int.
  - create a VARIABLE of your new TYPE with the IDENTIFIER "d" using the "VAR" keyword
  - in func main
    - print out the value of the variable "d"
    - print out the type of the variable "d"
    - assign 42 to the VARIABLE "d" using the "=" OPERATOR
    - print out the value of the variable "d"
###code:
here’s the solution: https://play.golang.org/p/snm4WuuYmG 

---
#Hands-on exercise #5
- Building on the code from the previous example
  - at the package level scope, using the "var" keyword, create a VARIABLE with the IDENTIFIER "f". The variable should be of the UNDERLYING TYPE of your custom TYPE "e"
    - eg:
    ```
    type hotdog int
    var e hotdog
    var f int
    ```
  - in func main
    - this should already be done
      - print out the value of the variable "e"
      - print out the type of the variable "e"
      - assign your own VALUE to the VARIABLE "e" using the "=" OPERATOR
      - print out the value of the variable "e"
    - now do this
      - now use CONVERSION to convert the TYPE of the VALUE stored in "e" to the UNDERLYING TYPE
        - then use the "=" operator to ASSIGN that value to "f"
        - print out the value stored in "f"
        - print out the type of "f"
###code:
here’s the solution: https://play.golang.org/p/cj8RrYgBOD 

---
#Hands-on exercise #6
[Take this quiz.](https://goo.gl/forms/dfwmTuYTe5ox8nyt1)
