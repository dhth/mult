# mult

✨ Overview
---

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

### Specify number of runs

```bash
mult \
    -n=10 \
    yourcommand --with=flags --and=args as usual
```

### Ask for number of runs

```bash
mult \
    -i \
    yourcommand --with=flags --and=args as usual
```

### Run sequentially

By default, `mult` executes all runs concurrently. Use `-s` for sequentially
execution.

```bash
mult \
    -s \
    yourcommand --with=flags --and=args as usual
```

### Add delay (in milliseconds) between runs

```bash
mult \
    -s \
    -delay-ms=500 \
    yourcommand --with=flags --and=args as usual
```

### Stop at first failure

```bash
mult \
    -s \
    -ff \
    yourcommand --with=flags --and=args as usual
```

*Note: `-delay-ms`, `-ff` only apply in sequential run mode.*
