package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	seen := make(map[uintptr]bool)
	trimValue(reflect.ValueOf(a), seen)
}

func trimValue(v reflect.Value, seen map[uintptr]bool) {
	if !v.IsValid() {
		return
	}

	// 追 pointer 但避免循環引用
	for v.Kind() == reflect.Ptr {
		ptr := v.Pointer()
		if ptr == 0 { // nil pointer
			return
		}
		if seen[ptr] { // 循環引用
			return
		}
		seen[ptr] = true

		v = v.Elem()
	}

	switch v.Kind() {

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			if f.CanSet() {
				trimValue(f, seen)
			}
		}

	case reflect.String:
		v.SetString(strings.TrimSpace(v.String()))

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			trimValue(v.Index(i), seen)
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			trimValue(val, seen)
		}
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	// 指回自己
	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
