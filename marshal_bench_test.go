package utils

import (
    "fmt"
    "time"
    "testing"
    "encoding/json"
    "github.com/legion-zver/utils/unixtime"
)

type ut struct {
    Time unixtime.Time `json:"time"`
}

type tt struct {
    Time time.Time `json:"time"`
}

func TestBenchmarkMarshalJSON(t *testing.T) {
    now     := time.Now()    
    ts := ut{ Time: unixtime.Now() }
    for i := 0; i < 1000000; i++ {
        json.Marshal(&ts)
    }
    fmt.Println("Use unixtime.Time by 1000000 items:\t", time.Now().Sub(now))
    now = time.Now()
    js := tt{ Time: now }
    for i := 0; i < 1000000; i++ {
        json.Marshal(&js)
    }
    fmt.Println("Use time.Time by 1000000 items:\t\t", time.Now().Sub(now))
}
