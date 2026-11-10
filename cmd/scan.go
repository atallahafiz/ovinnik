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
			return fmt.Errorf("harap tentukan path file menggunakan flag --from-json <path>")
		}

		// Validasi keberadaan database
		dbPath := filepath.Join(".ovinnik", "ovinnik.db")
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			return fmt.Errorf("database tidak ditemukan. Harap jalankan 'ovinnik init' terlebih dahulu")
		}

		conn, err := db.Connect()
		if err != nil {
			return err
		}
		defer conn.Close()

		// Baca file JSON
		data, err := os.ReadFile(fromJSON)
		if err != nil {
			return fmt.Errorf("gagal membaca file JSON: %w", err)
		}

		// Definisi struktur JSON Checkov
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

		// Parse JSON (Mendukung format Array of Objects atau Single Object)
		var outputs []CheckovOutput
		if err := json.Unmarshal(data, &outputs); err != nil {
			var singleOutput CheckovOutput
			if err2 := json.Unmarshal(data, &singleOutput); err2 != nil {
				return fmt.Errorf("format JSON Checkov tidak valid: %w", err)
			}
			outputs = append(outputs, singleOutput)
		}

		// Eksekusi transaksi database
		tx, err := conn.Begin()
		if err != nil {
			return fmt.Errorf("gagal memulai transaksi database: %w", err)
		}

		query := `INSERT OR IGNORE INTO findings (check_id, resource, file_path, raw_json, priority) VALUES (?, ?, ?, ?, ?)`
		stmt, err := tx.Prepare(query)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("gagal menyiapkan statement SQL: %w", err)
		}
		defer stmt.Close()

		count := 0
		for _, output := range outputs {
			for _, check := range output.Results.FailedChecks {
				if check.CheckID == "" || check.FilePath == "" || check.Resource == "" {
					continue // Lewati data yang tidak lengkap
				}
				
				rawJSON, _ := json.Marshal(check)
				priority := 0 // Default priority (nantinya diupdate pada tahap triage)

				_, err := stmt.Exec(check.CheckID, check.Resource, check.FilePath, string(rawJSON), priority)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("gagal menyisipkan temuan ke database: %w", err)
				}
				count++
			}
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("gagal melakukan commit transaksi: %w", err)
		}

		fmt.Printf("Berhasil memproses dan menyimpan %d temuan dari %s\n", count, fromJSON)
		return nil
	},
}

func init() {
	scanCmd.Flags().StringVar(&fromJSON, "from-json", "", "Path ke file output JSON Checkov")
}