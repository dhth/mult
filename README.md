<p align="center">
  <h1 align="center">mult</h1>
  <p align="center">
    <a href="https://github.com/dhth/mult/actions/workflows/main.yml"><img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/dhth/mult/main.yml?style=flat-square"></a>
    <a href="https://github.com/dhth/mult/actions/workflows/vulncheck.yml"><img alt="Vulnerability Check" src="https://img.shields.io/github/actions/workflow/status/dhth/mult/vulncheck.yml?style=flat-square&label=vulncheck"></a>
    <a href="https://github.com/dhth/mult/releases/latest"><img alt="Latest release" src="https://img.shields.io/github/release/dhth/mult.svg?style=flat-square"></a>
    <a href="https://github.com/dhth/mult/releases/latest"><img alt="Commits since latest release" src="https://img.shields.io/github/commits-since/dhth/mult/latest?style=flat-square"></a>
  </p>
</p>

Run a command multiple times and glance at the outputs.

```bash
mult -- curl -sSif https://flaky.service.example
```

[![demo](https://asciinema.org/a/E06efxcFkA4RSSTI7keeUtFqp.svg)](https://asciinema.org/a/E06efxcFkA4RSSTI7keeUtFqp)

🧰 Use Cases
---

`mult` can come in handy in a few scenarios, such as:

- Investigating inconsistent responses from a web server
- Checking the outcomes of a flaky test suite
- Running quick and dirty performance/stress tests

💾 Installation
---

### Pre-built binaries

Download a pre-built binary from the [latest
release](https://github.com/dhth/mult/releases/latest). See [Verifying release
artifacts](#-verifying-release-artifacts) for instructions on verifying your
download.

### Install from source

You can also install from source using the `go` toolchain:

```sh
go install github.com/dhth/mult@latest
```

⚡️ Quick start
---

```sh
mult -n 10 -- curl -sSif https://flaky.service.example
```

This runs the command 10 times, shows the status and duration of each run in the
left pane, and displays the selected run's output in the right pane. Press
`<tab>` to switch panes and `?` for help.

`mult` runs commands concurrently by default. Without `--num-runs`, it starts 5
runs. Use `--sequential` to run them one at a time.

The `--` separator prevents command flags from being interpreted as `mult`
flags. `mult` executes the command directly, without invoking a shell, and
captures its combined standard output and standard error. To use shell features
such as pipes, redirects, or variable expansion, invoke a shell explicitly:

```sh
mult -- sh -c 'curl -sS https://example.com | wc -c'
```

Each invocation receives a 1-indexed `MULT_RUN_NUM` environment variable.

🎛️ Execution Options
---

### Choose the number of runs

```sh
mult -n 10 -- yourcommand
```

The number of runs must be between 2 and 1000, inclusive. To enter it
interactively instead, use:

```sh
mult -i -- yourcommand
```

### Run sequentially

```sh
mult -s -- yourcommand
```

[![sequential-runs](https://asciinema.org/a/1AeDGwgIF7bJ7QcV73Zhgqkb3.svg)](https://asciinema.org/a/1AeDGwgIF7bJ7QcV73Zhgqkb3)

### Pause between runs

Use `--delay` to wait between sequential runs. The value is in milliseconds.

```sh
mult -s -d 500 -- yourcommand
```

[![add-delay-between-runs](https://asciinema.org/a/vdGk7tf5sXYFZyb77PmN7ciaG.svg)](https://asciinema.org/a/vdGk7tf5sXYFZyb77PmN7ciaG)

### Stop on the first failure or success

```sh
mult -s -F -- yourcommand # stop on the first failure
mult -s -S -- yourcommand # stop on the first success
```

[![stop-on-the-first-failure-or-success](https://asciinema.org/a/IeasGG4AVDLlLfTxKqETjskDD.svg)](https://asciinema.org/a/IeasGG4AVDLlLfTxKqETjskDD)

### Follow the latest run

In sequential mode, follow mode automatically selects the latest run:

```sh
mult -s -f -- yourcommand
```

Press `<ctrl+f>` to toggle follow mode from the run list or output pane.

The `--delay`, `--follow`, `--stop-on-first-failure`, and
`--stop-on-first-success` flags only apply in sequential mode.

📟 TUI
---

![TUI](https://tools.dhruvs.space/images/mult/v0-3-0/tui.png)

`mult` has three views:

- **Command Run List View** — Shows the status and duration of each run
- **Output View** — Shows the combined output of the selected run
- **Help View** — Shows the available keyboard shortcuts

### Keyboard shortcuts

#### General

| Key                     | Action                     |
|-------------------------|----------------------------|
| `<tab>` / `<shift+tab>` | Switch focus between panes |
| `?`                     | Show or hide the help view |
| `q` / `<esc>`           | Go back or quit            |
| `<ctrl+c>`              | Quit immediately           |

#### Command Run List View

| Key        | Action                              |
|------------|-------------------------------------|
| `j` / `↓`  | Go to next run                      |
| `k` / `↑`  | Go to previous run                  |
| `l` / `→`  | Go to next page (if applicable)     |
| `h` / `←`  | Go to previous page (if applicable) |
| `g`        | Go to start of the list             |
| `G`        | Go to the end of the list           |
| `<ctrl+r>` | Restart all runs                    |
| `<ctrl+f>` | Toggle follow mode                  |

#### Output View

| Key        | Action             |
|------------|--------------------|
| `j` / `↓`  | Scroll output down |
| `k` / `↑`  | Scroll output up   |
| `l` / `→`  | Go to next run     |
| `h` / `←`  | Go to previous run |
| `<ctrl+r>` | Restart all runs   |
| `<ctrl+f>` | Toggle follow mode |

`>_` CLI reference
---

```text
mult [flags] -- <command>
```

| Flag                            | What it does                                              |
|---------------------------------|-----------------------------------------------------------|
| `-n`, `--num-runs <number>`     | Set the number of runs (default: 5)                       |
| `-i`, `--interactive`           | Prompt for the number of runs; takes precedence over `-n` |
| `-s`, `--sequential`            | Run commands sequentially                                 |
| `-d`, `--delay <milliseconds>`  | Wait between sequential runs                              |
| `-f`, `--follow`                | Start sequential runs with follow mode enabled            |
| `-F`, `--stop-on-first-failure` | Stop sequential execution after the first failure         |
| `-S`, `--stop-on-first-success` | Stop sequential execution after the first success         |
| `-h`, `--help`                  | Show help                                                 |

Run `mult --help` for the authoritative command-line reference.

🔐 Verifying release artifacts
---

Each release includes checksums for all artifacts. The checksum file is signed
using [cosign](https://docs.sigstore.dev/cosign/installation/) (version
`3.1.3`).

Replace `x.y.z` below with the release version you want to verify.

1. Get the checksum and cosign signature from the release:

    ```shell
    curl -sSLO https://github.com/dhth/mult/releases/download/vx.y.z/mult_x.y.z_checksums.txt
    curl -sSLO https://github.com/dhth/mult/releases/download/vx.y.z/mult_x.y.z_checksums.txt.sigstore.json
    ```

2. Verify the checksum file's signature:

    ```shell
    cosign verify-blob \
        --bundle mult_x.y.z_checksums.txt.sigstore.json \
        --certificate-identity-regexp 'https://github\.com/dhth/mult/\.github/workflows/.+' \
        --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
        mult_x.y.z_checksums.txt
    ```

3. Download the archive for your platform and validate its checksum. For example,
   for Linux x86-64:

    ```shell
    curl -sSLO https://github.com/dhth/mult/releases/download/vx.y.z/mult_x.y.z_linux_amd64.tar.gz
    sha256sum --ignore-missing -c mult_x.y.z_checksums.txt
    ```

4. Once both checks pass, extract the archive:

    ```shell
    tar -xzf mult_x.y.z_linux_amd64.tar.gz
    ./mult -h
    # profit!
    ```
