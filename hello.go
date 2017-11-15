package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/daviddengcn/go-colortext"
)

func main() {
	fmt.Println("System environment:")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("Name: %s", pair[0])
		ct.Foreground(ct.Yellow, false)
		fmt.Printf(" Value: %s\n", pair[1])
		ct.ResetColor()
	}
}
