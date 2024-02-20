package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		src     string
		newPath string
	)
	flag.StringVar(&src, "src", "", "dir path to copy")
	flag.StringVar(&newPath, "newpath", "", "new dir path")
	flag.Parse()
	os.Exit(run(src, newPath))
}

func run(src string, newPath string) int {
	err := Copy(src, newPath)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
