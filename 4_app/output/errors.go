package output

import (
	"fmt"
)

func PrintError(value any) {
	switch value.(type) {
	case string:
		fmt.Println("This is a string error")
	case int:
		fmt.Println("This is a integer error")
	default:
		fmt.Println("Unknown type of error")
	}

	fmt.Println(value)
}
