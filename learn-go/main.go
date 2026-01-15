package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	var b string
	a = 1230
	b = "bonjour"
	var c int = 123
	d := "hi"

	fmt.Println("Hello World! ", a, b, c, d)

	// beyondHello()
	// learnTypes()
	// learnControlFlow()
	// learnFunctions()
	// learnStruct()
	// learnStructMethods()

	// learnArray()
	// learnMyString()

	learnSlice()

	// fmt.Println("++++ examples +++++")
	// RunExamples()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("Sum:", sum, "Product:", prod)
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn GO!"
	s2 := `A "raw" string literal
can include line breaks.`

	g := 'Σ' // rune type, alias for int32, holds a unicode code point

	fmt.Println("str: ", str, "\n s2:", s2, "\n g:", g)

	// fixed length

	const bVar bool = false
	var byteVar byte = 123
	fmt.Println("bVar ", bVar, " byteVar ", byteVar)
	{
		a := 1
		var b int64
		b = int64(a)
		fmt.Println("int64 b", b)
	}

	{
		a := 42
		var ptr *int
		ptr = &a
		b := *ptr
		fmt.Println("a ptr b: ", a, ptr, b)
		*ptr = 23
		fmt.Println("a ", a)
	}
}

func learnControlFlow() {
	a := 10
	c := 20
	if a != c {
		fmt.Println("a != c")
	} else {
		fmt.Println("a == c")
	}
	switch a {
	case 10:
		fmt.Println("a case 10")
	case 20:
		fmt.Println("a case 20")
	default:
		fmt.Println("a default case")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("for i: ", i)
	}
}

func f(a int, b string) error {
	return nil
}

func g() (string, bool) {
	return "hello", true
}

func h() (s string, b bool) {
	s = "World"
	b = false
	return
}

func foo() (string, error) {
	return "error msg: success", nil
}

func learnFunctions() {
	err := f(42, "hello")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("err is nil")
	}

	g_str, b := g()
	fmt.Println("str: ", g_str, " b: ", b)

	s, b := h()
	fmt.Println("s ", s, " b ", b)

	errmsg, err := foo()
	fmt.Println("errmsg ", errmsg, " err ", err)
	if err != nil {
		fmt.Println("handle error ", err)
	} else {
		fmt.Println("")
	}
}

type C struct {
	foo string
	bar int
}

func learnStruct() {
	c := C{foo: "hi", bar: 42}
	c.bar = 4242
	c.foo = "hi hi hi"
	fmt.Println("struct C ", c)

	c.bar = 123
	bar1 := c.bar
	fmt.Println("bar1", bar1)

	pc := &c
	bar2 := pc.bar
	bar3 := (*pc).bar
	fmt.Println("bar2 bar3", bar2, bar3)

}

func foo1(c *C) {
	c.foo = "hi"
}

func (c *C) foo2() {
	c.foo = "hi from foo2"
}

func (c C) foo3() {
	fmt.Println("bar ", c.bar)
}

func learnStructMethods() {
	c := C{foo: "hi", bar: 42}
	foo1(&c)
	fmt.Println("c ", c)
	c.foo2()
	fmt.Println("c ", c)
	c.foo3()
}

func learnArray() {
	var a [2]byte
	var b [2][4]byte

	b = [2][4]byte{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	fmt.Println("a ", a)
	fmt.Println("b ", b)

	b[1][3] = 123
	fmt.Println("b[1][3] ", b[1][3])
}

type MyString struct {
	ptr *byte
	len int
}

func (s MyString) String() string {
	return string(unsafe.Slice(s.ptr, s.len))
}

func learnMyString() {
	data := []byte("hello")
	s := MyString{ptr: &data[0], len: 5}
	fmt.Println("s ", s)

	go_str := "asdf"
	len_go_str := len(go_str)
	fmt.Println("len_go_str ", len_go_str)
	go_str = "hi " + go_str
	fmt.Println("go_str ", go_str)
}

func learnSlice() {
	nums := []int{1, 2, 3}
	nums2 := make([]int, 5)     // 5 zeros, len=5, cap=5
	nums3 := make([]int, 0, 10) // len=0, reserved for 10
	fmt.Println("nums ", nums, "nums2 ", nums2, "nums3 ", nums3)
}
