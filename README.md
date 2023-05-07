# GinPing

![Build](https://github.com/epomatti/gin-ping/actions/workflows/go.yml/badge.svg) [![codecov](https://codecov.io/gh/epomatti/gin-ping/branch/main/graph/badge.svg?token=FNZD2AMY6K)](https://codecov.io/gh/epomatti/gin-ping)

GinPing is a small middleware library for [Gin][1] to add health check routes.

## Usage

To enable a default `/health` route:

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

Testing locally: `curl localhost:8080/health`.

[1]: https://github.com/gin-gonic/gin
