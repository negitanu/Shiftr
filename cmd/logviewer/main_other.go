//go:build !windows
// +build !windows

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "logviewer は Windows 専用です。GOOS=windows でビルドしてください。")
	os.Exit(1)
}

