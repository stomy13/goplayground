package main

import (
	"context"
	"fmt"
)

type mySalaryKey string

const salaryKey mySalaryKey = "私の年収"

func main() {
	ctx := context.WithValue(context.Background(), salaryKey, "100000000")

	if v, ok := ctx.Value(salaryKey).(string); ok {
		fmt.Println(v)
	} else {
		fmt.Println("私の年収は取得できませんでした")
	}

	switch t := ctx.Value(salaryKey).(type) {
	case string:
		fmt.Printf("string:私の年収は%sです\n", t)
	case int:
		fmt.Printf("int:私の年収は%dです\n", t)
	case float64:
		fmt.Printf("float64:私の年収は%fです\n", t)
	default:
		fmt.Printf("default:私の年収は%vです\n", t)
	}
}
