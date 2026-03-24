# Eridian — CLAUDE.md

@../README.md
@../docs/img/architecture

## Key Design Decisions

- **Centroid matching, not a classifier** — adding a new word is just a new centroid. No retraining. Inference is nearest-neighbour in embedding space.
- **Go owns the DB, Python never writes** — avoids concurrent write conflicts. Python is a pure compute layer: audio in, vector out.
- **Ephemeral port** — Go finds a free port at startup and passes it to Python. No hardcoded ports, no config needed.
- **Auto-rebuild after edit/clean** — any DB mutation triggers centroid rebuild. User never thinks about centroid state.
- **Meaning units, not sentences** — no LLM reordering or sentence construction. Explicitly out of scope.

## Storage Layout

```
~/.eridian/
    config.json        { "active": "elvish" }
    elvish.db
    japanese.db
```

One SQLite DB per language. Each stores: audio path, embedding vector, English label, timestamp, speaker ID. Active language tracked in config.json.

## Tech Stack

| Layer | Tech |
|-------|------|
| CLI | Go (Cobra) |
| TUI (`eridian edit`) | Bubble Tea |
| Audio capture | Go (portaudio or similar) |
| ML service | Python, FastAPI |
| Embeddings | Wav2Vec 2.0 (HuggingFace) |
| Python dependency mgmt | UV |
| Storage | SQLite |

## Style Guidelines

- **Go** — idiomatic, no clever abstractions. Thin CLI layer. Error messages must be actionable — tell the user what to do, not just what went wrong.
- **Python** — thin HTTP wrapper around HuggingFace inference. No business logic.
- Never add LLM-based reordering or sentence construction.
- Never have Python write directly to SQLite.