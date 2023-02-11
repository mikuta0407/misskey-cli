//go:build !windows
// +build !windows

package misskey

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func terminalWidth() int {
	width, _, err := terminal.GetSize(syscall.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %+v", err)
		os.Exit(1)
	}
	return width
}
