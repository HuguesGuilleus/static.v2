# static.v2

[![GoDoc](https://godoc.org/github.com/HuguesGuilleus/static.v2?status.svg)](https://godoc.org/github.com/HuguesGuilleus/static.v2)

Create easily a http Handler for static file.

## Installation

```bash
go get -u github.com/HuguesGuilleus/static.v2
```

## Example

```go
package main

import (
	"github.com/HuguesGuilleus/static.v2"
	"log"
	"net/http"
)

func main() {
	// To pass in Dev mode.
	// static.Dev = true

	http.HandleFunc("/", static.Html(nil, "front/index.html"))
	http.HandleFunc("/style.css", static.Css(nil, "front/style.css"))
	http.HandleFunc("/app.js", static.Js(nil, "front/app.js"))
	http.HandleFunc("/favicon.png", static.Png(nil, "front/favicon.png"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
