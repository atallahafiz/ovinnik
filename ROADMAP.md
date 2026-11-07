# Roadmap — 14 Days

| Days | Branch | Done when |
|---|---|---|
| 1–2 | feat/init-scan-checkov | `ovinnik init && ovinnik scan` populates findings; smoke test with fixture passes |
| 3–4 | feat/inventory | `ovinnik inventory` populates assets from CODEOWNERS + asset.yaml |
| 5–6 | feat/triage-rules | `ovinnik triage` outputs P0–P3 from rules, no LLM |
| 7–8 | feat/llm-client | LLM provider abstraction + audit log + secret guard |
| 9–10 | feat/explain | `ovinnik explain` produces validated JSON narrative |
| 11–12 | feat/fix | `ovinnik fix` produces diff that passes output_check |
| 13 | feat/verify-evidence | `ovinnik verify` + `ovinnik evidence` work end-to-end |
| 14 | docs/readme-demo | v0.1.0a1 tagged and released |
