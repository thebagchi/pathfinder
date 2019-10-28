package pathfinder

import (
	"fmt"
	"testing"
)

func TestPathFinder(t *testing.T) {
	node := New()
	{
		err := node.Add("/hello/world", "static_path")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	{
		err := node.Add("/hello/world", "duplicate_path")
		if nil != err {
			fmt.Println("Error: ", err)
		}

	}
	{
		err := node.Add("/:x/:y", "dynamic_path")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	{
		err := node.Add("/hello/world/*", "wildcard_path")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	{
		leaf, values := node.Find("/hello/world")
		fmt.Println(leaf, values)
	}

	{
		leaf, values := node.Find("/goodbye/world")
		fmt.Println(leaf.Value)
		fmt.Println(leaf, values)
	}

	{
		leaf, values := node.Find("/hello/goodbye/world")
		//fmt.Println(leaf.Value)
		fmt.Println(leaf, values)
	}

	{
		leaf, values := node.Find("/hello/world/wildcard")
		//fmt.Println(leaf.Value)
		fmt.Println(leaf, values)
	}
}
