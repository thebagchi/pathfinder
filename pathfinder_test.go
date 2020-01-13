package pathfinder

import (
	"fmt"
	"testing"
)

func TestPathFinder(t *testing.T) {
	node := New()
	{
		err := node.Add("/:x", "/:x")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		err := node.Add("/:x/:y", "/:x/:y")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		err := node.Add("/:x/:y/:z", "/:x/:y/:z")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		err := node.Add("/hello/world", "/hello/world")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		err := node.Add("/hello/:x", "/hello/:x")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		err := node.Add("/api/:x", "/api/:x")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		err := node.Add("/api/1", "/api/1")
		if nil != err {
			fmt.Println("Error: ", err)
		}
	}
	fmt.Println("----")
	{
		v, e := node.Find("/hello/world")
		if nil != v && len(e) == 0 {
			fmt.Println(v.Value)
		}
	}
	{
		v, e := node.Find("/hello/:x")
		if nil != v && len(e) != 0 {
			fmt.Println(v.Value)
		}
	}
	{
		v, e := node.Find("/api/1")
		if nil != v && len(e) == 0 {
			fmt.Println(v.Value)
		}
	}
	{
		v, e := node.Find("/api/:x")
		if nil != v && len(e) != 0 {
			fmt.Println(v.Value)
		}
	}
	{
		v, e := node.Find("/x")
		if nil != v && len(e) != 0 {
			fmt.Println(v.Value)
		}
	}
	{
		v, e := node.Find("/x/y")
		if nil != v && len(e) != 0 {
			fmt.Println(v.Value)
		}
	}
	{
		v, e := node.Find("/x/y/z")
		if nil != v && len(e) != 0 {
			fmt.Println(v.Value)
		}
	}
}
