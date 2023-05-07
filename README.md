# gin-ping

The tiniest health check library for Gin.

Add health check routes with one line of code:

```go
ginping.Add(r)
```

Example:

```go
import "github.com/epomatti/gin-ping"

func main() {
	r := gin.Default()

  // Adds the health check routes
  ginping.Add(r)

	r.Run()
}
```

Testing it:

```sh
$ curl localhost:8080/health
OK
```
