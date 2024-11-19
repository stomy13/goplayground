package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	// bs := []byte{2, 3, 4, 5, 6, 3, 33, 3, 3, 3, 3, 3, 3, 3}
	// fmt.Println(binary.Size(bs))
	// fmt.Println(len(bs))

	src := []byte(`{
		"name":"mt",
		"id": 2
	}`)
	dst := bytes.NewBuffer(nil)
	json.Compact(dst, src)
	fmt.Println(dst)
}
