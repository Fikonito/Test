// newArray
package main

import (
	"fmt"
)

func newArray() {
	second := [5]int{1, 2, 10, 15}
	fmt.Printf("newArray is: len(%d) - %v", len(second), second)
}
