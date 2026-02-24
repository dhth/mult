# AGENTS.md

This file provides guidance to AI coding agents when working with code in this repository.

## Project Overview

`mult` is a CLI tool that runs a command multiple times and displays outputs in an interactive TUI. Useful for investigating flaky tests, inconsistent server responses, and quick stress tests.

## Common Commands

All commands use `just` (task runner):

```bash
just build      # go build -ldflags='-s -w' .
just run        # go run .
just install    # go install -ldflags='-s -w' .
just lint       # golangci-lint run
just fmt        # gofumpt -w .
```

Note: always use `just` to run commands.

## Architecture

The project follows the Elm architecture via the Bubbletea TUI framework.

### Package Structure

- **`internal/cmd/`** — CLI setup via Cobra. Parses flags, validates config, calls `ui.RenderUI()`.
- **`internal/domain/`** — Pure data types: `CommandRun` (run state/results), `Config` (user settings), `RunStatus` enum.
- **`internal/ui/`** — All TUI logic following MVU:
  - `model.go` — Central state (run list, viewport, results cache, stats)
  - `update.go` — Message handler (keyboard input, command lifecycle events)
  - `view.go` — Render function (two-pane layout: command list + output viewport)
  - `cmds.go` — Bubbletea commands for executing shell commands and scheduling
  - `msgs.go` — Custom message types (`CmdRanMsg`, `CmdRunChosenMsg`, etc.)
  - `delegate.go` / `list_item.go` — List item rendering with status-based styling
  - `styles.go` — All lipgloss styling (Gruvbox color scheme)

### Execution Model

- **Concurrent mode** (default): All runs launch immediately via `tea.Batch()`
- **Sequential mode** (`-s`): Runs execute one-at-a-time, with optional delay (`-d`) and stop conditions (`-F`/`-S`)
- Commands run via `os/exec`, capturing combined stdout/stderr. Each run gets `MULT_RUN_NUM` env var (1-indexed).

### Key Dependencies

- `bubbletea` — TUI framework (MVU event loop)
- `bubbles` — Reusable components (list, viewport)
- `lipgloss` — Terminal styling
- `cobra` — CLI argument parsing

## Code Style

- Formatter: **gofumpt** (not gofmt)
- Linter: **golangci-lint** with revive, errorlint, goconst, and others (see `.golangci.yml`)
