# What is scope? How do you determine the scope of a variable in Go?

Scope is the space where a variable is accesible, and to determine it you
only need to see the nearest wrapping curly braces (these denote a block of code),
a variable will be accesible (be in scope) to its containing block of code and all
the nested blocks of code defined after the variable definition. If a variable
is not inside any block of code then it is a package level variable, and is
available anywhere in the package.

```go
var name = "Synthia" // package level variable

func TheFunc() {
  surname := "Mathews"

  for i := 0; i < 22; i++ {
    pets := fmt.Sprintf("only %d", i)
    fmt.Printf("%s %s, do you have pets? %s\n", name, surname, pets)
  }
}

func OtherFunc() {
  fmt.Println(surname) // Error!
}
```
