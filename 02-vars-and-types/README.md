# Variables, values, & type

## Playground

URL: https://play.golang.org/

### Features
- share
- import
- format
- run

### Forum to share code
 - https://forum.golangbridge.org/ 
   
### **“idiomatic go”**
- idioms are patterns of speech
- when someone writes “idiomatic Go” they are writing Go code in the way Go code community writes code
---
## Introduction to packages
- ***variadic parameters***
  - the “...<some type>” is how we signify a variadic parameter
  - the type “interface{}” is the empty interface
    - every value is also of type ***"interface{}"***
  - so the parameter "...interface{}" means “give me as many arguments of any type as you’d like 
- throwing away returns
  - use ***the "_" underscore character*** to throw away returns
- you can’t have unused variables in your code
  - this is code pollution
  - the compiler doesn’t allow it
- we use this notation in Go
  - ***package.Identifier***
    - ex: fmt.Println()
      - we would read that: “from package fmt I am using the Println func”
  - an identifier is the name of the variable, constant, func
- packages
  - code that is already written which you can use
  - imports

code: https://play.golang.org/p/30rOKO3qT8s

---
