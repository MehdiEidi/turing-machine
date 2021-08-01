# Turing Machine
Turing machine simulator and input acceptance checker.

![turing machine](https://www.sciencealert.com/images/2019-06/turing-machine.jpg)


# Sample Input
```
Transition Rules:
q0,a=q1,b,R
q1,a=q1,b,R
q1,b=q2,_,L
end

Final States:
q2
end

Initial State:
q0

Tape Content:
aab
```

# Output
```
Accepted, Machine halted in a final state.
```

# Notes:
* All states have a q in the beginning. Like: q1, q2
* Free to use any alphabet. The alphabet will be detected automatically.
* Use 'end' to tell the program that you are finished inputting transition rules and final states.
* Underline '_' is for blank character.

# Requirements 
* In order to run the program you only need to run the executable file (.exe) attached.
* For running the code you will need Go installed in your machine.

# Authors
Mehdi Eidi (mehdiadq@gmail.com)

