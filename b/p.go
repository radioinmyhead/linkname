package b

import (
	_ "unsafe"
)

//go:linkname say1 github.com/radioinmyhead/linkname/a.Show
func say1(s string) string {
	return "show, " + s
}

//go:linkname say2 a.hello
func say2(s string) string {
	return "hello, " + s
}
