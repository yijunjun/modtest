package modtest

import (
	"fmt"
	"reflect"
)

func Hi(greeting string)string{
	return fmt.Sprintf("echo hi %v", greeting)
}

func ReflectLaw(){
	// 第一条规则，接口到反射对象
	x := 10
	xVal := reflect.ValueOf(x)
	fmt.Println("xVal:", xVal, xVal.String())
	xType := reflect.TypeOf(x)
	fmt.Println("xType:", xType, xType.String())

	// 第二条规则，反射对象到接口
	xInter := xVal.Interface()
	fmt.Println("xInter:", xInter)

	// 第三条规则,修改反射对象值
	px := &x
	pxVal := reflect.ValueOf(px)
	fmt.Println("pxVal:", pxVal, pxVal.String())

	fmt.Println("can set:", xVal.CanSet(), pxVal.CanSet())

	pxEle := pxVal.Elem()

	fmt.Println("pxEle:", pxEle, pxEle.CanSet())

	pxEle.SetInt(1000)

	fmt.Println("set:", x)
}