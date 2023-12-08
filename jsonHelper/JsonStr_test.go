package jsonHelper_test

import (
	"Database_Homework/jsonHelper"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var a jsonHelper.JsonStr
	var b jsonHelper.JsonStr
	a.JsonArrayInit()
	a.JsonArrayAddStr("hello")
	a.JsonArrayEnd()
	b.JsonArrayInit()
	b.JsonArrayEnd()

	a.CombineWith(b)
	fmt.Println(string(a.Str))
	b.CombineWith(a)
	fmt.Println(string(a.Str))
	b.Str = []byte("[\"world\"]")
	a.CombineWith(b)
	fmt.Printf(string(a.Str))
}
