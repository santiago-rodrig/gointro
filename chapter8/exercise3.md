# What is a package alias? How do you make one?

A package alias is a way for changing the package identifier
that would be normally used for accesing the package
exported identifiers. The way you create one is by placing
the alias right before the import path.

```go
import fmtalias "fmt"
```
