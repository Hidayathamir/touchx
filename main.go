package main

import (
	"log/slog"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		slog.Error("usage: touchx <file-path>")
		os.Exit(1)
	}

	for _, path := range os.Args[1:] {
		dir := filepath.Dir(path)

		if err := os.MkdirAll(dir, 0755); err != nil {
			slog.Error("failed to create directories", "path", path, "error", err)
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			slog.Error("failed to create file", "path", path, "error", err)
			continue
		}
		file.Close()
	}
}
