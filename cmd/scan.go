package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atallahafiz/ovinnik/internal/db"

	"github.com/spf13/cobra"
)

var fromJSON string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and ingest findings from Checkov JSON output",
	RunE: func(cmd *cobra.Command, args []string) error {
		if fromJSON == "" {
			return fmt.Errorf("please specify the file path using the --from-json <path> flag")
		}

		// Validate database existence
		dbPath := filepath.Join(".ovinnik", "ovinnik.db")
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			return fmt.Errorf("database not found. Please run 'ovinnik init' first")
		}

		conn, err := db.Connect()
		if err != nil {
			return err
		}
		defer conn.Close()

		// Read JSON file
		data, err := os.ReadFile(fromJSON)
		if err != nil {
			return fmt.Errorf("failed to read JSON file: %w", err)
		}

		// Checkov JSON structure definitions
		type CheckovCheck struct {
			CheckID  string `json:"check_id"`
			FilePath string `json:"file_path"`
			Resource string `json:"resource"`
		}

		type CheckovResult struct {
			FailedChecks []CheckovCheck `json:"failed_checks"`
		}

		type CheckovOutput struct {
			Results CheckovResult `json:"results"`
		}

		// Parse JSON (Supports Array of Objects or Single Object format)
		var outputs []CheckovOutput
		if err := json.Unmarshal(data, &outputs); err != nil {
			var singleOutput CheckovOutput
			if err2 := json.Unmarshal(data, &singleOutput); err2 != nil {
				return fmt.Errorf("invalid Checkov JSON format: %w", err)
			}
			outputs = append(outputs, singleOutput)
		}

		// Execute database transaction
		tx, err := conn.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin database transaction: %w", err)
		}

		query := `INSERT OR IGNORE INTO findings (check_id, resource, file_path, raw_json, priority) VALUES (?, ?, ?, ?, ?)`
		stmt, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to prepare SQL statement: %w", err)
		}
		defer stmt.Close()

		count := 0
		for _, output := range outputs {
			for _, check := range output.Results.FailedChecks {
				if check.CheckID == "" || check.FilePath == "" || check.Resource == "" {
					continue // Skip incomplete data
				}
				
				rawJSON, _ := json.Marshal(check)
				priority := 0 // Default priority (will be updated in the triage phase)

				_, err := stmt.Exec(check.CheckID, check.Resource, check.FilePath, string(rawJSON), priority)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to insert finding into database: %w", err)
				}
				count++
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}

		fmt.Printf("Successfully processed and saved %d findings from %s\n", count, fromJSON)
		return nil
	},
}

func init() {
	scanCmd.Flags().StringVar(&fromJSON, "from-json", "", "Path to the Checkov JSON output file")
}