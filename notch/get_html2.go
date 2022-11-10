package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var depth int
var tag_sign int = 1

func pre(n *html.Node) {
	if n.Type == html.ElementNode {
		depth++
		if n.FirstChild != nil {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, v := range n.Attr {
				fmt.Printf(" %s=\"%s\"", v.Key, v.Val)
			}
			fmt.Println(">")
		} else {
			fmt.Printf("%*s<%s/>", depth*2, "", n.Data)
		}
	}
}

func post(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		} else {
			fmt.Printf("\n")
		}
		depth--
	}
}

func gather_html(n *html.Node, pre func(*html.Node), post func(*html.Node)) {
	pre(n)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		gather_html(c, pre, post)
	}

	post(n)
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

// 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，
// 根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
func get_element_by_id(doc *html.Node, id string) *html.Node {
	return nil
}

// 编写函数expand，将s中的"foo"替换为f("foo")的返回值
func expand(s string, f func(string) string) string {
	t := f("foo")
	q := strings.Join(strings.Split(s, "foo"), t)
	return q
}

func main() {
	// doc, err := get_http("http://www.baidu.com")
	// if err != nil {
	// 	log.Fatal("get_http error...")
	// }

	// gather_html(doc, pre, post)
	s := expand("this_is_foo_haha", func(t string) string {
		return strings.Map(func(a rune) rune {
			return a + 1
		}, t)
	})
	fmt.Println(s)
}
