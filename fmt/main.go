package main

import "fmt"

func main() {
	scan()
}

func scan() {
	// "1 2 3 4 5 a b c d e\nfg h"
	s := ""
	s2 := ""
	n, err := fmt.Scan(&s, &s2)
	if err != nil {
		panic(err)
	}
	fmt.Println(n) // scanできた数
	fmt.Println(s)
	fmt.Println(s2)
}

func append() {
	fmt.Println(string(fmt.Appendln([]byte("a0"), "a1", "a2", "a3")))
}

func format() {

	s := "hello world"
	ps := &s
	// ダブルクォートありで出力
	fmt.Printf("%q\n", s)

	fmt.Printf("%v\n", ps)

	fmt.Printf("% x\n", s)

	// [n] で指定できる
	fmt.Printf("%[2]d %[1]d\n", 11, 22)

	// %!で出力されたらエラーである。

	f("hello world")
	f(1)
	f([]int{1234567890, 12345})
}

func f(v any) {
	// %#v ソースコード上での表現のまま出力
	fmt.Printf("%#v\n", v)

	fmt.Printf("%T\n", v)

	// literal %
	fmt.Printf("%%\n")

	fmt.Printf("%t\n", v)

	fmt.Printf("-----------------------\n")
}
