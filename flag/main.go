package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "default", "usageeee")
	flag.Parse()
	fmt.Println(flag.Arg(0))
	fmt.Println(flag.Args())
	fmt.Println(*name)
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value, f.Usage, f)
	})
}
