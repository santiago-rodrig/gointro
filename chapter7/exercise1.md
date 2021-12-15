# What's the difference between a method and a function?

A method is a function with a user-defined type receiver, a function
is simply a function, they don't have a receiver. Methods are called
on user-defined types with dot notation (a.b).

```go
type MyInt int

// Method
func (mi MyInt) double() MyInt {
  return mi * 2
}

// Function
func doubleMyInt(mi MyInt) MyInt {
  return mi * 2
}
```
