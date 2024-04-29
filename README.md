# mult

✨ Overview
---

Run a command multiple times and glance at the outputs.

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
    -n=10 \
    curl -s -i https://httpbin.org/delay/1
```

### Run sequentially

```bash
mult \
    -n=10 \
    -s=true \
    curl -s -i https://httpbin.org/delay/1
```
