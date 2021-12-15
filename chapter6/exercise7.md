# How do you get the memory address of a variable?

With the `&` unary operator, that _gets_ the address
of the variable that it operates on.

```go
var omega = "WarnTheForeigners"
var theta = &omega

func main() {
  *theta = "WissleTheFrench"
}
```
