package gee

import "strings"

// node 节点结构体
type node struct {
	pattern  string  // 待匹配完整路由，例如 `/p/:lang`
	part     string  // 当前节点对应的路由片段，例如 `:lang`
	children []*node // 子节点 例如 [doc, tutorial, intro]
	// 只要该路径片段（part）以 : 或 * 开头，就说明它是一个通配符，不需要精确匹配具体字符串
	isWild bool // 是否模糊匹配，part 含有 : 或 * 时为true
}

// matchChild 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if part == child.part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 递归调用插入，将路由模式拆分成片段（如 "/p/:lang/doc" -> ["p", ":lang", "doc"]），然后从根节点递归插入
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	// 没有子树，则创建
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	// 递归插入
	child.insert(pattern, parts, height+1)
}

// search 递归调用查找，给定请求路径片段数组（如 ["p","go","doc"]），从根节点递归查找。
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
