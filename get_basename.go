package main

import (
	"fmt"
	"strings"
)

// method 1
func get_basename(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			name = name[i+1:]
			break
		}
	}

	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			name = name[:i]
			break
		}
	}
	return name
}

// method 2
func get_basename2(name string) string {
	length := strings.LastIndex(name, "/")
	name = name[length+1:]
	if dot := strings.LastIndex(name, "."); dot != -1 {
		name = name[:dot]
	}

	// fmt.Println(name)
	return name
}

func main() {
	fmt.Println(get_basename2("a/b/c.go")) // "c"
	fmt.Println(get_basename2("c.d.go"))   // "c.d"
	fmt.Println(get_basename2("abc"))      // "abc"

}
