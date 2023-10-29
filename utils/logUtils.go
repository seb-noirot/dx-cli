package utils

import "fmt"

func Println(verbose bool, a ...interface{}) {
	if verbose {
		fmt.Println(a...)
	}
}

func Printf(verbose bool, format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}
