package context

import (
	"github.com/spf13/cobra"
)

var ContextCmd = &cobra.Command{
	Use:   "context",
	Short: "Effortlessly Switch Between Your Contexts 🔄",
	Long: `Tired of juggling between different environments or projects? 😓

The 'context' command is your context-switching savior! 🙌
Configure and manage multiple contexts with ease, and let DevBot automatically
apply the right settings and configurations based on your context.

Define your contexts once, and switch like a pro! 🔄`,
}
