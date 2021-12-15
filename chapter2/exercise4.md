# What is a string? How do you find its length?

A string is a sequence of bytes (uint8),
that's the way they're defined in Go, to find the length of a string
you use the builtin function `len`, that returns the number of bytes
in the string.

```go
str := "こんにちはみなさん、げんきですか？"
fmt.Println(len(str)) // 51! This is the number of bytes held by the string
```
