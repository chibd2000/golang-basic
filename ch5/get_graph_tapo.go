package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var order []string
var visited map[string]bool = make(map[string]bool, 0)

func topo_sort(m map[string][]string) {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// dfs(keys)
	bfs(keys)
}

func dfs(items []string) {
	for _, item := range items {
		if !visited[item] {
			visited[item] = true
			order = append(order, item)
			dfs(prereqs[item])
		}
	}
}

// 使用breadthFirst遍历其他数据结构。比如，topoSort例子中的课程依赖关系（有向图）
func bfs(items []string) {
	var queue []string
	queue = append(queue, items[0])
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if !visited[item] {
			visited[item] = true
			order = append(order, item)
			for k, _ := range prereqs {
				queue = append(queue, k)
			}
		}
	}
}

// 练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）
func topo_sort2(m map[string][]string) {
	var gmap map[string]bool = make(map[string]bool, 0)
	for key := range m {
		gmap[key] = true
	}
	dfs2(gmap)
}

func dfs2(gmap map[string]bool) {
	for item, _ := range gmap {
		if !visited[item] {
			visited[item] = true
			order = append(order, item)
			dfs(prereqs[item])
		}
	}
}

// 问题描述：有一串数字1到5，按照下面的关于顺序的要求，重新排列并打印出来。要求如下：2在5前出现，3在2前出现，4在1前出现，1在3前出现
//  41325
var gmap2 map[string]string = map[string]string{
	"2": "5",
	"3": "2",
	"4": "1",
	"1": "3",
}

//反转数组顺序
func reverse(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

var q []string
var visited_ map[string]bool

func tapo_sort3(k string) {
	if !visited_[k] {
		visited_[k] = true
		q = append(q, k)
		if gmap2[gmap2[k]] != "" {
			q = append(q, gmap2[k])
			tapo_sort3(gmap2[gmap2[k]])
		}
	}
}

func main() {
	for k, _ := range gmap2 {
		fmt.Printf("\t %s\n", k)
		q = make([]string, 0)
		visited_ = make(map[string]bool, 0)
		tapo_sort3(k)
		// reverse(q)
		fmt.Printf("topusort: %v \n", q)
	}

}
