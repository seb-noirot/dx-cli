package utils

import "fmt"

func Println(verbose bool, a ...any) {
	if verbose {
		fmt.Println(a...)
	}
}

func Printf(verbose bool, format string, a ...any) {
	if verbose {
		fmt.Printf(format, a...)
	}
}
