package main

import "fmt"

// 排序二叉树线序遍历构造练习
// root@localhost src % go run get_sort_tree.go
// 1
// 2
// 3
// 4
// 5
type TreeNode struct {
	l_node, r_node *TreeNode
	value          int
}

func add_node(node **TreeNode, node_value int) {
	if *node == nil {
		*node = new(TreeNode)
		(*node).value = node_value
	} else {
		if node_value > (*node).value {
			add_node(&(*node).l_node, node_value)
		} else {
			add_node(&(*node).r_node, node_value)
		}
	}
}

func print_node(node *TreeNode) {
	if node != nil {
		fmt.Println(node.value)
		print_node(node.l_node)
		print_node(node.r_node)
	}
}

func main() {
	var root_node *TreeNode = nil
	node_value := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(node_value); i++ {
		add_node(&root_node, node_value[i])
	}

	print_node(root_node)
}
