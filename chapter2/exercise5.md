# What's the value of the following expression?

```go
expr := (true && false) || (false && true) || !(false && false)
fmt.Println(expr) // ???
```

The value is `true`.
