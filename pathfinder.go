package pathfinder

import (
	"errors"
	"strings"
)

type Node struct {
	Nodes      map[string]*Node
	Variable   *Node
	Leaf       *Leaf
	Extensions map[string]*Leaf
	Wildcard   *Leaf
	Count      int
}

type Leaf struct {
	Value    interface{}
	Segments []string
	Order    int
}

func valid(path string) bool {
	if !blank(path) {
		if strings.HasPrefix(path, "/") {
			return true
		}
	}
	return false
}

func blank(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func split(key string) []string {
	var (
		elements = strings.Split(key, "/")
		n        = 0
	)
	for _, element := range elements {
		if !blank(element) {
			elements[n] = element
			n++
		}
	}
	return elements[:n]
}

func New() *Node {
	return &Node{
		Nodes: make(map[string]*Node),
	}
}

func (n *Node) Add(path string, val interface{}) error {
	if !valid(path) {
		return errors.New("path must begin with /")
	}
	n.Count++
	return n.AddSegment(n.Count, PathToSegments(path), nil, val)
}

func (n *Node) AddLeaf(leaf *Leaf) error {
	extension := GetExtensionFromLastSegment(leaf.Segments)
	if extension != "" {
		if n.Extensions == nil {
			n.Extensions = make(map[string]*Leaf)
		}
		if n.Extensions[extension] != nil {
			return errors.New("duplicate path")
		}
		n.Extensions[extension] = leaf
		return nil
	}

	if n.Leaf != nil {
		return errors.New("duplicate path")
	}
	n.Leaf = leaf
	return nil
}

func (n *Node) AddSegment(order int, segments, wildcards []string, value interface{}) error {
	if len(segments) == 0 {
		return n.AddLeaf(
			&Leaf{
				Order:    order,
				Value:    value,
				Segments: wildcards,
			},
		)
	}
	segment, segments := segments[0], segments[1:]
	if segment == "" {
		return errors.New("empty path segment are not allowed")
	}

	switch segment[0] {
	case ':':
		if n.Variable == nil {
			n.Variable = New()
		}
		return n.Variable.AddSegment(order, segments, append(wildcards, segment[1:]), value)
	case '*':
		if n.Wildcard != nil {
			return errors.New("duplicate path")
		}
		n.Wildcard = &Leaf{
			Order:    order,
			Value:    value,
			Segments: append(wildcards, segment[1:]),
		}
		return nil
	}
	edge, ok := n.Nodes[segment]
	if !ok {
		edge = New()
		n.Nodes[segment] = edge
	}
	return edge.AddSegment(order, segments, wildcards, value)
}

func (n *Node) Find(key string) (leaf *Leaf, expansions []string) {
	if len(key) == 0 || key[0] != '/' {
		return nil, nil
	}
	return n.FindLeaf(PathToSegments(key), nil)
}

func (n *Node) FindLeaf(elements, extensions []string) (leaf *Leaf, expansions []string) {
	if len(elements) == 0 {
		if len(extensions) > 0 && n.Extensions != nil {
			last := extensions[len(extensions)-1]
			prefix, extension := ExtensionForPath(last)
			if leaf := n.Extensions[extension]; leaf != nil {
				extensions[len(extensions)-1] = prefix
				return leaf, extensions
			}
		}
		return n.Leaf, extensions
	}
	var start string
	if n.Wildcard != nil {
		start = strings.Join(elements, "/")
	}
	var element string
	element, elements = elements[0], elements[1:]
	if next, ok := n.Nodes[element]; ok {
		leaf, expansions = next.FindLeaf(elements, extensions)
	}
	if n.Variable != nil {
		_leaf, _expansion := n.Variable.FindLeaf(elements, append(extensions, element))
		if _leaf != nil && (leaf == nil || leaf.Order > _leaf.Order) {
			leaf = _leaf
			expansions = _expansion
		}
	}
	if n.Wildcard != nil && (leaf == nil || leaf.Order > n.Wildcard.Order) {
		leaf = n.Wildcard
		expansions = append(extensions, start)
	}
	return
}

func ExtensionForPath(path string) (string, string) {
	position := strings.LastIndex(path, ".")
	if position != -1 {
		return path[:position], path[position:]
	}
	return "", ""
}

func PathToSegments(key string) []string {
	elements := strings.Split(key, "/")
	if elements[0] == "" {
		elements = elements[1:]
	}
	if elements[len(elements)-1] == "" {
		elements = elements[:len(elements)-1]
	}
	return elements
}

func GetExtensionFromLastSegment(segments []string) string {
	if len(segments) == 0 {
		return ""
	}
	last := segments[len(segments)-1]
	prefix, extension := ExtensionForPath(last)
	if extension != "" {
		segments[len(segments)-1] = prefix
	}
	return extension
}
