# How would you document an exported function in a package?

With a comment right above the function signature explaining
the behavior of the function.

```go
package stuffs

type Stuff struct {
  Name string
  Purpose string
}

// Gets some stuffs.
func GetStuffs() []Stuff {
  return []Stuff{Stuff{Name: "Kerchief", Purpose: "Cleaning"}, Stuff{Name: "Fork", Purpose: "Eating"}}
}
```
