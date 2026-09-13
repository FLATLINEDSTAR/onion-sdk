# onion-sdk

## Project Purpose
The **onion-sdk** repository provides language‑agnostic, versioned SDKs and shared data contracts for the FLATLINEDSTAR ecosystem. It defines stable types (e.g., `OnionTarget`, `OnionObservation`, `Technology`, `Entity`, `ChangeEvent`, `SearchQuery`, `SearchResult`) and helper utilities that other libraries and applications import.

## Current Status
**Planning** – only a placeholder README exists. No code has been added yet.

## Why It Exists
Separating data models and helper functions into a dedicated SDK prevents circular dependencies and lets each component (validator, crawler, platform, etc.) evolve independently while sharing a single source of truth.

## Architecture Role
- **Library** – consumed by `onion-intelligence`, `onion-validator`, `onion-crawler`, `tech-fingerprint`, `entity-extractor`, `change-detector`, `query-engine`, `onion-cli`, and `onion-platform`.
- Exposes Go modules (primary) with optional bindings for Python and TypeScript.

## Planned Features (MVP)
1. Define core structs/interfaces for observations and results.
2. Generate language bindings via `gopy` (Go ↔ Python) and `ts-proto` (Go ↔ TypeScript).
3. Publish to package registries (Go modules, PyPI, npm) with semantic versioning.

## Installation / Usage (placeholder)
```bash
# Go (primary)
go get github.com/FLATLINEDSTAR/onion-sdk

# Python (once released)
pip install onion-sdk

# TypeScript (once released)
npm i onion-sdk
```
*Actual packages will become available after MVP release.*

## Development
- Language: **Go** (module `github.com/FLATLINEDSTAR/onion-sdk`).
- Follow the organization‑wide contribution guidelines in the `.github` repo.

## Testing
- Unit tests for each data type and JSON‑marshalling round‑trip.
- Compatibility tests for generated language bindings.

## Contributing
See the organization‑wide `CONTRIBUTING.md` in the `.github` repository.

## Roadmap
- **Phase 1** – Define core data models and CI pipeline for regeneration of bindings.
- **Phase 2** – Publish first versions of Go, Python, and TypeScript packages.
- **Phase 3** – Add helper utilities (validation, conversion, mock generators).

## Relationship to FLATLINEDSTAR Ecosystem
All other repositories import this SDK for their public interfaces. It is the foundational contract layer that enables independent evolution of each component.

## License
MIT – see LICENSE in the repository root.