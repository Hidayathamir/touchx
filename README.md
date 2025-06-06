# touchx
`touch` extra, **touchx** is like `touch` but automatically creates any missing parent directories before creating the file.

`touchx` is not a replacement for `touch`.

```bash
hidayat@thinkbuntu:~$ touch path/file.txt
touch: cannot touch 'path/file.txt': No such file or directory

🤮 everybody hate this error because directory `path` does not exist

hidayat@thinkbuntu:~$ touchx path/file.txt 
2025/06/06 13:53:15 INFO done
```

## 📦 Installation

### Option 1: Using go install

```bash
go install github.com/Hidayathamir/touchx@latest
```

```bash
touchx path/file.txt
```

### Option 2: Clone and build manually

```bash
git clone https://github.com/Hidayathamir/touchx.git
cd touchx
go build -o touchx
```

```bash
./touchx path/file.txt
```

move binary `./touchx` to your `PATH`.
