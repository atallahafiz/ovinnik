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

		if err := os.MkdirAll(".ovinnik", 0755); err != nil {
			return fmt.Errorf("gagal membuat direktori .ovinnik: %w", err)
		}

		conn, err := db.Connect()
		if err != nil {
			return err
		}
		defer conn.Close()

		if err := db.InitSchema(conn); err != nil {
			return err
		}

		fmt.Println("Ovinnik berhasil diinisialisasi di direktori .ovinnik/")
		return nil
	},
}