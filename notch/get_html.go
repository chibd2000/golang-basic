package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func gather_link(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = gather_link(links, c)
	}
	return links
}

// 修改代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用
// 相当于树的线序遍历，根左右的方式进行
func gather_link2(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		gather_link2(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		gather_link2(links, n.NextSibling)
	}

	return links
}

func get_tag(tags []string, node *html.Node) {
	if node.Type == html.ElementNode {
		tags = append(tags, node.Data)
		fmt.Println(tags)
	}

	if node.FirstChild != nil {
		get_tag(tags, node.FirstChild)
	}

	if node.NextSibling != nil {
		get_tag(tags, node.NextSibling)
	}
}

// 记录同类型标签名出现的次数count
func get_count(counts map[string]int, node *html.Node) map[string]int {
	if node.Type == html.ElementNode {
		counts[node.Data]++
	}

	if node.FirstChild != nil {
		counts = get_count(counts, node.FirstChild)
	}

	if node.NextSibling != nil {
		counts = get_count(counts, node.NextSibling)
	}

	return counts
}

func get_http(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	return doc, err
}

func main() {
	doc, err := get_http(os.Args[1])
	if err != nil {
		log.Fatal("get_http error...")
	}

	for _, link := range gather_link(nil, doc) {
		fmt.Println(link)
	}

	// gmap := make(map[string]int)
	// gcount := get_count(gmap, doc)
	// fmt.Println(gcount)

	//get_tag(nil, doc)
}
