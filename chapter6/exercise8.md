# How do you assign a value to a pointer?

Using the `&` unary operator on variables.

```go
var name = "Santiago"
var nameptr *string

func main() {
  nameptr = &name
  *nameptr = "サンティアゴ"
  // The name is now different
}
```
