# touchx
`touch` extra, **touchx** is like `touch` but automatically creates any missing parent directories before creating the file.

## ✨ Features

- Automatically creates missing parent directories (like `mkdir -p`)
- Dependency-free (pure Go)

## 📦 Installation

### Clone and build manually

```bash
git clone https://github.com/Hidayathamir/touchx.git
cd touchx
go build -o touchx
```

```bash
./touchx path/to/your/file.txt
```