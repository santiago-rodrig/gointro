# What are `defer`, `panic`, and `recover`? How do you recover from a runtime panic?

```go
package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		recover()
		fmt.Println("Recovered in f")
	}()
	g()
	fmt.Println("This never prints")
}

func g() {
	panic(nil)
}
```

## `defer`

Builtin statement that precedes a function call which will be
deferred and added to the call stack right after the end of the
containing function, all deferred function calls are called in
last-in-first-out order.

## `panic`

Builtin function that takes one argument of type `interface{}` that
interrupts the call stack and starts _panicking_ until a deferred
function call _recovers_ from the panic.

## `recover`

Builtin function that returns an `interface{}`, the return value is
not really important, the simple call to this function **in a deferred**
function call will stop the _panicking_ state caused by other functions
down the call stack.
