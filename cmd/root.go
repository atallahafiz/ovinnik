package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ovinnik",
	Short: "Ovinnik is a high-performance CLI security tool.",
	Long:  `Ovinnik is a high-performance, deterministic, and rule-based CLI security tool for analyzing static analysis findings.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(listCmd)
}