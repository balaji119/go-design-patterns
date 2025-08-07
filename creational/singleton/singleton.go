package singleton

import (
    "fmt"
    "sync"
)

type singleton struct {
    value string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        fmt.Println("Creating singleton instance")
        instance = &singleton{value: "I am the only one"}
    })
    return instance
}
