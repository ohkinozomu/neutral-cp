# neutral-cp

neutral-cp is facade for continuous profiling.

## Example

```go
package main

import (
	"net/http"
	ncp "github.com/ohkinozomu/neutral-cp"
)

func startProfiler() {
	c := ncp.Config{
		Registry:        ncp.CLOUD_PROFILER,
		ApplicationName: "example-app",
	}
	ncp := ncp.NeutralCP{Config: c}

	ctx := context.Background()
	err := ncp.Start(ctx)
	if err != nil {
		panic(err)
	}
}

func main() {
	go startProfiler()
	http.ListenAndServe(":8080", nil)
}
```

## Registries

- Pyroscope

- Cloud Profiler