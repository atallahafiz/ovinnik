package cmd

import (
	"fmt"
	"os"

	"github.com/atallahafiz/ovinnik/internal/db"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Ovinnik workspace and database",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Create .ovinnik/ directory
		if err := os.MkdirAll(".ovinnik", 0755); err != nil {
			return fmt.Errorf("failed to create .ovinnik directory: %w", err)
		}

		// 2. Initialize database connection
		conn, err := db.Connect()
		if err != nil {
			return err
		}
		defer conn.Close()

		// 3. Create findings table schema
		if err := db.InitSchema(conn); err != nil {
			return err
		}

		fmt.Println("Ovinnik successfully initialized in .ovinnik/ directory")
		return nil
	},
}