package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/atallahafiz/ovinnik/internal/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored findings from the local Ovinnik database",
	RunE: func(cmd *cobra.Command, args []string) error {
		dbPath := filepath.Join(".ovinnik", "ovinnik.db")
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			return fmt.Errorf("database not found. Please run 'ovinnik init' first")
		}

		conn, err := db.Connect()
		if err != nil {
			return err
		}
		defer conn.Close()

		rows, err := conn.Query(`
			SELECT id, check_id, resource, priority
			FROM findings
			ORDER BY id
		`)
		if err != nil {
			return fmt.Errorf("failed to query findings: %w", err)
		}
		defer rows.Close()

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

		fmt.Fprintln(w, "ID\tCHECK_ID\tRESOURCE\tPRIORITY")
		fmt.Fprintln(w, "--\t--------\t--------\t--------")

		count := 0

		for rows.Next() {
			var (
				id       int64
				checkID  string
				resource string
				priority int
			)

			if err := rows.Scan(&id, &checkID, &resource, &priority); err != nil {
				return fmt.Errorf("failed to scan finding row: %w", err)
			}

			fmt.Fprintf(w, "%d\t%s\t%s\t%d\n", id, checkID, resource, priority)
			count++
		}

		if err := rows.Err(); err != nil {
			return fmt.Errorf("failed while iterating finding rows: %w", err)
		}

		if err := w.Flush(); err != nil {
			return fmt.Errorf("failed to flush table output: %w", err)
		}

		if count == 0 {
			fmt.Println("No findings stored yet. Run 'ovinnik scan --from-json <path>' first.")
			return nil
		}

		fmt.Printf("\nTotal findings: %d\n", count)

		return nil
	},
}