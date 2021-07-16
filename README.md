# air

## Showcase

### Hello world
```
package air

phrase: String
    blow "Hello"
    blow "World"


print(phrase ->)
```

### For-loop
Forget for loops... Try use the `consumer` instead!
```
>>> print(5 ->)
1 2 3 4 5
```

### While-loop
Do you need a while loop? Okay... it can help you :D.
```
package while-loop

air(n) (if n < 5)
    return air(n+1)

air(n)  # "Else" case.
    return n
```
```
>>> air(0)
5
```
