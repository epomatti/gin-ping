# gin-ping

A tiny health check library for Gin.

How to add it:

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

Testing the health check route:

```sh
curl localhost:8080/health
```
