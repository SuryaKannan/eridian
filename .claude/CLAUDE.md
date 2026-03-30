# Eridian — CLAUDE.md

@../README.md
@../docs/img/architecture

## Key Design Decisions

- **Single Go binary** — no Python service, no IPC, no port management. Audio capture, MFCC feature extraction, and KNN matching all live in Go.
- **Centroid matching, not a classifier** — adding a new word is just a new centroid. No retraining. Inference is nearest-neighbour in embedding space.
- **MFCC for embeddings** — audio features are extracted via Mel-Frequency Cepstral Coefficients directly in Go, replacing the previous Wav2Vec 2.0 / Python approach.
- **Auto-rebuild after edit/clean** — any DB mutation triggers centroid rebuild. User never thinks about centroid state.
- **Meaning units, not sentences** — no LLM reordering or sentence construction. Explicitly out of scope.

## Storage Layout

```
~/.eridian/
    config.json        { "active": "elvish" }
    elvish.db
    japanese.db
```

One SQLite DB per language. Each stores: MFCC embedding vector and English label (see `Entry` model). Centroids (averaged embeddings per label) stored separately for fast KNN lookup. Active language tracked in config.json.

## Tech Stack

| Layer | Tech |
|-------|------|
| CLI | Go (Cobra) |
| TUI | Bubble Tea v2 |
| Audio capture | Go (portaudio or similar) |
| Feature extraction | MFCC (in Go) |
| Matching | KNN centroid lookup |
| Storage | SQLite (GORM) |

## Style Guidelines

- **Go** — idiomatic, no clever abstractions. Thin CLI layer. Error messages must be actionable — tell the user what to do, not just what went wrong.
- Never add LLM-based reordering or sentence construction.