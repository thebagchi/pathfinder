package pathfinder

import (
	"strings"
	"testing"
)

var paths = map[string]string{
	"/":            "/",
	"/:a":          "/:a",
	"/:a/:b":       "/:a/:b",
	"/:a/:b/:c":    "/:a/:b/:c",
	"/:a/:b/:c/:d": "/:a/:b/:c/:d",
	"/a":           "/a",
	"/b":           "/b",
	"/c":           "/c",
	"/d":           "/d",
	"/a/:b":        "/a/:b",
	"/a/:b/:c":     "/a/:b/:c",
	"/a/:b/:c/:d":  "/a/:b/:c/:d",
	"/:a/b":        "/:a/b",
	"/:a/c":        "/:a/c",
	"/:a/d":        "/:a/d",
	"/:a/a":        "/:a/a",
	"/:a/b/c":      "/:a/b/c",
	"/:a/c/b":      "/:a/c/b",
	"/:a/b/c/d":    "/:a/b/c/d",
	"/:a/b/d/c":    "/:a/b/d/c",
	"/:a/c/b/d":    "/:a/c/b/d",
	"/:a/c/d/b":    "/:a/c/d/b",
	"/:a/d/c/b":    "/:a/d/c/b",
	"/:a/d/b/c":    "/:a/d/b/c",
	"/a/b":         "/a/b",
	"/a/b/:c":      "/a/b/:c",
	"/a/b/:c/:d":   "/a/b/:c/:d",
	"/:a/:b/c":     "/:a/:b/c",
	"/:a/:b/c/d":   "/:a/:b/c/d",
	"/:a/:b/d/c":   "/:a/:b/d/c",
	"/a/b/c":       "/a/b/c",
	"/a/b/c/:d":    "/a/b/c/:d",
	"/:a/:b/:c/d":  "/:a/:b/:c/d",
	"/:a/:b/:c/a":  "/:a/:b/:c/a",
	"/:a/:b/:c/b":  "/:a/:b/:c/b",
	"/:a/:b/:c/c":  "/:a/:b/:c/c",
	"/a/b/c/d":     "/a/b/c/d",
	"/a/b/d/c":     "/a/b/d/c",
	"/a/c/b/d":     "/a/c/b/d",
	"/a/c/d/b":     "/a/c/d/b",
	"/a/d/c/b":     "/a/d/c/b",
	"/a/d/b/c":     "/a/d/b/c",
	"/b/c/d/a":     "/b/c/d/a",
	"/b/c/a/d":     "/b/c/a/d",
	"/b/d/c/a":     "/b/d/c/a",
	"/b/d/a/c":     "/b/d/a/c",
	"/b/a/c/d":     "/b/a/c/d",
	"/b/a/d/c":     "/b/a/d/c",
	"/c/d/a/b":     "/c/d/a/b",
	"/c/d/b/a":     "/c/d/b/a",
	"/c/b/a/d":     "/c/b/a/d",
	"/c/b/d/a":     "/c/b/d/a",
	"/c/a/b/d":     "/c/a/b/d",
	"/c/a/d/b":     "/c/a/d/b",
	"/d/a/b/c":     "/d/a/b/c",
	"/d/a/c/b":     "/d/a/c/b",
	"/d/b/a/c":     "/d/b/a/c",
	"/d/b/c/a":     "/d/b/c/a",
	"/d/c/a/b":     "/d/c/a/b",
	"/d/c/b/a":     "/d/c/b/a",
}

func TestPathFinder(t *testing.T) {
	node := New()
	for key, value := range paths {
		err := node.Add(key, value)
		if nil != err {
			t.Error("Failed adding path: ", key)
			t.Fail()
		}
	}
	for key, value := range paths {
		err := node.Add(key, value)
		if nil == err {
			t.Error("Duplicate path added: ", key)
			t.Fail()
		}
	}
	for key, value := range paths {
		data, params := node.Find(key)
		if nil == data {
			t.Error("Failed finding path: ", key)
			t.Fail()
		}
		if nil != data {
			if data.Value != value {
				t.Error("Wrong value found for path: ", key)
				t.Error("Expected :", value, " Got: ", data.Value)
				t.Fail()
			}
		}
		if count := strings.Count(key, ":"); count > 0 {
			if len(params) != count {
				t.Error("Wrong params found for path: ", key)
				t.Error("Expected count: ", count, " Got count: ", len(params))
				t.Fail()
			}
		}
	}
	for key, value := range paths {
		path := key
		path = strings.ReplaceAll(path, ":a", "w")
		path = strings.ReplaceAll(path, ":b", "x")
		path = strings.ReplaceAll(path, ":c", "y")
		path = strings.ReplaceAll(path, ":d", "z")
		data, params := node.Find(path)
		if nil == data {
			t.Error("Failed finding path: ", key)
			t.Fail()
		}
		if nil != data {
			if data.Value != value {
				t.Error("Wrong value found for path: ", key)
				t.Error("Expected :", value, " Got: ", data.Value)
				t.Fail()
			}
		}
		if count := strings.Count(key, ":"); count > 0 {
			if len(params) != count {
				t.Error("Wrong params found for path: ", key)
				t.Error("Expected count: ", count, " Got count: ", len(params))
				t.Fail()
			}
		}
	}
}
