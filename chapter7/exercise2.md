# Why would you use an embedded anonymous field instead of a normal named field?

From now own, a user-defined type whose underlying type is a `struct` will be referred
as a `struct`.

When a `struct` augments a pre-defined `struct` it makes sense to use _type embedding_.

```go
type Circle struct {
  Radius int
}

type Wheel struct {
  Spokes int
  Circle
}

func main() {
  wheel := Wheel{Spokes: 33}
  // The following would be if you wanted
  // to define the Circle properties right
  // in the Wheel struct literal.
  // wheel := Wheel{
  //   Spokes: 33,
  //   Circle: Circle{
  //     Radius: 22,
  //   },
  // }
  wheel.Radius = 22
}
```
