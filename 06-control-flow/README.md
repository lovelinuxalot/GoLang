# Control Flow
## Understanding control flow
Computers read programs in a certain way, just like we read books in a certain way. When we, as humans, read books, in Western culture, we read from front to back, left to right, top to bottom. When computers read code in a program, they read it from top to bottom. This is known as reading it in SEQUENCE. This is also known as SEQUENTIAL control flow. In addition to sequential control flow, there are two other statements which can affect how a computer reads code. A computer might hit a LOOP control flow. If it hits one of these, it will loop over the code in a specified manner. Loop control flow is also known as ITERATIVE control flow. Finally there is also CONDITIONAL control flow. When the computer hits a CONDITION, like an “if statement” or a “switch statement” the computer will take some course of action depending upon some condition. So the three ways a computer reads code are: SEQUENTIAL, LOOP, CONDITIONAL
- #### Sequence
- #### Loop / iterative
  - for loop
    - init, cond, post
  - bool (while-ish)
    - infinite
  - do-while-is
    - break
  - continue
  - nested
- #### Conditionals
  - switch / case / default statements
    - no default fall-through
    - creating fall-through
    - multiple cases
    - cases can be expressions
      - evaluate to true, they run
    - type
  - if
    - the not operator
      - !
    - initialization statement
    - if / else
    - if / else if / else
    - if / else if / else if / … / else
