# Ovinnik — Specification v0.1 (LOCKED)

> Dokumen ini adalah kontrak. Setiap perubahan harus disengaja, bukan
> hasil improvisasi AI.

## 1. Product
CLI tool that triages Checkov findings on Terraform. Reads findings,
correlates with infra context, explains risk, drafts patch, produces
Markdown evidence. Runs locally. No cloud.

## 2. Position
Scanner finds → AI decides & explains → Human approves → Scanner verifies.

## 3. Scope v0.1
- Scanner: Checkov (Terraform), whitelist 10–15 check_id
- DB: SQLite at `.ovinnik/ovinnik.db`
- Priority: rule-based (exposure + data_class + rule class) → P0–P3
- LLM: Qwen2.5-Coder via Ollama (7B dev, 32B user)
- Output: narration + draft patch + Markdown evidence
- CLI commands: init, scan, inventory, triage, explain, fix, verify, evidence

## 4. Not in scope
Web UI, multi-scanner, Antares, multi-tenant, SSO, cloud LLM,
auto-apply patches, post-incident forensics.

## 5. Database schema
See db.py. Tables: findings, assets, remediations, evidence, audit_log.
`findings.raw` immutable.

## 6. LLM output
Two prompts: explain.md → {what, why_it_matters, attack_scenario,
priority, priority_reason, citations[]}; fix.md → {diff, rationale,
risk_of_change, test_to_run[], needs_human}. JSON strict, Pydantic v2.

## 7. Guards
- secret_scan.py before every LLM call
- output_check.py after every patch draft: reject patch outside
  target_file, reject patch that only adds tags/description
- audit/logger.py wraps every LLM call

## 8. Metrics
Priority accuracy (P0 correctness) and patch applicability
(terraform validate + apply + re-scan clean). Measured on 30 findings.
Report honest numbers in README.

## 9. Naming
Package `ovinnik`, CLI `ovinnik`, dir `.ovinnik/`.

## 10. License
Apache-2.0 + DCO for contributors.
