# PathFinder
Simple http router library.

## Sample Usage:
```go
package main
import (
  "github.com/thebagchi/pathfinder"
)

func VersionHandler() {
  // Handle Version
}

func ParamHandler() {
  // Handle Parameter
}

func main() {
  node := pathfinder.New()
  _ = node.Add("/api/v1/version", VersionHandler)
  _ = node.Add("/api/v1/:param", ParamHandler)
  _ = node.Add("/api/v2/{param}", ParamHandler)
  {
    data, params := node.Find("/api/v1/version")
    _ = params // params is array of string containing dynamic parts
    if handler, ok := data.Value.(func()); ok {
        handler()
    }
  }
  // Similarly other paths can be handled 
}
```
