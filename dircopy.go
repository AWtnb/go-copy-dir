package main

import (
	"os"
	"path/filepath"
)

func Copy(src string, newPath string) error {
	err := copy(src, newPath)
	return err
}

func copy(src string, newPath string) error {
	fs, err := os.Stat(src)
	if err != nil {
		return err
	}

	if fs.IsDir() {
		if err := addDir(src, newPath); err != nil {
			return err
		}
	} else {
		if err := addFile(src, newPath); err != nil {
			return err
		}
	}

	return nil
}

func addDir(src string, newPath string) error {
	if err := os.Mkdir(newPath, 0700); err != nil {
		return err
	}

	fi, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, f := range fi {
		sp := filepath.Join(src, f.Name())
		np := filepath.Join(newPath, f.Name())
		err := copy(sp, np)
		if err != nil {
			return err
		}
	}

	return nil
}

func addFile(src string, newPath string) error {
	d, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	df, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer df.Close()

	if _, err = df.Write(d); err != nil {
		return err
	}

	return nil
}
