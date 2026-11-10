package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ovinnik",
	Short: "Ovinnik is a high-performance CLI security tool.",
	Long:  `Ovinnik adalah alat CLI keamanan berkinerja tinggi yang bersifat deterministik dan berbasis aturan (rule-based) untuk menganalisis temuan static analysis.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(scanCmd)
}