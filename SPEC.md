# Ovinnik — Specification v0.1 (Go, No AI)

## 1. Product
Ovinnik is a high-performance CLI tool written in Go. It ingests Checkov 
findings (JSON), correlates them with infrastructure context (YAML/CODEOWNERS), 
and outputs prioritized, deterministic security reports. Zero AI. Zero Cloud.

## 2. Tech Stack
- Language: Go 1.22+
- CLI Framework: github.com/spf13/cobra
- Database: SQLite via modernc.org/sqlite (pure Go, no CGO)
- Config Parsing: gopkg.in/yaml.v3

## 3. Core Logic
- Read Checkov JSON output.
- Read `asset.yaml` (exposure, data_class) and `CODEOWNERS`.
- Rule-based priority: match Checkov `check_id` against `rules.yaml`.
- `ovinnik report` generates a standalone `report.md`.

## 4. Database Schema (SQLite)
- `findings`: id, check_id, resource, file_path, raw_json, priority.
- `assets`: resource_ref, owner_team, exposure, data_class.
