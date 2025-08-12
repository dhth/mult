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
mult -- command --you=want --to=run
```

![Usage](https://tools.dhruvs.space/images/mult/v0-3-0/usage.gif)

🧰 Use Cases
---

`mult` can come in handy in a few scenarios, such as:

- Investigating inconsistent responses from a web server
- Checking the outcomes of a flaky test suite
- Running quick and dirty performance/stress tests

💾 Installation
---

**homebrew**:

```sh
brew install dhth/tap/mult
```

**go**:

```sh
go install github.com/dhth/mult@latest
```

**Arch Linux**:

```sh
yay -S mult
```

Or get a binary directly from a
[release](https://github.com/dhth/mult/releases). Read more about verifying the
authenticity of released artifacts [here](#-verifying-release-artifacts).

⚡️ Usage
---

```text
Usage:
  mult [flags] -- <command>

Examples:
mult -s -n 10 -d 1000 -- curl -sif -m 5 'https://some.url/that?fails=sometimes'

Flags:
  -d, --delay int               time (in ms) to sleep for between runs
  -f, --follow                  start with "follow mode" ON (ie, automatically select the latest command run)
  -h, --help                    help for mult
  -i, --interactive             accept flag values interactively (takes precendence over -n)
  -n, --num-runs int            number of times to run the command (default 5)
  -s, --sequential              whether to invoke the command sequentially
  -F, --stop-on-first-failure   whether to stop after first failure
  -S, --stop-on-first-success   whether to stop after first success
```

### Specify number of runs

```bash
mult -n=10 -- yourcommand
```

### Ask for number of runs

```bash
mult -i -- yourcommand
```

### Run sequentially

By default, `mult` executes all runs concurrently. Use `-s` for sequentially
execution.

```bash
mult -s -- yourcommand
```

### Add delay (in milliseconds) between runs

```bash
mult -s -d=500 -- yourcommand
```

### Stop at first failure

```bash
mult -s -F -- yourcommand
```

### Stop at first success

```bash
mult -s -S -- yourcommand
```

*Note: `-d`, `-F`, `-S` only apply in sequential run mode.*

📟 TUI
---

![TUI](https://tools.dhruvs.space/images/mult/v0-3-0/tui.png)

`mult`'s TUI has 3 views:
- Command Run List View
- Output View
- Help View

### Keyboard Shortcuts

**General**

| Key         | Action                     |
|-------------|----------------------------|
| `tab`       | Switch focus between panes |
| `?`         | Show help view             |
| `q` / `Esc` | Go back or quit            |
| `ctrl+c`    | Quit immediately           |

**Command Run List View**

| Key       | Action                              |
|-----------|-------------------------------------|
| `j` / `↓` | Go to next run                      |
| `k` / `↑` | Go to previous run                  |
| `l` / `→` | Go to next page (if applicable)     |
| `h` / `←` | Go to previous page (if applicable) |
| `g`       | Go to start of the list             |
| `G`       | Go to the end of the list           |
| `ctrl+r`  | Restart all runs                    |
| `ctrl+f`  | Toggle follow mode                  |

**Output View**

| Key       | Action             |
|-----------|--------------------|
| `j` / `↓` | Scroll output down |
| `k` / `↑` | Scroll output up   |
| `l` / `→` | Go to next run     |
| `h` / `←` | Go to previous run |
| `ctrl+r`  | Restart all runs   |
| `ctrl+f`  | Toggle follow mode |

🔐 Verifying release artifacts
---

In case you get the `mult` binary directly from a
[release](https://github.com/dhth/mult/releases), you may want to verify its
authenticity. Checksums are applied to all released artifacts, and the resulting
checksum file is signed using
[cosign](https://docs.sigstore.dev/cosign/installation/).

Steps to verify (replace `A.B.C` in the commands listed below with the version
you want):

1. Download the following files from the release:

    - mult_A.B.C_checksums.txt
    - mult_A.B.C_checksums.txt.pem
    - mult_A.B.C_checksums.txt.sig

2. Verify the signature:

   ```shell
   cosign verify-blob mult_A.B.C_checksums.txt \
       --certificate mult_A.B.C_checksums.txt.pem \
       --signature mult_A.B.C_checksums.txt.sig \
       --certificate-identity-regexp 'https://github\.com/dhth/mult/\.github/workflows/.+' \
       --certificate-oidc-issuer "https://token.actions.githubusercontent.com"
   ```

3. Download the compressed archive you want, and validate its checksum:

   ```shell
   curl -sSLO https://github.com/dhth/mult/releases/download/vA.B.C/mult_A.B.C_linux_amd64.tar.gz
   sha256sum --ignore-missing -c mult_A.B.C_checksums.txt
   ```

3. If checksum validation goes through, uncompress the archive:

   ```shell
   tar -xzf mult_A.B.C_linux_amd64.tar.gz
   ./mult
   # profit!
   ```
