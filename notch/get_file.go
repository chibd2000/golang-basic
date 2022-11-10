package main

import (
	"os"
)

func tempDirs() []string {
	return []string{"123", "456"}
}
func main() {
	var rmdirs []func()
	for _, dir := range tempDirs() {
		d := dir // NOTE: necessary!
		os.MkdirAll(d, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(d)
		})
	}
}
