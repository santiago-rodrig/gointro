# What is the value of `x` after running this program?

```go
package main

func square(x *float64) {
  *x = *x * *x
}

func main() {
  x := 1.5
  square(&x)
}
```

The value of `x` will be 9/4, because an assignment to a _dereference_ is being made on the
pointer to the value that was passed as argument to the `square` function, this will
change the value that the pointer is pointing to, effectively modifying it after
calling `square`.
