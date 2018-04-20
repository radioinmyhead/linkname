package a

import (
	_ "unsafe"

	_ "github.com/radioinmyhead/linkname/b"
)

func Show(s string) string

func Hello(s string) string {
	return hello(s)
}

//go:linkname hello a.hello
func hello(s string) string
