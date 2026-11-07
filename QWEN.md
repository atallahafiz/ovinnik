# Ovinnik — Project Instructions

You are helping build Ovinnik, a CLI tool that triages Checkov findings
on Terraform. Read the following before writing any code:

- @SPEC.md
- @CONSTRAINTS.md
- @ROADMAP.md

## Hard rules (non-negotiable)
1. Scanner v0.1 = Checkov only. Do NOT add other scanners.
2. Database = SQLite only. Do NOT add Postgres/pgvector.
3. Do NOT touch LLM code before day 7 of ROADMAP.
4. `findings.raw` is immutable. No UPDATE after insert.
5. CLI is a thin client. NO business logic in cli.py.
6. If unsure, STOP and ask. Do not improvise.

## Naming
- Package: `ovinnik` (not `privai`)
- CLI command: `ovinnik`
- Config/DB dir: `.ovinnik/`
- Any reference to "PrivAI" in prior documents means "Ovinnik"

## Workflow
- One 2-day block per session. Stop when the block's "done when" is met.
- Do NOT proceed to next block without explicit instruction.
- After each block: show a summary of files changed and how to test.
