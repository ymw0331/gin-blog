package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra
// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md

var rootCmd  = &cobra.Command{
	Use:   "help",
	Short: "Help command",
	Long:  "Display help command",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
