# What are the two ways to create a new variable?

With the following statement you can create a variable at
any level in a file.

```go
package omega

var varName string = "Hey there"
var otherVar int
var japaneseName = "サンティアゴ"

func OmegaFunc() {
  var insideFunc = 32_987
}

// And the code continues...
```

The following statement can only be used inside functions.

```go
package omega

// The following statement would generate an error
// data := "Useful stuff"

func UltraFunc() {
  info := map[string]string{
    "date": "Today",
    "time": "Yes",
    "hungry?": "Always",
  }
}

// And the code continues...
```
