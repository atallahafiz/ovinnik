# Hard Constraints — Ovinnik

1. Scanner = Checkov only. No Gitleaks/Trivy/Semgrep in v0.1.
2. DB = SQLite. No Postgres, no pgvector, no ORM.
3. LLM = Ollama with OpenAI-compatible fallback. No api.openai.com.
4. AI does NOT auto-apply patches. Human approves every remediation.
5. `findings.raw` is immutable after insert.
6. Every LLM call MUST go through audit/logger.py AND guard/secret_scan.py.
7. Every patch MUST pass guard/output_check.py before display.
8. Patch scope = access, encryption, logging only. Reject tag/description-only.
9. CLI is a thin client. NO business logic in cli.py.
10. Python 3.10+, Typer, Pydantic v2. No FastAPI, no Next.js.

If a request conflicts with these, refuse and cite the constraint number.
