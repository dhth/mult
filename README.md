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
mult command --you=want to run
```

<p align="center">
  <img src="https://tools.dhruvs.space/images/mult/mult-usage-1.gif" alt="Usage" />
</p>

🧰 Use Cases
---

`mult` can come in handy in a few scenarios, such as:

- Investigating inconsistent responses from a web server
- Checking the outcomes of a flaky test suite
- Running quick and dirty performance/stress tests

💾 Installation
---

**Arch Linux**:

```sh
yay -S mult
```

**homebrew**:

```sh
brew install dhth/tap/mult
```

**go**:

```sh
go install github.com/dhth/mult@latest
```

Or get the binaries directly from a
[release](https://github.com/dhth/mult/releases).

⚡️ Usage
---

```text
Usage: mult [flags]

Flags:
  -F    whether to stop after first failure
  -S    whether to stop after first success
  -d int
        time (in ms) to sleep for between runs
  -i    accept flag values interactively (takes precendence over -n)
  -n int
        number of times to run the command (default 5)
  -s    whether to invoke the command sequentially
```

### Specify number of runs

```bash
mult \
    -n=10 \
    yourcommand --with=flags --and args as usual
```

### Ask for number of runs

```bash
mult \
    -i \
    yourcommand --with=flags --and args as usual
```

### Run sequentially

By default, `mult` executes all runs concurrently. Use `-s` for sequentially
execution.

```bash
mult \
    -s \
    yourcommand --with=flags --and args as usual
```

### Add delay (in milliseconds) between runs

```bash
mult \
    -s \
    -d=500 \
    yourcommand --with=flags --and args as-usual
```

### Stop at first failure

```bash
mult \
    -s \
    -F \
    yourcommand --with=flags --and args as usual
```

*Note: `-d`, `-F`, `-S` only apply in sequential run mode.*
