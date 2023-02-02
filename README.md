# Terminal colors for golang


## Usage

```go
package main

import (
  "fmt"
  color "github.com/dghaehre/termcolor"
)

func main() {
  color.Println(color.Green, "this is green")
  fmt.Println(color.Str(color.Green, "this is also green"))
}
```
