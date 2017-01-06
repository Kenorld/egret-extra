## Middleware information

This folder contains a middleware which enables net/http/pprof.


## Install

```sh
$ go get -u github.com/eject-contrib/middleware/pprof
```

## Usage

```go
package main

import (
	"github.com/eject-contrib/middleware/pprof"
	"bitbucket.org/kenorld/eject-core"
)

func main() {
  
	eject.Get("/", func(ctx *eject.Context) {
		ctx.HTML(eject.StatusOK, "<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	eject.Get("/debug/pprof/*action", pprof.New())

	eject.Listen(":8080")
}

```
