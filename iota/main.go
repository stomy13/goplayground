//go:generate stringer -type=Country
package main

import (
	"fmt"
)

type Country int

const (
	Japan Country = iota + 1
	China
	Korea
)

func main() {
	for _, country := range []Country{Japan, China, Korea} {
		fmt.Printf("%d: %s\n", country, country)
	}
}
