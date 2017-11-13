package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("System environment:")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Printf("Name: %s | Value: %s\n", pair[0], pair[1])
	}
}
