package main

import (
	"fmt"
	"sync"
)

var Cache = struct {
	mappings map[string]string
	mutex    sync.Mutex
}{
	mappings: make(map[string]string),
}

func get_cache(key string) string {
	Cache.mutex.Lock()
	defer Cache.mutex.Unlock()
	if t := Cache.mappings[key]; t != "" {
		return t
	}

	return ""
}
func main() {
	t := get_cache("haha")
	if t == "" {
		fmt.Println("not get...")
	} else {
		fmt.Println("get ...")
	}
}
