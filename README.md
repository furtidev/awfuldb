# awfuldb
## Simple key-value database written in Go (using JSON)

### Example
```go
package main

import (
  "fmt"
  "log"
  
  db "github.com/furtidev/awfuldb"
)

func main() {
  obj := db.Database{DBPath="db.json"}
  obj.Set("key", "value")
  key, err := obj.Get("key")
  if err != nil {
    log.Fatalf("ERROR: %s\n", err)
  }

  fmt.Println(key) // value
}
```

### Inspiration
Inspired by [ForgiveDB](https://github.com/hui-z/ForgiveDB/).

### License
This project is licensed under the [MIT](./LICENSE) license.