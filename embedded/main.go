package main

import (
	_ "embed"
)

//go:embed hello.txt
var b []byte

func main() {
	print(string(b))
}
