package main

import (
	"fmt"
	"reflect"
)

func swap(a, b interface{}) {
	// 允許顯式 panic
	if a == nil || b == nil {
		panic("nil pointer received")
	}

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// a, b 必須是指向可設定值的 pointer
	if va.Kind() != reflect.Ptr || vb.Kind() != reflect.Ptr {
		panic("swap requires pointer arguments")
	}

	// 取出指向的實際值
	va = va.Elem()
	vb = vb.Elem()

	// 型別必須一致
	if va.Type() != vb.Type() {
		panic("type mismatch")
	}

	// 地址不變 → 直接交換內容
	tmp := reflect.New(va.Type()).Elem()
	tmp.Set(va)
	va.Set(vb)
	vb.Set(tmp)
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
