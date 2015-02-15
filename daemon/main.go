inpackage main

import (
  "fmt"
  "github.com/sonkwo/zoro/nexus"
)

func main() {
  a := nexus.Add1(1, 2)
  fmt.Println("main output...", a)
}
