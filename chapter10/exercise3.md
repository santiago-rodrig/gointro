# What is a buffered channel? How would you create one with a capacity of 20?

A buffered channel is one that only blocks when its maximum capacity has been reached (trying to send blocks),
or when its capacity is 0 (trying to receive blocks). To create a buffered channel you use the following
statement:

```go
c := make(chan int, 20)
```
