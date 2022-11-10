package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal("http.Get error...")
	}

	if response.StatusCode != http.StatusOK {
		log.Fatal("error...")
	}
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	response.Body.Close()

}
