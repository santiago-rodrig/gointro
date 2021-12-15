# How are integers stored on a computer?

As sequences of bits (1 or 0), integers can hold a sequence of at most 64 bits,
leaving one bit to be the sign in case of signed integers. A 64 bits signed
integer ranges from 2^63 to 2^63-1, and you can guess the ranges for 32 bits signed integers.
Unsigned integers take ranges that start from 0, a 64 bits unsigned integer goes
from 0 to 2^64-1, guess what the top number would be for 32 bits unsigned integers.

```go
var a int8 = 16 // 2^7 to 2^7-1
var b int16 = 31799 // 2^15 to 2^15-1
var c uint64 = 12315123123 // 0 to 2^64-1
```
