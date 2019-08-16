package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := &Point{X: 10, Y: 5}

	rv := reflect.ValueOf(p)
	fmt.Printf("rv.Type = %v\n", rv.Type())
	fmt.Printf("rv.Kind = %v\n", rv.Kind())
	fmt.Printf("rv.Interface = %v\n", rv.Interface())

	// 本だと `xv := rv.Field(0)` になっていてポインタの値になっていない
	xv := rv.Elem().Field(0)
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after SetInt) = %d\n", xv.Int())

	/*
	rv1 := reflect.ValueOf(1)
	rv2 := reflect.ValueOf("Hello World")
	rv3 := reflect.ValueOf([]byte{0xa,0xb})
	rv4 := reflect.ValueOf(make(chan struct{}))

	rv1.Int() // This is OK
	rv2.Int() // This causes panic
	rv3.Int() // This causes panic
	rv4.Int() // This causes panic
	 */

	rv1 := reflect.ValueOf(map[string]int{"foo": 1})
	value := rv1.MapIndex(reflect.ValueOf("foo"))
	fmt.Printf("value = %v\n", value)
	rv1.SetMapIndex(reflect.ValueOf("foo"), reflect.ValueOf(2))

	var num int64
	if rv1.Kind() == reflect.Int {
		num = rv1.Int()
	} else {
		num = 0
	}

	fmt.Printf("num = %v\n", num)
}
