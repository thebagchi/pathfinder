package pathfinder

import (
	"errors"
	"strings"
)

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

func segments(key string) []string {
	elements := strings.Split(key, "/")
	if elements[0] == "" {
		elements = elements[1:]
	}
	if elements[len(elements)-1] == "" {
		elements = elements[:len(elements)-1]
	}
	return elements
}

type Node struct {
	Children   map[string]*Node
	Parameters *Node
	Value      *Leaf
	Count      int
}

type Leaf struct {
	Value      interface{}
	Parameters []string
	Order      int
}

func New() *Node {
	return &Node{
		Children: make(map[string]*Node),
	}
}

func (n *Node) Add(path string, val interface{}) error {
	if !valid(path) {
		return errors.New("path must begin with /")
	}
	n.Count++
	return n.AddSegment(n.Count, segments(path), nil, val)
}

func (n *Node) AddLeaf(leaf *Leaf) error {
	if n.Value != nil {
		return errors.New("duplicate path")
	}
	n.Value = leaf
	return nil
}

func (n *Node) AddSegment(order int, segments, parameters []string, value interface{}) error {
	if len(segments) == 0 {
		return n.AddLeaf(
			&Leaf{
				Order:      order,
				Value:      value,
				Parameters: parameters,
			},
		)
	}
	segment, segments := segments[0], segments[1:]
	if segment == "" {
		return errors.New("empty path segment are not allowed")
	}

	switch segment[0] {
	case ':':
		if n.Parameters == nil {
			n.Parameters = New()
		}
		return n.Parameters.AddSegment(order, segments, append(parameters, segment[1:]), value)
	default:
		edge, ok := n.Children[segment]
		if !ok {
			edge = New()
			n.Children[segment] = edge
		}
		return edge.AddSegment(order, segments, parameters, value)
	}
}

func (n *Node) Find(key string) (leaf *Leaf, expansions []string) {
	if len(key) == 0 || key[0] != '/' {
		return nil, nil
	}
	return n.FindLeaf(segments(key), nil)
}

func (n *Node) FindLeaf(elements, parameters []string) (leaf *Leaf, expansions []string) {
	if len(elements) == 0 {
		return n.Value, parameters
	}
	element, elements := elements[0], elements[1:]
	if next, ok := n.Children[element]; ok {
		leaf, expansions = next.FindLeaf(elements, parameters)
	}
	if n.Parameters != nil {
		_leaf, _expansion := n.Parameters.FindLeaf(elements, append(parameters, element))
		if _leaf != nil && leaf == nil {
			leaf = _leaf
			expansions = _expansion
		}
	}
	return
}
