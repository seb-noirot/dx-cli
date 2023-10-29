package utils

import (
	"fmt"
	"log"
)

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

// LogError logs an error message along with the error
func LogError(action string, err error) {
	log.Printf("🛑 Error %s: %s\n", action, err)
}

// LogWarning logs a warning message
func LogWarning(message string) {
	log.Printf("⚠️  Warning: %s\n", message)
}

// LogInfo logs an informational message
func LogInfo(message string) {
	log.Printf("ℹ️  Info: %s\n", message)
}
