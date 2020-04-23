# mfile-consul - Consul provider for mfile

Consul provider for github.com/voytechnology/mfile

## Usage

```go
package main

import (
    "github.com/voytechnology/mfile"
    _ "github.com/voytechnology/mfile-consul"
)

func main() {
    // You should now be able to read data from consul.
    v, _ := mfile.ReadFile("consul://foo")

    // ...
}
```