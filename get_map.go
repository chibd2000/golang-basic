package main

import "fmt"

func main() {

	// first method
	tmap := map[string]int{
		"test2": 2,
		"test":  1,
	}

	// second method
	tmap2 := make(map[string]int)
	tmap2["test"] = 1
	tmap2["test2"] = 2
	fmt.Println(tmap2)

	// third method
	tmap3 := map[string]int{}
	tmap3["test"] = 1
	tmap3["test2"] = 2
	fmt.Println(tmap3)

	// print map
	for k, v := range tmap {
		fmt.Println("k: ", k, " v: ", v)
	}
	fmt.Println("------------")

	delete(tmap, "test3")
	fmt.Println(tmap)
	delete(tmap, "test")
	fmt.Println(tmap)
}
