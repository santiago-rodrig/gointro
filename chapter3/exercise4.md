# What is the difference between `var` and `const`?

Variables can change their value, constans can't, and
constants can only have values that can be determined
at compile-time.

```go
const name = "Santiago"
const address = GetAddress() // Error!
var hobby = GetHobby("Getting into")

func Life() {
  hobby = GetHobby("Respecting")
}

func GetHobby(prefix string) string {
  return prefix + " other people's conversations"
}
```
