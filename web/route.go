package web

import (
	"fmt"
	"strings"
)

type router struct {
	trees map[string]*node
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

func (r *router) addRoute(method string, path string, handler HandleFunc) {

	if path == "" {
		panic("web: 路由是空字符串")
	}

	if path[0] != '/' {
		panic("web: 路由必须以 / 开头")
	}

	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 路由不能以 / 结尾")
	}

	root, ok := r.trees[method]

	if !ok {
		root = &node{path: "/"}
		r.trees[method] = root
	}

	if path == "/" {
		if root.handler != nil {
			panic("web: 路由冲突[/]")
		}

		root.handler = handler
		return
	}

	segs := strings.Split(path[1:], "/")

	for _, s := range segs {
		if s == "" {
			panic(fmt.Sprintf("web: 非法路由。不允许使用 //a/b, /a//b 之类的路由, [%s]", path))
		}
		root = root.childOrCreate(s)
	}

	if root.handler != nil {
		panic(fmt.Sprintf("web: 路由冲突[%s]", path))
	}

	root.handler = handler

}

// findRoute 查找对应的节点
// 注意，返回的 node 内部 HandleFunc 不为 nil 才算是注册了路由
func (r *router) findRoute(method string, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}

	if path == "/" {
		return root, true
	}

	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, s := range segs {
		root, ok = root.childOf(s)
		if !ok {
			return nil, false
		}
	}
	return root, true
}

type node struct {
	path string

	// 静态匹配的节点
	// 子 path 到子节点的映射
	children map[string]*node

	// 加一个通配符匹配
	starChild *node

	handler HandleFunc
}

// childOrCreate 查找子节点，如果子节点不存在就创建一个
// 并且将子节点放回去了 children 中
func (n *node) childOrCreate(seg string) *node {

	if seg == "*" {
		n.starChild = &node{
			path: seg,
		}
		return n.starChild
	}

	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[seg]
	if !ok {
		// 要新建一个
		child = &node{path: seg}
		n.children[seg] = child
	}
	return child
}

// childOf 优先考虑静态匹配，匹配不上，再考虑通配符匹配
func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return n.starChild, n.starChild != nil
	}
	child, ok := n.children[path]

	if !ok {
		return n.starChild, n.starChild != nil
	}
	return child, ok
}
