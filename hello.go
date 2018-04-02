package main

import (
	"fmt"
)

type example struct {
	flag    bool
	counter int32
	pi      float64
}

type MyInt int

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

func main() {
	ex1 := example{}
	ex2 := example{
		flag: true,
	}
	mi := MyInt(100)
	mi.showYourSelf()
	mi.add(50)
	mi.showYourSelf()
	fmt.Printf("%+v", ex1)
	fmt.Printf("%+v", ex2)

}
