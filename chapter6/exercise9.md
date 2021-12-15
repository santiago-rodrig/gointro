# How do you create a new pointer?

The same way you'll create a new variable of any type,
just that now the type will have `*` right in front of
it.

```go
var nameptr *string // A pointer to a string value
```

You can also use the `&` operator to make the compiler know
that the variable is a pointer type.

```go
var nameptr = &name
```

It is also possible to do create a pointer with the `:=` notation.

```go
func f() {
  ptr := &somePackageVar
}
```

Additionally, we have the builtin function `new`, that takes a type
as argument and returns a pointer to some value in memory with the type
that was passed as argument.

```go
func g() {
  ptr := new(int) // int*
}
```
