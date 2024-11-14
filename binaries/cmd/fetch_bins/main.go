package main

import (
	"github.com/steebchen/prisma-client-go/binaries"
	"github.com/steebchen/prisma-client-go/logger"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	outDir := filepath.Join("./", "out")
	outDir2 := filepath.Join("./", "out2")
	err := os.MkdirAll(outDir, 0755)
	if err != nil {
		panic(err)
	}
	err = binaries.FetchAllNative(outDir)
	if err != nil {
		panic(err)
	}
	err = filepath.WalkDir(outDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(d.Name()) != ".tmp" {
			f, err := os.Open(filepath.Join(outDir, path))
			if err != nil {
				logger.Debug.Printf("Failed to open %s", path)
				return nil
			}
			newFile, err := os.Create(filepath.Join(outDir2, d.Name()))
			if err != nil {
				logger.Debug.Printf("Failed to open %s", path)
				return nil
			}
			_, err = io.Copy(f, newFile)
			if err != nil {
				logger.Debug.Printf("Failed to copy %s", path)
			}
			f.Close()
			newFile.Close()
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
