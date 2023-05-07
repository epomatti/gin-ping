# GinPing

A tiny health check library for Gin.

## How to

```go
ginping.Add(r)
```

Example:

```go
package main

import (
  "github.com/epomatti/gin-ping"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  ginping.Add(r)
  r.Run()
}
```

Test it:

```sh
curl localhost:8080/health
```
