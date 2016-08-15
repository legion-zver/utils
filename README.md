# utils - A set of tools that I use in every project

[![Build Status](https://drone.io/github.com/legion-zver/utils/status.png)](https://drone.io/github.com/legion-zver/utils/latest) [![GitHub license](https://img.shields.io/badge/license-AGPL-blue.svg?style=flat-square)](https://raw.githubusercontent.com/legion-zver/utils/master/LICENSE)

My Utils for GoLang

# Installing

```
go get github.com/legion-zver/utils
```

# UnixTime

Support in ORM such as GORM and marshaller to JSON to int64

## Example JSON

```go
import (
    "fmt"
    "encoding/json"
    "github.com/legion-zver/utils/unixtime"
)

type Struct struct {
    Created unixtime.Time `json:"created"`
}

func main() {
    b, err := json.Marshal(&Struct{Created: unixtime.Now()})
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
}
```

Result:
```
{"created":1471270764}
```
