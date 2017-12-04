package main

import (
	"fmt"
	"strconv"
	"reflect"
	//"os"
	//"strings"
	//"github.com/daviddengcn/go-colortext"
)

func main() {
	stringValue := "24.90"
	stringToBool, err := strconv.ParseBool(stringValue)
	stringToFloat, err := strconv.ParseFloat(stringValue, 64)

	intNumber := 34

	intNumberToStr := strconv.Itoa(intNumber)

	i, err := strconv.ParseFloat(intNumberToStr, 32)
	intNumberToFloat32 := float32(i)
	intNumberToFloat64 := float64(i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v with type %s\n", stringToBool, reflect.TypeOf(stringToBool))
	fmt.Printf("%v with type %s\n", stringToFloat, reflect.TypeOf(stringToFloat))
	fmt.Printf("%v with type %s\n", intNumber, reflect.TypeOf(intNumber))
	fmt.Printf("%v with type %s\n", intNumberToStr, reflect.TypeOf(intNumberToStr))
	fmt.Printf("%v with type %s\n", intNumberToFloat32, reflect.TypeOf(intNumberToFloat32))
	fmt.Printf("%v with type %s\n", intNumberToFloat64, reflect.TypeOf(intNumberToFloat64))



	//fmt.Println("System environment:")
	//for _, e := range os.Environ() {
	//	pair := strings.Split(e, "=")
	//	fmt.Printf("Name: %s", pair[0])
	//	ct.Foreground(ct.Yellow, false)
	//	fmt.Printf(" Value: %s\n", pair[1])
	//	ct.ResetColor()
	//}
}
