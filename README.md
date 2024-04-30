# mult

✨ Overview
---

Run a command multiple times and glance at the outputs.

```bash
mult command --you=want to run
```

<p align="center">
  <img src="./assets/mult.png?raw=true" alt="Usage" />
</p>


💾 Installation
---

**go**:

```sh
go install github.com/dhth/mult@latest
```

⚡️ Usage
---

### Run concurrently

```bash
mult \
    curl -s -i https://httpbin.org/delay/1
```

### Specify number of runs

```bash
mult \
    -n=10 \
    curl -s -i https://httpbin.org/delay/1
```

### Ask for number of runs

```bash
mult \
    -i \
    curl -s -i https://httpbin.org/delay/1
```

### Run sequentially

```bash
mult \
    -s \
    curl -s -i https://httpbin.org/delay/1
```

### Add delay (in milliseconds) between runs

```bash
mult \
    -s \
    -delay-ms=500 \
    curl -s -i https://httpbin.org/delay/1
```

### Stop at first failure

```bash
mult \
    -s \
    -ff \
    curl -s -i https://httpbin.org/delay/1
```

*Note: `-delay-ms`, `-ff` only apply in sequential run mode.*
