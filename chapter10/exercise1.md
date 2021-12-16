# How do you specify the direction of a channel type?

By using `<-` before or after the `chan` type keyword.

```go
func receivesFromChannel(c <-chan string) {
  // ...
}

func sendsToChannel(c chan<- string) {
  // ...
}
```
