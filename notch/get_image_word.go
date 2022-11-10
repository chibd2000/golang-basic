package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// 实现countWordsAndImages分词
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images += 1
		} else if n.Data == "div" {
			words += 1
		}
	}

	if n.FirstChild != nil {
		a, b := countWordsAndImages(n.FirstChild)
		words += a
		images += b
	}

	if n.NextSibling != nil {
		a, b := countWordsAndImages(n.NextSibling)
		words += a
		images += b
	}

	return words, images
}

func main() {
	fmt.Println(CountWordsAndImages("http://www.baidu.com"))
}
